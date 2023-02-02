package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/whj1990/app-mine/internal/handler/stru"
	"github.com/whj1990/app-mine/kitex_gen/api"
	"github.com/whj1990/app-mine/kitex_gen/api/appine"
	"github.com/whj1990/go-core/handler"
	"github.com/whj1990/go-core/util"
)

type MineHandler interface {
	ReviewProjectList() gin.HandlerFunc
	ReviewProjectDetails() gin.HandlerFunc
	ReviewProjectSave() gin.HandlerFunc
	ReviewProjectStatus() gin.HandlerFunc
	ReviewProjectDelete() gin.HandlerFunc
}
type mineHandler struct {
	appineClient appine.Client
}

// @Summary 获取项目列表
// @Tags review
// @Produce json
// @Param authorization header string true "token"
// @Param query query stru.ReviewProjectListParams false "query"
// @Success 200 {object} handler.response{data=api.ReviewProjectListResponse} "ok"
// @Router /v1/recvisual/review/project/list [GET]
func (r *mineHandler) ReviewProjectList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req stru.ReviewProjectListParams
		if err := ctx.ShouldBindQuery(&req); err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		res, err := r.appineClient.ReviewProjectList(util.GetRPCContext(ctx), &api.ReviewProjectListParams{
			PageNum:    req.PageNum,
			PageSize:   req.PageSize,
			Id:         req.Id,
			Name:       req.Name,
			ModeCode:   req.ModeCode,
			Status:     req.Status,
			ShowStatus: req.ShowStatus,
		})
		if err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		handler.HandleSuccessDataResponse(ctx, res)
	}
}

// @Summary 获取项目详情
// @Tags review
// @Produce json
// @Param authorization header string true "token"
// @Param query query stru.IdsInt64Params true "query"
// @Success 200 {object} handler.response{data=api.ReviewProjectDetailResponse} "ok"
// @Router /v1/recvisual/review/project/detail [GET]
func (r *mineHandler) ReviewProjectDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req stru.IdsInt64Params
		if err := ctx.ShouldBindQuery(&req); err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		res, err := r.appineClient.ReviewProjectDetails(util.GetRPCContext(ctx), &api.IdsInt64Params{
			req.Ids,
		})
		if err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		handler.HandleSuccessDataResponse(ctx, res)
	}
}

// @Summary 项目信息存储
// @Tags review
// @Produce json
// @Param authorization header string true "token"
// @Param body body stru.ReviewProjectSaveParam true "body"
// @Success 200 {object} handler.response{data=api.SaveResponse} "ok"
// @Router /v1/recvisual/review/project/save [POST]
func (r *mineHandler) ReviewProjectSave() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req stru.ReviewProjectSaveParam
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		validate := validator.New()
		err := validate.Struct(req)
		if err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		res, err := r.appineClient.ReviewProjectSave(util.GetRPCContext(ctx), &api.ReviewProjectSaveParam{
			Action: req.Action,
			Data: &api.ReviewProjectSaveData{
				Id:          req.Data.Id,
				Name:        req.Data.Name,
				Description: req.Data.Description,
				ModeCode:    req.Data.ModeCode,
			},
		})
		if err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		handler.HandleSuccessDataResponse(ctx, res)
	}
}

// @Summary 项目状态更新
// @Tags review
// @Produce json
// @Param authorization header string true "token"
// @Param body body stru.StatusParam true "body"
// @Success 200 {object} handler.response{data=api.SaveResponse} "ok"
// @Router /v1/recvisual/review/project/upStatus [POST]
func (r *mineHandler) ReviewProjectStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req stru.StatusParam
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		res, err := r.appineClient.ReviewProjectStatus(util.GetRPCContext(ctx), &api.StatusParam{
			Id:     req.Id,
			Status: req.Status,
		})
		if err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		handler.HandleSuccessDataResponse(ctx, res)
	}
}

// @Summary 项目信息删除
// @Tags review
// @Produce json
// @Param authorization header string true "token"
// @Param body body stru.IdParam true "body"
// @Success 200 {object} handler.response{data=api.SaveResponse} "ok"
// @Router /v1/recvisual/review/project/delete [POST]
func (r *mineHandler) ReviewProjectDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req stru.IdParam
		if err := ctx.ShouldBindJSON(&req); err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		res, err := r.appineClient.ReviewProjectDelete(util.GetRPCContext(ctx), &api.IdParam{
			Id: req.Id,
		})
		if err != nil {
			handler.HandleErrorResponse(ctx, err)
			return
		}
		handler.HandleSuccessDataResponse(ctx, res)
	}
}
