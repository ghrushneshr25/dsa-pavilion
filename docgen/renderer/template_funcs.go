package renderer

import (
	"strings"
	"text/template"
)

func SplitTags(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" || p == "—" {
			continue
		}
		out = append(out, p)
	}
	return out
}

func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"inc":       inc,
		"splitTags": SplitTags,
	}
}
