//go:build wireinject
// +build wireinject

package main

import (
	"git.dp.ibbtv.cn/dp-backend/bff-rec-visual/internal/handler"
	"git.dp.ibbtv.cn/dp-backend/bff-rec-visual/internal/router"
	"git.dp.ibbtv.cn/dp-backend/go-core/launch"
	"github.com/google/wire"
)

func initServer() *launch.RouterQuote {
	panic(wire.Build(router.ProviderSet, handler.ProviderSet, newRouterQuote))
}
