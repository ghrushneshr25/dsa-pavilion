package parser

import (
	"os"
	"regexp"
	"strings"
)

var reTopLevelFunc = regexp.MustCompile(`(?m)^func `)

func ExtractCode(file string) string {
	data, _ := os.ReadFile(file)
	re := regexp.MustCompile(`(?s)/\*.*?\*/`)
	out := string(re.ReplaceAll(data, []byte("")))
	out = strings.TrimSpace(trimLeadingMetadataComments(out))
	return FormatGoSource(out)
}

// ExtractPackageImportPreamble returns the package clause and import block before the first
// top-level func (used to prefix per-operation snippets in multi-approach problem docs).
func ExtractPackageImportPreamble(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		return ""
	}
	reDoc := regexp.MustCompile(`(?s)/\*.*?\*/`)
	s := string(reDoc.ReplaceAll(data, []byte("")))
	s = strings.TrimSpace(trimLeadingMetadataComments(s))
	loc := reTopLevelFunc.FindStringIndex(s)
	if loc == nil {
		return ""
	}
	pre := strings.TrimSpace(s[:loc[0]])
	pre = trimTrailingMetadataComments(pre)
	return pre
}

// trimTrailingMetadataComments removes trailing blank lines and consecutive // @... lines
// (docgen directives before the first func, e.g. // @operation / @description).
func trimTrailingMetadataComments(s string) string {
	lines := strings.Split(s, "\n")
	j := len(lines)
	for j > 0 {
		t := strings.TrimSpace(lines[j-1])
		if t == "" {
			j--
			continue
		}
		if isMetadataCommentLine(t) {
			j--
			continue
		}
		break
	}
	return strings.TrimSpace(strings.Join(lines[:j], "\n"))
}

// trimLeadingMetadataComments removes file-header lines like "// @problem: ..." that
// are used for docgen but should not appear in the rendered solution block.
func trimLeadingMetadataComments(s string) string {
	lines := strings.Split(s, "\n")
	i := 0
	for i < len(lines) {
		t := strings.TrimSpace(lines[i])
		if t == "" {
			i++
			continue
		}
		if isMetadataCommentLine(t) {
			i++
			continue
		}
		break
	}
	return strings.Join(lines[i:], "\n")
}

func isMetadataCommentLine(trimmed string) bool {
	if !strings.HasPrefix(trimmed, "//") {
		return false
	}
	rest := strings.TrimSpace(trimmed[2:])
	return strings.HasPrefix(rest, "@")
}
