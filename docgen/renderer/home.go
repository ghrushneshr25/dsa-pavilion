package renderer

import (
	"os"
	"text/template"
)

func inc(i int) int {
	return i + 1
}

type HomePage struct {
	Categories []SidebarCategory
	Topics     int
	Concepts   int
	Problems   int
}

// RenderHome writes docs/index.mdx (site root) listing category cards from sidebar data.
func RenderHome(categories []SidebarCategory, out string) {
	data := HomePage{
		Categories: categories,
		Topics:     len(categories),
	}
	for _, c := range categories {
		data.Concepts += len(c.Concepts)
		data.Problems += len(c.Problems)
	}

	t := template.Must(template.New("home.tpl").Funcs(template.FuncMap{"inc": inc}).ParseFiles(templateFile("home.tpl")))
	f, err := os.Create(out)
	if err != nil {
		return
	}
	defer f.Close()
	_ = t.ExecuteTemplate(f, "home.tpl", data)
}
