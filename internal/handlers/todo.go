package handlers

import (
	"bytes"
	"fmt"
	"htmx-practice/internal/core"
	"htmx-practice/internal/entities"
	"net/http"
	"strconv"
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
	if err := h.Tmpl.ExecuteTemplate(w, "view", globalTodos); err != nil {
		response500(w, fmt.Errorf("tmpl.ExecuteTemplate: %w", err))
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

	target := "view"
	if r.Header.Get("Hx-Request") == "true" {
		target = "list"
	}

	resBuf := bytes.NewBuffer(nil)
	if err := h.Tmpl.ExecuteTemplate(resBuf, target, globalTodos); err != nil {
		response500(w, fmt.Errorf("tmpl.ExecuteTemplate: %w", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resBuf.Bytes())
}

func (h TodoHandler) DoneHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		response500(w, fmt.Errorf("r.ParseForm: %w", err))
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response500(w, fmt.Errorf("strconv.Atoi: %w", err))
		return
	}

	doneStr := r.Form.Get(idStr + "-done")
	done := false
	if doneStr == "on" {
		done = true
	}

	// update
	for i, t := range globalTodos {
		if t.Id == id {
			globalTodos[i].Done = done
			break
		}
	}

	target := "view"
	if r.Header.Get("Hx-Request") == "true" {
		target = "list"
	}

	if err := h.Tmpl.ExecuteTemplate(w, target, globalTodos); err != nil {
		response500(w, fmt.Errorf("tmpl.ExecuteTemplate: %w", err))
		return
	}
}
