package main

import (
	_ "github.com/whj1990/bff-mine/docs"
	"github.com/whj1990/bff-mine/internal/router"
	"github.com/whj1990/go-core/launch"
)

func main() {
	logger, closer := launch.InitPremise()
	defer logger.Sync()
	defer closer.Close()

	launch.InitHttpServer(initServer().Routes...)
}

func newRouterQuote(mineHandler *router.MineHandler) *launch.RouterQuote {
	return &launch.RouterQuote{Routes: []launch.HttpRouter{
		mineHandler,
	}}
}
