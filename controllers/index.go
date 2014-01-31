package controllers

import (
	"github.com/hoisie/web"
	"github.com/plausibility/word-by-word/util"
)

func index(ctx *web.Context) {
	items := make(map[string]interface{})
	util.Render(ctx, "index.html", items)
}
