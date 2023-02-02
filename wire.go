//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/whj1990/bff-mine/internal/handler"
	"github.com/whj1990/bff-mine/internal/router"
	"github.com/whj1990/go-core/launch"
)

func initServer() *launch.RouterQuote {
	panic(wire.Build(router.ProviderSet, handler.ProviderSet, newRouterQuote))
}
