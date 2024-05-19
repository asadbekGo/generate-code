package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateChapterStock godoc
// @Security ApiKeyAuth
// @ID create_chapterStock
// @Router /chapter-stock [POST]
// @Summary Create ChapterStock
// @Description Create ChapterStock
// @Tags ChapterStock
// @Accept json
// @Produce json
// @Param ChapterStock body storehouse_service.CreateChapterStockRequest true "CreateChapterStockRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.ChapterStock} "ChapterStock data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateChapterStock(c *gin.Context) {

	var chapterStock storehouse_service.CreateChapterStockRequest
	err := c.ShouldBindJSON(&chapterStock)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().ChapterStock().CreateChapterStock(
		context.Background(),
		&chapterStock,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleChapterStock godoc
// @Security ApiKeyAuth
// @ID get_chapterStock_by_id
// @Router /chapter-stock/{chapter_stock_id} [GET]
// @Summary Get single ChapterStock
// @Description Get single ChapterStock
// @Tags ChapterStock
// @Accept json
// @Produce json
// @Param chapter_stock_id path string true "chapter_stock_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.ChapterStock} "ChapterStockBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleChapterStock(c *gin.Context) {

	var chapterStockId = c.Param("chapter_stock_id")
	if !util.IsValidUUID(chapterStockId) {
		h.HandleResponse(c, status_http.InvalidArgument, "chapterStock id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().ChapterStock().GetByIDChapterStock(
		context.Background(),
		&storehouse_service.ChapterStockPrimaryKey{Id: chapterStockId},
	)
	if response == nil {
		err := errors.New("not Found")
		h.HandleResponse(c, status_http.NoContent, err.Error())
		return
	}
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.OK, response)
}

// GetChapterStockList godoc
// @Security ApiKeyAuth
// @ID get_chapterStock_list
// @Router /chapter-stock [GET]
// @Summary Get ChapterStock list
// @Description Get ChapterStock list
// @Tags ChapterStock
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListChapterStockRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListChapterStockResponse} "ChapterStockBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetChapterStockList(c *gin.Context) {

	offset, err := h.GetOffsetParam(c)
	if err != nil {
		h.HandleResponse(c, status_http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.GetLimitParam(c)
	if err != nil {
		h.HandleResponse(c, status_http.InvalidArgument, err.Error())
		return
	}

	response, err := h.services.StorehouseService().ChapterStock().GetListChapterStock(
		context.Background(),
		&storehouse_service.GetListChapterStockRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.OK, response)
}

// UpdateChapterStock godoc
// @Security ApiKeyAuth
// @ID update_chapterStock
// @Router /chapter-stock [PUT]
// @Summary Update ChapterStock
// @Description Update ChapterStock
// @Tags ChapterStock
// @Accept json
// @Produce json
// @Param ChapterStock body storehouse_service.UpdateChapterStockRequest true "UpdateChapterStockRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.ChapterStock} "ChapterStock data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateChapterStock(c *gin.Context) {

	var updateChapterStock storehouse_service.UpdateChapterStockRequest
	err := c.ShouldBindJSON(&updateChapterStock)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().ChapterStock().UpdateChapterStock(
		context.Background(),
		&updateChapterStock,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteChapterStock godoc
// @Security ApiKeyAuth
// @ID delete_chapterStock
// @Router /chapter-stock/{chapter_stock_id} [DELETE]
// @Summary Delete ChapterStock
// @Description Delete ChapterStock
// @Tags ChapterStock
// @Accept json
// @Produce json
// @Param chapter_stock_id path string true "chapter_stock_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteChapterStock(c *gin.Context) {

	var chapterStockId = c.Param("chapter_stock_id")
	if !util.IsValidUUID(chapterStockId) {
		h.HandleResponse(c, status_http.InvalidArgument, "chapterStock id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().ChapterStock().DeleteChapterStock(
		context.Background(),
		&storehouse_service.ChapterStockPrimaryKey{Id: chapterStockId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
