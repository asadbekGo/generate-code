package client_handler

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_client_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateComing godoc
// @Security ApiKeyAuth
// @ID create_coming
// @Router /v1/coming [POST]
// @Summary Create Coming
// @Description Create Coming
// @Tags Coming
// @Accept json
// @Produce json
// @Param Coming body storehouse_client_service.CreateComingRequest true "CreateComingRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_client_service.Coming} "Coming data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateComing(c *gin.Context) {

	var coming storehouse_client_service.CreateComingRequest
	err := c.ShouldBindJSON(&coming)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().Coming().CreateComing(
		context.Background(),
		&coming,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleComing godoc
// @Security ApiKeyAuth
// @ID get_coming_by_id
// @Router /v1/coming/{coming_id} [GET]
// @Summary Get single Coming
// @Description Get single Coming
// @Tags Coming
// @Accept json
// @Produce json
// @Param coming_id path string true "coming_id"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.Coming} "ComingBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleComing(c *gin.Context) {

	var comingId = c.Param("coming_id")
	if !util.IsValidUUID(comingId) {
		h.HandleResponse(c, status_http.InvalidArgument, "coming id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().Coming().GetByIDComing(
		context.Background(),
		&storehouse_client_service.ComingPrimaryKey{Id: comingId},
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.OK, response)
}

// GetComingList godoc
// @Security ApiKeyAuth
// @ID get_coming_list
// @Router /v1/coming [GET]
// @Summary Get Coming list
// @Description Get Coming list
// @Tags Coming
// @Accept json
// @Produce json
// @Param filters query storehouse_client_service.GetListComingRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.GetListComingResponse} "ComingBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetComingList(c *gin.Context) {

	page, err := h.GetPageParam(c)
	if err != nil {
		h.HandleResponse(c, status_http.InvalidArgument, err.Error())
		return
	}


	limit, err := h.GetLimitParam(c)
	if err != nil {
		h.HandleResponse(c, status_http.InvalidArgument, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().Coming().GetListComing(
		context.Background(),
		&storehouse_client_service.GetListComingRequest{
			Limit:  int32(limit),
			Page:   int32(page),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.OK, response)
}

// UpdateComing godoc
// @Security ApiKeyAuth
// @ID update_coming
// @Router /v1/coming [PUT]
// @Summary Update Coming
// @Description Update Coming
// @Tags Coming
// @Accept json
// @Produce json
// @Param Coming body storehouse_client_service.UpdateComingRequest true "UpdateComingRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.Coming} "Coming data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateComing(c *gin.Context) {

	var updateComing storehouse_client_service.UpdateComingRequest
	err := c.ShouldBindJSON(&updateComing)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().Coming().UpdateComing(
		context.Background(),
		&updateComing,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteComing godoc
// @Security ApiKeyAuth
// @ID delete_coming
// @Router /v1/coming/{coming_id} [DELETE]
// @Summary Delete Coming
// @Description Delete Coming
// @Tags Coming
// @Accept json
// @Produce json
// @Param coming_id path string true "coming_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteComing(c *gin.Context) {

	var comingId = c.Param("coming_id")
	if !util.IsValidUUID(comingId) {
		h.HandleResponse(c, status_http.InvalidArgument, "coming id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().Coming().DeleteComing(
		context.Background(),
		&storehouse_client_service.ComingPrimaryKey{Id: comingId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
