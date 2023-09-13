package main

import (
	"context"
	"flag"
	"github.com/huerni/miao-gateway/internal/app"
	"github.com/huerni/miao-gateway/internal/config"
)

var configFile = flag.String("f", "etc/cfg.toml", "the config file")

func main() {
	c, err := config.InitConfig(*configFile)
	if err != nil {
		panic(err)
	}

	g := app.NewGmServer(c)
	g.Start(context.Background())
	g.WaitForShutdown(context.Background())
}
