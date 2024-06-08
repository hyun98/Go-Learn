package main

import (
	"eCommerce/config"
	"eCommerce/init/app"
	"flag"
	"fmt"
)

var envFlag = flag.String("config", "./env.toml", "env file not found")

func main() {
	flag.Parse()
	c := config.NewConfig(*envFlag)
	app.NewApp(c)
	fmt.Println("Hello World")
}
