package main

import (
	"log"
	"github.com/plausibility/word-by-word/server"
	"github.com/plausibility/word-by-word/controllers"
)

func main() {
	log.SetFlags(log.Ldate|log.Ltime)

	if err := server.Setup(); err != nil {
		panic(err) // programming
	}

	controllers.Bind()

	server.Run()
}
