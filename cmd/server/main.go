package main

import (
	"fmt"
	"htmx-practice/internal/core"
	"htmx-practice/internal/handlers"
	"net/http"
	"os"
	"strings"
)

const (
	baseTemplateDir = "./resources/templates"
	staticDir       = "./static"
)

func main() {
	template, err := core.NewTemplate(baseTemplateDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "core.NewTemplate: %v", err)
		os.Exit(1)
	}
	h := handlers.New(core.App{
		Tmpl: *template,
	})

	mux := http.NewServeMux()
	mux.Handle("GET /static/", disableListFile(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir)))))
	mux.HandleFunc("GET /todos", core.Logger(h.TodoHandler.ListHandler))
	mux.HandleFunc("POST /todo", core.Logger(h.TodoHandler.CreateHandler))
	mux.HandleFunc("PUT /todo/{id}/done", core.Logger(h.TodoHandler.DoneHandler))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Fprintf(os.Stderr, "http.ListenAndServe: %v", err)
		os.Exit(1)
	}
}

func disableListFile(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}
