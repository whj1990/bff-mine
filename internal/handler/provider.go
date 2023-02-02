package handler

import (
	"github.com/google/wire"
	"github.com/whj1990/app-mine/kitex_gen/api/appmine"
	"github.com/whj1990/app-mine/rpc"
)

var ProviderSet = wire.NewSet(NewMineHandler, rpc.NewMineClient)

func NewMineHandler(client appmine.Client) MineHandler {
	return &minelHandler{client}
}
