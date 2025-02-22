{{ block "view" . }}
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta
        name="viewport"
        content="width=device-width, initial-scale=1.0"
    >
    <title>Todo</title>
    <script src="/static/js/htmx.min.js"></script>
</head>
<body>
    <h1>Todo List</h1>
    <form
        hx-post="/todo"
        hx-target="#list"
        hx-swap="innerHTML"
        hx-on::after-request="this.reset()"
        hx-trigger="submit"
    >
        <input
            type="text"
            name="content"
        />
        <button>Add</button>
    </form>
    <div id="list">
        {{ block "list" . }}
        <div>items: {{ . | len }}</div>
        <ul>
            {{ range . }}
            <li>
                <input
                    type="checkbox"
                    name="{{ .Id }}-done"
                    id="{{ .Id }}"
                    {{ if .Done }}checked{{ end }}
                    hx-put="/todo/{{ .Id }}/done"
                >{{ .Id }}: {{ .Content }}
            </li>
            {{ end }}
        </ul>
        {{ end }}
    </div>
    <script>
        document.addEventListener('htmx:responseError', (e) => {
            alert(e.detail.xhr.statusText)
        })
    </script>
</body>
</html>
{{ end }}
