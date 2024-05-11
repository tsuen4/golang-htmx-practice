package handlers

import (
	"fmt"
	"htmx-practice/internal/core"
	"net/http"
	"os"
)

type handlers struct {
	TodoHandler
}

func New(app core.App) handlers {
	return handlers{
		TodoHandler: newTodoHandler(app),
	}
}

func response500(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	errorLog(err)
}

func errorLog(err error) {
	fmt.Fprint(os.Stderr, err.Error()+"\n")
}
