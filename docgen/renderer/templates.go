package renderer

import "path/filepath"

// templateDir is the directory containing *.tpl files. main sets this to
// <pavilion>/docgen/templates so generation works from the repo root or docgen/.
var templateDir = "templates"

func SetTemplateDir(dir string) {
	if dir != "" {
		templateDir = dir
	}
}

func templateFile(name string) string {
	return filepath.Join(templateDir, name)
}
