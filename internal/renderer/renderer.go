package renderer

import (
	"html/template"
	"io"

	"github.com/Masterminds/sprig/v3"
	"github.com/bdreece/notable/web"
	"github.com/labstack/echo/v4"
)

type renderer struct {
	*template.Template
}

// Render implements echo.Renderer.
func (r *renderer) Render(w io.Writer, name string, data any, _ echo.Context) error {
    t, err := r.Clone()
    if err != nil {
        return err
    }

    t, err = t.ParseFS(web.Templates, "templates/**/"+name)
    if err != nil {
        return err
    }

    return t.ExecuteTemplate(w, name, data)
}

func New() (echo.Renderer, error) {
    t, err := template.New("").
        Funcs(sprig.FuncMap()).
        ParseFS(web.Templates, "templates/**/*.gotmpl")
    if err != nil {
        return nil, err
    }

	return &renderer{t}, nil
}
