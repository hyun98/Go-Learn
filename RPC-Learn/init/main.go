package main

import (
	"RPC-Learn/config"
	"RPC-Learn/gRPC/server"
	"RPC-Learn/init/app"
	"flag"
	"time"
)

var configFlag = flag.String("config", "./env.toml", "env file not found")

func main() {
	flag.Parse()
	c := config.NewConfig(*configFlag)

	if err := server.NewGRPCServer(c); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	app.NewApp(c)
}
