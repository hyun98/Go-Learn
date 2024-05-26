package main

import (
	"CRUD_SERVER/init/cmd"
	"flag"
)

var configPath = flag.String("config", "./init/config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPath)
}
