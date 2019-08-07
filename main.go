package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo"

	"go-echo-test/handler"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		err := errors.New("Template not foun -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	e := echo.New()

	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("view/home.html", "view/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("view/about.html", "view/base.html"))

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.GET("/", handler.HomeHandler)
	e.GET("/about", handler.AboutHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
