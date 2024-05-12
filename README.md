# golang-htmx-practice

- [HTMX](https://htmx.org/)を触ってみたやつ
    - Todoの登録・状態の更新のみ
    - データストア部はとりあえずグローバルに持ったスライスで試している状態
- ついでにやりたかったこと
    - [html/template](https://pkg.go.dev/html/template)の扱いを知りたかった
    - [go1.22のルーティング機能](https://go.dev/blog/routing-enhancements)も試したかったのでサーバー部分は[net/http](https://pkg.go.dev/net/http)で実装
- パッケージ分けなんもわからん

## 起動方法

以下のコマンドでサーバーが起動できます。

```shell
go run cmd/server/main.go
```

[go-task](https://taskfile.dev/)が使用できる環境では以下のコマンドでサーバーが起動できます。

```shell
task server
```
