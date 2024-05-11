package core

import (
	"html/template"
	"path"
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

func (h Template) NewTemplate(tplPaths ...string) (*template.Template, error) {
	return template.ParseFiles(h.resolveTemplatePaths(tplPaths)...)

}

func (h Template) resolveTemplatePath(tplPath string) string {
	return path.Clean(h.BaseTemplateDir) + "/" + tplPath
}

func (h Template) resolveTemplatePaths(tplPaths []string) []string {
	paths := []string{}
	for _, path := range tplPaths {
		paths = append(paths, h.resolveTemplatePath(path))
	}
	return paths
}
