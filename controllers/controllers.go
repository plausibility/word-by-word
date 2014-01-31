package controllers

import "github.com/plausibility/word-by-word/server"

func Bind() {
	server.Server.Get("/", index)
}
