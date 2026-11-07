package renderer

import (
	"os"
	"text/template"
)

func RenderDoc(doc Doc) {
	tmplFile := templateFile("problem.tpl")
	if doc.Type == "concept" {
		tmplFile = templateFile("concept.tpl")
	}

	tmpl := template.Must(template.ParseFiles(tmplFile))

	f, _ := os.Create(doc.Output)
	defer f.Close()

	tmpl.Execute(f, doc)
}
