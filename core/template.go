package core

import (
	"bufio"
	"bytes"
	"html/template"
	"path"
)

func populatTemplate(msg Message) (string, error) {
	tmplName := msg.TemplateName

	if len(tmplName) == 0 {
		tmplName = "base.html"
	}

	files := []string{
		path.Join("templates", "base.html"),
		path.Join("templates", msg.TemplateName),
	}

	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err = tmpl.Execute(w, msg)

	if err != nil {
		return "", err
	}

	return b.String(), nil
}
