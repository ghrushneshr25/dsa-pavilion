package parser

import (
	"go/ast"
	"strings"
)

// ParseOperationDirectiveDoc extracts prose from a function's leading comment block:
// @description, @time:, @space:. Lines starting with @operation are skipped.
func ParseOperationDirectiveDoc(doc *ast.CommentGroup) (description, timeC, spaceC string) {
	if doc == nil {
		return "", "", ""
	}
	text := doc.Text()
	var descParts []string
	for _, raw := range strings.Split(text, "\n") {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		low := strings.ToLower(line)
		switch {
		case strings.HasPrefix(low, "@operation"):
			continue
		case strings.HasPrefix(low, "@description"):
			idx := strings.Index(low, "@description")
			if idx >= 0 {
				rest := strings.TrimSpace(line[idx+len("@description"):])
				rest = strings.TrimPrefix(rest, ":")
				rest = strings.TrimSpace(rest)
				if rest != "" {
					descParts = append(descParts, rest)
				}
			}
		case strings.HasPrefix(low, "@time:"):
			idx := strings.Index(low, "@time:")
			if idx >= 0 {
				v := strings.TrimSpace(line[idx+len("@time:"):])
				v = strings.TrimPrefix(v, ":")
				timeC = strings.TrimSpace(v)
			}
		case strings.HasPrefix(low, "@space:"):
			idx := strings.Index(low, "@space:")
			if idx >= 0 {
				v := strings.TrimSpace(line[idx+len("@space:"):])
				v = strings.TrimPrefix(v, ":")
				spaceC = strings.TrimSpace(v)
			}
		}
	}
	description = strings.Join(descParts, " ")
	return description, timeC, spaceC
}
