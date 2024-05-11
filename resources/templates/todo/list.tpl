{{ define "list" }}
<ul>
    {{ range . }}
    <li><input type="checkbox" name="done[]" id="{{ .Id }}" disabled {{ if .Done }}checked{{ end }}>{{ .Id }}: {{ .Content }}</li>
    {{ end }}
</ul>
{{ end }}
