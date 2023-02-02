package router

import (
	"github.com/google/wire"
	"github.com/whj1990/bff-mine/internal/handler"
)

var ProviderSet = wire.NewSet(NewMinelRouter)

func NewMinelRouter(recVisualHandler handler.MineHandler) *RecVisualHandler {
	return &RecVisualHandler{recVisualHandler}
}
