<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Todo</title>
</head>
<body>
    <h1>Todo List</h1>
    <form action="/todo" method="post">
        <input type="text" name="content" /><button type="submit">Add</button>
    </form>
    <div>items: {{ . | len }}</div>
    <div>
        {{ template "list" . }}
    </div>
</body>
</html>
