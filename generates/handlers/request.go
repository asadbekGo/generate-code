package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateRequest godoc
// @Security ApiKeyAuth
// @ID create_request
// @Router /request [POST]
// @Summary Create Request
// @Description Create Request
// @Tags Request
// @Accept json
// @Produce json
// @Param Request body storehouse_service.CreateRequestRequest true "CreateRequestRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.Request} "Request data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateRequest(c *gin.Context) {

	var request storehouse_service.CreateRequestRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Request().CreateRequest(
		context.Background(),
		&request,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleRequest godoc
// @Security ApiKeyAuth
// @ID get_request_by_id
// @Router /request/{request_id} [GET]
// @Summary Get single Request
// @Description Get single Request
// @Tags Request
// @Accept json
// @Produce json
// @Param request_id path string true "request_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.Request} "RequestBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleRequest(c *gin.Context) {

	var requestId = c.Param("request_id")
	if !util.IsValidUUID(requestId) {
		h.HandleResponse(c, status_http.InvalidArgument, "request id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Request().GetByIDRequest(
		context.Background(),
		&storehouse_service.RequestPrimaryKey{Id: requestId},
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

// GetRequestList godoc
// @Security ApiKeyAuth
// @ID get_request_list
// @Router /request [GET]
// @Summary Get Request list
// @Description Get Request list
// @Tags Request
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListRequestRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListRequestResponse} "RequestBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetRequestList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().Request().GetListRequest(
		context.Background(),
		&storehouse_service.GetListRequestRequest{
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

// UpdateRequest godoc
// @Security ApiKeyAuth
// @ID update_request
// @Router /request [PUT]
// @Summary Update Request
// @Description Update Request
// @Tags Request
// @Accept json
// @Produce json
// @Param Request body storehouse_service.UpdateRequestRequest true "UpdateRequestRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.Request} "Request data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateRequest(c *gin.Context) {

	var updateRequest storehouse_service.UpdateRequestRequest
	err := c.ShouldBindJSON(&updateRequest)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Request().UpdateRequest(
		context.Background(),
		&updateRequest,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteRequest godoc
// @Security ApiKeyAuth
// @ID delete_request
// @Router /request/{request_id} [DELETE]
// @Summary Delete Request
// @Description Delete Request
// @Tags Request
// @Accept json
// @Produce json
// @Param request_id path string true "request_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteRequest(c *gin.Context) {

	var requestId = c.Param("request_id")
	if !util.IsValidUUID(requestId) {
		h.HandleResponse(c, status_http.InvalidArgument, "request id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Request().DeleteRequest(
		context.Background(),
		&storehouse_service.RequestPrimaryKey{Id: requestId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
