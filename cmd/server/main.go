package main

import (
	"fmt"
	"htmx-practice/internal/core"
	"htmx-practice/internal/handlers"
	"net/http"
	"os"
)

const baseTemplateDir = "./resources/templates"

func main() {
	app := core.App{
		Template: *core.NewTemplateHandler(baseTemplateDir),
	}

	h := handlers.New(app)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", core.Logger(h.TodoHandler.ListHandler))
	mux.HandleFunc("POST /todo", core.Logger(h.TodoHandler.CreateHandler))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "http.ListenAndServe: %v", err)
		os.Exit(1)
	}
}
