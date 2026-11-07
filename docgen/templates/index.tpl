---
title: {{.Title}}
slug: /{{.Slug}}
pagination_prev: null
pagination_next: null
---

<div class="dsa-kicker">{{.Slug}}</div>

# {{.Title}}

{{if .ConceptItems}}
## Concepts

| Concept | Difficulty | Tags |
| --- | --- | --- |
{{- range .ConceptItems }}
| [{{.Title}}]({{.Link}}) | {{.Difficulty}} | <span class="dsa-tags">{{- range splitTags .Tags -}}<span class="dsa-tag">{{.}}</span>{{- end -}}</span> |
{{- end }}

{{end}}
{{- if and .ConceptItems .ProblemItems }}

---

{{ end }}
{{- if .ProblemItems }}
## Problems

| Problem | Difficulty | Tags |
| --- | --- | --- |
{{- range .ProblemItems }}
| [{{.Title}}]({{.Link}}) | {{.Difficulty}} | <span class="dsa-tags">{{- range splitTags .Tags -}}<span class="dsa-tag">{{.}}</span>{{- end -}}</span> |
{{- end }}

{{end}}
