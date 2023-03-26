package main

import (
	"web-app/internal/cli"
	"web-app/internal/conf"
	"web-app/internal/server"
)

func main() {
	cli.Parse()
	server.Start(conf.NewConfig())
}
