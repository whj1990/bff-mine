package router

import (
	"github.com/gin-gonic/gin"
	"github.com/whj1990/bff-mine/internal/handler"
	"github.com/whj1990/go-common/constant"
)

type RecVisualHandler struct {
	mineHandler handler.MineHandler
}

func (r *RecVisualHandler) SetRouter(app *gin.Engine) {
	group := app.Group(constant.MineRouterGroup)
	group.GET("/review/project/list", r.mineHandler.ReviewProjectList())
	group.GET("/review/project/detail", r.mineHandler.ReviewProjectDetails())
	group.POST("/review/project/save", r.mineHandler.ReviewProjectSave())
	group.POST("/review/project/upStatus", r.mineHandler.ReviewProjectStatus())
	group.POST("/review/project/delete", r.mineHandler.ReviewProjectDelete())
}
