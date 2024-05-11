package handlers

import (
	"fmt"
	"htmx-practice/internal/core"
	"htmx-practice/internal/entities"
	"net/http"
)

var globalTodos = []entities.Todo{
	{
		Id:      1,
		Content: "<script>alert(1)</script>",
		Done:    false,
	},
	{
		Id:      2,
		Content: "ご飯を作る",
		Done:    true,
	},
}

type TodoHandler struct {
	core.App
}

func newTodoHandler(app core.App) TodoHandler {
	return TodoHandler{
		App: app,
	}
}

func (h TodoHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := h.NewTemplate("todo/list.tpl")
	if err != nil {
		response500(w, fmt.Errorf("h.NewTemplate: %w", err))
		return
	}

	if err := tmpl.Execute(w, globalTodos); err != nil {
		response500(w, fmt.Errorf("tmpl.Execute: %w", err))
		return
	}
}

func (h TodoHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		response500(w, fmt.Errorf("r.ParseForm: %w", err))
		return
	}

	id := globalTodos[len(globalTodos)-1].Id + 1
	globalTodos = append(globalTodos, entities.Todo{
		Id:      id,
		Content: r.Form.Get("content"),
	})

	tmpl, err := h.NewTemplate("todo/list.tpl")
	if err != nil {
		response500(w, fmt.Errorf("h.NewTemplate: %w", err))
		return
	}

	if err := tmpl.Execute(w, globalTodos); err != nil {
		response500(w, fmt.Errorf("tmpl.Execute: %w", err))
		return
	}
}
