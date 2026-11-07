package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"docgen/parser"
	"docgen/renderer"
	"docgen/utils"
)

func main() {
	pavilionRoot, docgenDir, err := resolveRoots()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	renderer.SetTemplateDir(filepath.Join(docgenDir, "templates"))
	categoryOrderFile := filepath.Join(docgenDir, "category_order.txt")

	var (
		readmeOnly       = flag.Bool("readme-only", false, "Only write the dsa readme table. Does not touch docs/site.")
		codeDir          = flag.String("code", envOr("DOCGEN_CODE", filepath.Join(pavilionRoot, "..", "dsa")), "Path to the DSA Go module root (category subdirs, e.g. linked-list/)")
		docsDir          = flag.String("docs", envOr("DOCGEN_DOCS", filepath.Join(pavilionRoot, "docs")), "Output directory for generated .mdx")
		sidebar          = flag.String("sidebar", envOr("DOCGEN_SIDEBAR", filepath.Join(pavilionRoot, "sidebars.js")), "Output path for sidebars.js")
		readmeOut        = flag.String("readme", envOr("DOCGEN_README", filepath.Join(pavilionRoot, "..", "dsa", "readme.md")), "Output path for generated readme (code repo index table)")
		sitePrefix       = flag.String("base", envOr("DOCGEN_BASE", "/dsa-pavilion/"), "Site base path for internal doc links (must end with /)")
		githubRepo       = flag.String("github-repo", envOr("DOCGEN_GITHUB_REPO", "https://github.com/ghrushneshr25/dsa"), "DSA repo URL (no trailing slash) for navbar source links")
		githubBranch     = flag.String("github-branch", envOr("DOCGEN_GITHUB_BRANCH", "master"), "Branch name for GitHub blob/tree links")
		githubRoutesPath = flag.String("github-routes", envOr("DOCGEN_GITHUB_ROUTES", filepath.Join(pavilionRoot, "src", "data", "dsaGithubRoutes.json")), "Output JSON: doc id -> GitHub URL")
	)
	flag.Parse()

	if *readmeOnly {
		codeDirAbs, err := filepath.Abs(*codeDir)
		if err != nil {
			fmt.Fprintln(os.Stderr, "code path:", err)
			os.Exit(1)
		}
		readmeAbs, err := filepath.Abs(*readmeOut)
		if err != nil {
			fmt.Fprintln(os.Stderr, "readme path:", err)
			os.Exit(1)
		}
		if err := renderer.RenderReadme(codeDirAbs, categoryOrderFile, readmeAbs); err != nil {
			fmt.Fprintln(os.Stderr, "readme:", err)
			os.Exit(1)
		}
		fmt.Println("✅ readme.md generated")
		return
	}

	if *sitePrefix == "" {
		*sitePrefix = "/"
	} else if !strings.HasSuffix(*sitePrefix, "/") {
		*sitePrefix = *sitePrefix + "/"
	}

	codeDirAbs, err := filepath.Abs(*codeDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "code path:", err)
		os.Exit(1)
	}
	docsDirAbs, err := filepath.Abs(*docsDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "docs path:", err)
		os.Exit(1)
	}
	sidebarAbs, err := filepath.Abs(*sidebar)
	if err != nil {
		fmt.Fprintln(os.Stderr, "sidebar path:", err)
		os.Exit(1)
	}
	readmeAbs, err := filepath.Abs(*readmeOut)
	if err != nil {
		fmt.Fprintln(os.Stderr, "readme path:", err)
		os.Exit(1)
	}

	if err := os.RemoveAll(docsDirAbs); err != nil {
		fmt.Fprintln(os.Stderr, "remove docs dir:", err)
		os.Exit(1)
	}
	if err := os.MkdirAll(docsDirAbs, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "mkdir docs dir:", err)
		os.Exit(1)
	}

	var sidebarData []renderer.SidebarCategory

	repoRoot := strings.TrimSuffix(*githubRepo, "/")
	branch := *githubBranch
	githubRoutes := map[string]string{
		"index": repoRoot,
	}

	categories, err := os.ReadDir(codeDirAbs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read code dir:", err)
		os.Exit(1)
	}
	categories = utils.OrderedDirEntries(categories, categoryOrderFile)

	for _, category := range categories {
		if !category.IsDir() {
			continue
		}

		catName := category.Name()
		if strings.HasPrefix(catName, ".") {
			continue
		}
		catPath := filepath.Join(codeDirAbs, catName)
		docPath := filepath.Join(docsDirAbs, catName)

		files, err := filepath.Glob(filepath.Join(catPath, "*.go"))
		if err != nil {
			fmt.Fprintln(os.Stderr, "glob:", err)
			os.Exit(1)
		}
		// Skip repo folders that are not doc categories (e.g. scripts, .github)
		if len(files) == 0 {
			continue
		}

		if err := os.MkdirAll(docPath, 0o755); err != nil {
			fmt.Fprintln(os.Stderr, "mkdir category:", err)
			os.Exit(1)
		}

		var conceptItems []renderer.IndexItem
		var problemItems []renderer.IndexItem
		cat := renderer.SidebarCategory{Name: utils.FormatTitle(catName)}

		for i, src := range parser.OrderedSources(files) {
			file := src.Path
			meta := src.Meta
			docType := src.DocType
			title := src.Title
			sidebarPosition := i + 1

			slug := utils.DocSlug(title)

			var sections string
			var code string
			var problemPreamble string
			var problemApproaches []renderer.ProblemApproach
			var problemSummary string
			var structs []parser.StructInfo
			var conceptDesc, conceptStructIntro, conceptOpsMD, operationsCode string
			var opSubs []renderer.OperationSubsection

			if docType == "concept" {
				cdoc, _ := parser.ParseConceptDocBlock(file)
				inline, _ := parser.ParseConceptInlineFromFile(file)
				structs, _ = parser.ExtractStructs(file)
				if cdoc != nil {
					conceptDesc = cdoc.Description
					conceptStructIntro = cdoc.StructureIntro
					conceptOpsMD = cdoc.Operations
					for i := range structs {
						if d, ok := inline.StructDoc[structs[i].Name]; ok {
							structs[i].Doc = d
							continue
						}
						for k, v := range cdoc.StructureSubsections {
							if strings.EqualFold(strings.TrimSpace(k), structs[i].Name) {
								structs[i].Doc = v
								break
							}
						}
					}
					// Prefer in-file // @operation prose; else leading block @subsection: Operations.
					opProseList := cdoc.OperationsSubsections
					if len(inline.Operations) > 0 {
						opProseList = inline.Operations
					}
					if len(opProseList) > 0 {
						codeBySection, err := parser.ExtractOperationCodeBySection(file)
						if err == nil {
							for _, s := range opProseList {
								code := ""
								if codeBySection != nil {
									code = codeBySection[parser.NormalizeOpTitle(s.Title)]
								}
								opSubs = append(opSubs, renderer.OperationSubsection{Title: s.Title, Prose: s.Prose, Code: code, Time: s.Time, Space: s.Space})
							}
						} else {
							for _, s := range opProseList {
								opSubs = append(opSubs, renderer.OperationSubsection{Title: s.Title, Prose: s.Prose, Code: "", Time: s.Time, Space: s.Space})
							}
						}
					}
				}
				funcs, _ := parser.ExtractFunctions(file)
				operationsCode = strings.Join(funcs, "\n\n")
			} else {
				preamble, paDocs, summary, multi := parser.ParseProblemMultiApproach(file)
				if multi {
					snippets, err := parser.ExtractOperationCodeOrdered(file)
					if err == nil && len(snippets) == len(paDocs) {
						problemPreamble = preamble
						problemSummary = summary
						prefix := parser.ExtractPackageImportPreamble(file)
						for i := range paDocs {
							code := snippets[i].Source
							if prefix != "" && code != "" {
								code = parser.FormatGoSource(prefix + "\n\n" + code)
							}
							problemApproaches = append(problemApproaches, renderer.ProblemApproach{
								Markdown:    paDocs[i].Markdown,
								Code:        code,
								Description: snippets[i].Description,
								Time:        snippets[i].Time,
								Space:       snippets[i].Space,
							})
						}
					}
				}
				if len(problemApproaches) == 0 {
					sections = parser.ParseSections(file)
					code = parser.ExtractCode(file)
				}
			}

			var subtests []parser.SubtestInfo
			if docType == "problem" {
				if tp, ok := parser.TestFilePath(file); ok {
					subtests, _ = parser.ParseSubtests(tp)
				}
			}
			hasTests := len(subtests) > 0

			outFile := filepath.Join(docPath, slug+".mdx")

			renderer.RenderDoc(renderer.Doc{
				Title:                 title,
				SidebarPosition:       sidebarPosition,
				Type:                  docType,
				Sections:              sections,
				Code:                  code,
				ProblemPreamble:       problemPreamble,
				ProblemApproaches:     problemApproaches,
				ProblemSummary:        problemSummary,
				Subtests:              subtests,
				HasTests:              hasTests,
				Meta:                  meta,
				Structs:               structs,
				ConceptDescription:    conceptDesc,
				ConceptStructureIntro: conceptStructIntro,
				ConceptOperationsMD:   conceptOpsMD,
				OperationsCode:        operationsCode,
				OperationSubsections:  opSubs,
				Output:                outFile,
			})

			docID := catName + "/" + slug
			relSource := filepath.ToSlash(filepath.Join(catName, filepath.Base(file)))
			githubRoutes[docID] = repoRoot + "/blob/" + branch + "/" + relSource

			entry := renderer.IndexItem{
				Title:      title,
				Link:       *sitePrefix + docID,
				Difficulty: meta["difficulty"],
				Tags:       meta["tags"],
			}
			if docType == "concept" {
				cat.Concepts = append(cat.Concepts, docID)
				conceptItems = append(conceptItems, entry)
			} else {
				cat.Problems = append(cat.Problems, docID)
				problemItems = append(problemItems, entry)
			}
		}

		renderer.RenderIndex(renderer.Index{
			Title:        cat.Name,
			Slug:         catName,
			ConceptItems: conceptItems,
			ProblemItems: problemItems,
			Path:         filepath.Join(docPath, "index.mdx"),
		})

		githubRoutes[catName+"/index"] = repoRoot + "/tree/" + branch + "/" + catName

		sidebarData = append(sidebarData, renderer.SidebarCategory{
			Name:     cat.Name,
			Slug:     catName,
			Concepts: cat.Concepts,
			Problems: cat.Problems,
		})
	}

	renderer.RenderSidebar(sidebarData, sidebarAbs)
	renderer.RenderHome(sidebarData, filepath.Join(docsDirAbs, "index.mdx"))

	if err := renderer.RenderReadme(codeDirAbs, categoryOrderFile, readmeAbs); err != nil {
		fmt.Fprintln(os.Stderr, "readme:", err)
		os.Exit(1)
	}

	routesAbs, err := filepath.Abs(*githubRoutesPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "github-routes path:", err)
		os.Exit(1)
	}
	if err := renderer.WriteGithubRoutes(routesAbs, githubRoutes); err != nil {
		fmt.Fprintln(os.Stderr, "github routes:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Docs generated")
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// resolveRoots finds the Docusaurus repo root (this module lives in <root>/docgen).
// Works when cwd is the pavilion root, docgen/, or any subdirectory.
func resolveRoots() (pavilionRoot, docgenDir string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", "", err
	}
	dir := cwd
	for {
		if isPavilionRoot(dir) {
			return dir, filepath.Join(dir, "docgen"), nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", "", fmt.Errorf("could not find dsa-pavilion root from %s (need docusaurus.config.js and docgen/)", cwd)
}

func isPavilionRoot(dir string) bool {
	return fileExists(filepath.Join(dir, "docusaurus.config.js")) && fileExists(filepath.Join(dir, "docgen", "go.mod"))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
