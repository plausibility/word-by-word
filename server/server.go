package server

import (
	"os"
	"fmt"
	"flag"
	"path"
	"github.com/hoisie/web"
)

var Server *web.Server
var bindTo string
var StaticDir string
var TemplateDir string

func Setup() (err error) {
	host := flag.String("host", "0.0.0.0", "host to bind web.go to")
	port := flag.Int("port", 9999, "port to bind web.go to")

	cwd, err := os.Getwd()
	static_path := path.Join(cwd, "static")
	if err != nil {
		return err
	}
	static := flag.String("static", static_path, "static web directory")

	flag.Parse()

	if !path.IsAbs(*static) {
		*static = path.Join(cwd, *static)
	}

	bindTo = fmt.Sprintf("%s:%d", *host, *port)
	StaticDir = *static
	TemplateDir = path.Join(StaticDir, "templates")

	// web.go-y stuff.
	Server = web.NewServer()
	return nil
}

func Run() {
	Server.Run(bindTo)
}
