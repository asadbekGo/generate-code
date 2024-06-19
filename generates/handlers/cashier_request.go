package storehouse_supplier_handler

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_supplier_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateCashierRequest godoc
// @Security ApiKeyAuth
// @ID create_cashierRequest
// @Router /v1/cashier-request [POST]
// @Summary Create CashierRequest
// @Description Create CashierRequest
// @Tags CashierRequest
// @Accept json
// @Produce json
// @Param CashierRequest body storehouse_supplier_service.CreateCashierRequestRequest true "CreateCashierRequestRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_supplier_service.CashierRequest} "CashierRequest data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateCashierRequest(c *gin.Context) {

	var cashierRequest storehouse_supplier_service.CreateCashierRequestRequest
	err := c.ShouldBindJSON(&cashierRequest)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseSupplierService().CashierRequest().CreateCashierRequest(
		context.Background(),
		&cashierRequest,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleCashierRequest godoc
// @Security ApiKeyAuth
// @ID get_cashierRequest_by_id
// @Router /v1/cashier-request/{cashier_request_id} [GET]
// @Summary Get single CashierRequest
// @Description Get single CashierRequest
// @Tags CashierRequest
// @Accept json
// @Produce json
// @Param cashier_request_id path string true "cashier_request_id"
// @Success 200 {object} status_http.Response{data=storehouse_supplier_service.CashierRequest} "CashierRequestBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleCashierRequest(c *gin.Context) {

	var cashierRequestId = c.Param("cashier_request_id")
	if !util.IsValidUUID(cashierRequestId) {
		h.HandleResponse(c, status_http.InvalidArgument, "cashierRequest id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseSupplierService().CashierRequest().GetByIDCashierRequest(
		context.Background(),
		&storehouse_supplier_service.CashierRequestPrimaryKey{Id: cashierRequestId},
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

// GetCashierRequestList godoc
// @Security ApiKeyAuth
// @ID get_cashierRequest_list
// @Router /v1/cashier-request [GET]
// @Summary Get CashierRequest list
// @Description Get CashierRequest list
// @Tags CashierRequest
// @Accept json
// @Produce json
// @Param filters query storehouse_supplier_service.GetListCashierRequestRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_supplier_service.GetListCashierRequestResponse} "CashierRequestBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetCashierRequestList(c *gin.Context) {

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

	response, err := h.services.StorehouseSupplierService().CashierRequest().GetListCashierRequest(
		context.Background(),
		&storehouse_supplier_service.GetListCashierRequestRequest{
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

// UpdateCashierRequest godoc
// @Security ApiKeyAuth
// @ID update_cashierRequest
// @Router /v1/cashier-request [PUT]
// @Summary Update CashierRequest
// @Description Update CashierRequest
// @Tags CashierRequest
// @Accept json
// @Produce json
// @Param CashierRequest body storehouse_supplier_service.UpdateCashierRequestRequest true "UpdateCashierRequestRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_supplier_service.CashierRequest} "CashierRequest data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateCashierRequest(c *gin.Context) {

	var updateCashierRequest storehouse_supplier_service.UpdateCashierRequestRequest
	err := c.ShouldBindJSON(&updateCashierRequest)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseSupplierService().CashierRequest().UpdateCashierRequest(
		context.Background(),
		&updateCashierRequest,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteCashierRequest godoc
// @Security ApiKeyAuth
// @ID delete_cashierRequest
// @Router /v1/cashier-request/{cashier_request_id} [DELETE]
// @Summary Delete CashierRequest
// @Description Delete CashierRequest
// @Tags CashierRequest
// @Accept json
// @Produce json
// @Param cashier_request_id path string true "cashier_request_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteCashierRequest(c *gin.Context) {

	var cashierRequestId = c.Param("cashier_request_id")
	if !util.IsValidUUID(cashierRequestId) {
		h.HandleResponse(c, status_http.InvalidArgument, "cashierRequest id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseSupplierService().CashierRequest().DeleteCashierRequest(
		context.Background(),
		&storehouse_supplier_service.CashierRequestPrimaryKey{Id: cashierRequestId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
