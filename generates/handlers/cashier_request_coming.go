package client_handler

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_client_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateCashierRequestComing godoc
// @Security ApiKeyAuth
// @ID create_cashierRequestComing
// @Router /v1/cashier-request-coming [POST]
// @Summary Create CashierRequestComing
// @Description Create CashierRequestComing
// @Tags CashierRequestComing
// @Accept json
// @Produce json
// @Param CashierRequestComing body storehouse_client_service.CreateCashierRequestComingRequest true "CreateCashierRequestComingRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_client_service.CashierRequestComing} "CashierRequestComing data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateCashierRequestComing(c *gin.Context) {

	var cashierRequestComing storehouse_client_service.CreateCashierRequestComingRequest
	err := c.ShouldBindJSON(&cashierRequestComing)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().CashierRequestComing().CreateCashierRequestComing(
		context.Background(),
		&cashierRequestComing,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleCashierRequestComing godoc
// @Security ApiKeyAuth
// @ID get_cashierRequestComing_by_id
// @Router /v1/cashier-request-coming/{cashier_request_coming_id} [GET]
// @Summary Get single CashierRequestComing
// @Description Get single CashierRequestComing
// @Tags CashierRequestComing
// @Accept json
// @Produce json
// @Param cashier_request_coming_id path string true "cashier_request_coming_id"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.CashierRequestComing} "CashierRequestComingBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleCashierRequestComing(c *gin.Context) {

	var cashierRequestComingId = c.Param("cashier_request_coming_id")
	if !util.IsValidUUID(cashierRequestComingId) {
		h.HandleResponse(c, status_http.InvalidArgument, "cashierRequestComing id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().CashierRequestComing().GetByIDCashierRequestComing(
		context.Background(),
		&storehouse_client_service.CashierRequestComingPrimaryKey{Id: cashierRequestComingId},
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.OK, response)
}

// GetCashierRequestComingList godoc
// @Security ApiKeyAuth
// @ID get_cashierRequestComing_list
// @Router /v1/cashier-request-coming [GET]
// @Summary Get CashierRequestComing list
// @Description Get CashierRequestComing list
// @Tags CashierRequestComing
// @Accept json
// @Produce json
// @Param filters query storehouse_client_service.GetListCashierRequestComingRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.GetListCashierRequestComingResponse} "CashierRequestComingBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetCashierRequestComingList(c *gin.Context) {

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

	response, err := h.services.StorehouseClientService().CashierRequestComing().GetListCashierRequestComing(
		context.Background(),
		&storehouse_client_service.GetListCashierRequestComingRequest{
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

// UpdateCashierRequestComing godoc
// @Security ApiKeyAuth
// @ID update_cashierRequestComing
// @Router /v1/cashier-request-coming [PUT]
// @Summary Update CashierRequestComing
// @Description Update CashierRequestComing
// @Tags CashierRequestComing
// @Accept json
// @Produce json
// @Param CashierRequestComing body storehouse_client_service.UpdateCashierRequestComingRequest true "UpdateCashierRequestComingRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.CashierRequestComing} "CashierRequestComing data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateCashierRequestComing(c *gin.Context) {

	var updateCashierRequestComing storehouse_client_service.UpdateCashierRequestComingRequest
	err := c.ShouldBindJSON(&updateCashierRequestComing)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().CashierRequestComing().UpdateCashierRequestComing(
		context.Background(),
		&updateCashierRequestComing,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteCashierRequestComing godoc
// @Security ApiKeyAuth
// @ID delete_cashierRequestComing
// @Router /v1/cashier-request-coming/{cashier_request_coming_id} [DELETE]
// @Summary Delete CashierRequestComing
// @Description Delete CashierRequestComing
// @Tags CashierRequestComing
// @Accept json
// @Produce json
// @Param cashier_request_coming_id path string true "cashier_request_coming_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteCashierRequestComing(c *gin.Context) {

	var cashierRequestComingId = c.Param("cashier_request_coming_id")
	if !util.IsValidUUID(cashierRequestComingId) {
		h.HandleResponse(c, status_http.InvalidArgument, "cashierRequestComing id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().CashierRequestComing().DeleteCashierRequestComing(
		context.Background(),
		&storehouse_client_service.CashierRequestComingPrimaryKey{Id: cashierRequestComingId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
