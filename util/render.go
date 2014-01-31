package util

import (
	"fmt"
	"path"
	"html/template"
	"github.com/hoisie/web"
	"github.com/plausibility/word-by-word/server"
)

func Render(ctx *web.Context, templ string, items map[string]interface{}) {
	// TODO: cache template.ParseFiles
	tpl_path := path.Join(server.TemplateDir, templ)
	tpl, err := template.ParseFiles(tpl_path)
	if err != nil {
		// TODO: logging errors
		ctx.WriteHeader(500)
		fmt.Fprintf(ctx, "Couldn't parse template: %s", err)
		return
	}
	tpl.Execute(ctx, items)
	return
}
