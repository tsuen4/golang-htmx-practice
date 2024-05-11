package core

import (
	"html/template"
	"path"
	"path/filepath"
)

type Template struct {
	BaseTemplateDir string
	FuncMap         template.FuncMap
}

func NewTemplateHandler(baseTemplateDir string) *Template {
	return &Template{BaseTemplateDir: baseTemplateDir}
}

func (h *Template) AddFuncMap(funcMap template.FuncMap) *Template {
	for name, f := range funcMap {
		h.FuncMap[name] = f
	}
	return h
}

func (h Template) NewTemplate(tplPath string) (*template.Template, error) {
	tplName := filepath.Base(tplPath)
	return template.New(tplName).Funcs(h.FuncMap).ParseFiles(h.resolveTemplatePath(tplPath))

}

func (h Template) resolveTemplatePath(tplName string) string {
	return path.Clean(h.BaseTemplateDir) + "/" + tplName
}
