---
title: {{.Title}}
sidebar_position: {{.SidebarPosition}}
---

# {{.Title}}

<br />

<table>
  <tbody>
    <tr><td>Difficulty</td><td>{{index .Meta "difficulty"}}</td></tr>
    <tr><td>Tags</td><td>{{index .Meta "tags"}}</td></tr>
    <tr><td>Time complexity</td><td>{{index .Meta "time"}}</td></tr>
    <tr><td>Space complexity</td><td>{{index .Meta "space"}}</td></tr>
  </tbody>
</table>

---

{{if .ProblemApproaches}}
{{.ProblemPreamble}}

{{range $i, $a := .ProblemApproaches}}{{if gt $i 0}}

---

{{end}}{{$a.Markdown}}

{{if or $a.Description $a.Time $a.Space}}
{{if $a.Description}}
### Description

{{$a.Description}}

{{end}}
{{if or $a.Time $a.Space}}

<p class="dsa-complexity">{{- if $a.Time -}}<span class="dsa-complexity__item"><span class="dsa-complexity__k">time</span>{{$a.Time}}</span>{{- end -}}{{- if $a.Space -}}<span class="dsa-complexity__item"><span class="dsa-complexity__k">space</span>{{$a.Space}}</span>{{- end -}}</p>

{{end}}
{{end}}

---

## Solution (Go)

```go
{{$a.Code}}
```

{{end}}
{{if .ProblemSummary}}

---

{{.ProblemSummary}}
{{end}}
{{else}}
{{.Sections}}

---

## Solution (Go)

```go
{{.Code}}
```

{{end}}
{{if .HasTests}}

---

## Tests

{{range .Subtests}}
<div className="dsa-test">

#### {{.Name}}

{{if .Desc}}
{{.Desc}}

{{end}}
```go
{{.Code}}
```

</div>
{{end}}
{{end}}
