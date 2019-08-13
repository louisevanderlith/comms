package core

import (
	"html/template"
	"path"
	"strings"
)

func PopulatTemplate(msg Message) (string, error) {
	tmplName := msg.TemplateName

	if len(tmplName) == 0 {
		tmplName = "default.html"
	}

	files := []string{
		path.Join("templates", "base.html"),
		path.Join("templates", tmplName),
		path.Join("templates", "style.html"),
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		return "", err
	}

	var strBuild strings.Builder
	err = tmpl.Execute(&strBuild, msg)

	if err != nil {
		return "", err
	}

	return strBuild.String(), nil
}
