package server

import (
	"fmt"
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	t := &Template{}
	err := t.initTemplates()
	if err != nil {
		fmt.Println(err)
		panic("couldn't init templates")
	}
	return t
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (h *Template) initTemplates() error {
	t, err := parseTemplates()
	if err != nil {
		return err
	}
	h.templates = t
	return nil
}

func parseTemplates() (*template.Template, error) {
	t, err := template.ParseGlob("web/templates/base/*")
	if err != nil {

		return nil, fmt.Errorf("base templates: %v", err)
	}

	t, err = t.ParseGlob("web/components/*")
	if err != nil {

		return nil, fmt.Errorf("component templates: %v", err)
	}

	t, err = t.ParseGlob("web/templates/pages/*")
	if err != nil {

		return nil, fmt.Errorf("page templates: %v", err)
	}

	return t, nil

}

// func (h *Handler) executeTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
// 	w.Header().Set("Content-Type", "text/html")
// 	w.WriteHeader(200)

// 	err := h.Templates.ExecuteTemplate(w, tmpl, data)
// 	if err != nil {
// 		fmt.Printf("execute template: %v\n", err)
// 		print404(w)
// 		return
// 	}
// }
