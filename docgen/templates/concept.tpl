---
title: {{.Title}}
sidebar_position: {{.SidebarPosition}}
---

# {{.Title}}

<br />

<table>
  <tbody>
    <tr><td>Difficulty</td><td>{{index .Meta "difficulty"}}</td></tr>
{{- $tags := splitTags (index .Meta "tags") }}
{{- if $tags }}
    <tr><td>Tags</td><td><span class="dsa-tags">{{- range $tags -}}<span class="dsa-tag">{{.}}</span>{{- end -}}</span></td></tr>
{{- end }}
  </tbody>
</table>

---

{{if .ConceptDescription}}
## Description

{{.ConceptDescription}}
{{end}}

## Structure

{{if .ConceptStructureIntro}}
{{.ConceptStructureIntro}}

{{end}}
{{range .Structs}}
### {{.Name}}

{{if .Doc}}
{{.Doc}}

{{end}}
```go
{{.Code}}
```

{{end}}

{{if .OperationSubsections}}
## Operations
{{range .OperationSubsections}}
### {{.Title}}
{{.Prose}}

{{if or .Time .Space}}

<p class="dsa-complexity">{{- if .Time -}}<span class="dsa-complexity__item"><span class="dsa-complexity__k">time</span>{{.Time}}</span>{{- end -}}{{- if .Space -}}<span class="dsa-complexity__item"><span class="dsa-complexity__k">space</span>{{.Space}}</span>{{- end -}}</p>

{{end}}

```go
{{.Code}}
```

{{end}}
{{else if .ConceptOperationsMD}}
## Operations

{{.ConceptOperationsMD}}
{{else if .OperationsCode}}
## Operations

{{end}}
{{if and (eq (len .OperationSubsections) 0) .OperationsCode}}
```go
{{.OperationsCode}}
```

{{end}}
