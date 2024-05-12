package core

import (
	"html/template"
	"path"
)

func NewTemplate(baseTemplateDir string) (*template.Template, error) {
	template, err := template.ParseGlob(path.Join(baseTemplateDir, "**/*.tpl"))
	if err != nil {
		return nil, err
	}
	return template, nil
}
