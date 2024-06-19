package storehouse_supplier_handler

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_supplier_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateCashierRequestProduct godoc
// @Security ApiKeyAuth
// @ID create_cashierRequestProduct
// @Router /v1/cashier-request-product [POST]
// @Summary Create CashierRequestProduct
// @Description Create CashierRequestProduct
// @Tags CashierRequestProduct
// @Accept json
// @Produce json
// @Param CashierRequestProduct body storehouse_supplier_service.CreateCashierRequestProductRequest true "CreateCashierRequestProductRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_supplier_service.CashierRequestProduct} "CashierRequestProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateCashierRequestProduct(c *gin.Context) {

	var cashierRequestProduct storehouse_supplier_service.CreateCashierRequestProductRequest
	err := c.ShouldBindJSON(&cashierRequestProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseSupplierService().CashierRequestProduct().CreateCashierRequestProduct(
		context.Background(),
		&cashierRequestProduct,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleCashierRequestProduct godoc
// @Security ApiKeyAuth
// @ID get_cashierRequestProduct_by_id
// @Router /v1/cashier-request-product/{cashier_request_product_id} [GET]
// @Summary Get single CashierRequestProduct
// @Description Get single CashierRequestProduct
// @Tags CashierRequestProduct
// @Accept json
// @Produce json
// @Param cashier_request_product_id path string true "cashier_request_product_id"
// @Success 200 {object} status_http.Response{data=storehouse_supplier_service.CashierRequestProduct} "CashierRequestProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleCashierRequestProduct(c *gin.Context) {

	var cashierRequestProductId = c.Param("cashier_request_product_id")
	if !util.IsValidUUID(cashierRequestProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "cashierRequestProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseSupplierService().CashierRequestProduct().GetByIDCashierRequestProduct(
		context.Background(),
		&storehouse_supplier_service.CashierRequestProductPrimaryKey{Id: cashierRequestProductId},
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

// GetCashierRequestProductList godoc
// @Security ApiKeyAuth
// @ID get_cashierRequestProduct_list
// @Router /v1/cashier-request-product [GET]
// @Summary Get CashierRequestProduct list
// @Description Get CashierRequestProduct list
// @Tags CashierRequestProduct
// @Accept json
// @Produce json
// @Param filters query storehouse_supplier_service.GetListCashierRequestProductRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_supplier_service.GetListCashierRequestProductResponse} "CashierRequestProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetCashierRequestProductList(c *gin.Context) {

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

	response, err := h.services.StorehouseSupplierService().CashierRequestProduct().GetListCashierRequestProduct(
		context.Background(),
		&storehouse_supplier_service.GetListCashierRequestProductRequest{
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

// UpdateCashierRequestProduct godoc
// @Security ApiKeyAuth
// @ID update_cashierRequestProduct
// @Router /v1/cashier-request-product [PUT]
// @Summary Update CashierRequestProduct
// @Description Update CashierRequestProduct
// @Tags CashierRequestProduct
// @Accept json
// @Produce json
// @Param CashierRequestProduct body storehouse_supplier_service.UpdateCashierRequestProductRequest true "UpdateCashierRequestProductRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_supplier_service.CashierRequestProduct} "CashierRequestProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateCashierRequestProduct(c *gin.Context) {

	var updateCashierRequestProduct storehouse_supplier_service.UpdateCashierRequestProductRequest
	err := c.ShouldBindJSON(&updateCashierRequestProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseSupplierService().CashierRequestProduct().UpdateCashierRequestProduct(
		context.Background(),
		&updateCashierRequestProduct,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteCashierRequestProduct godoc
// @Security ApiKeyAuth
// @ID delete_cashierRequestProduct
// @Router /v1/cashier-request-product/{cashier_request_product_id} [DELETE]
// @Summary Delete CashierRequestProduct
// @Description Delete CashierRequestProduct
// @Tags CashierRequestProduct
// @Accept json
// @Produce json
// @Param cashier_request_product_id path string true "cashier_request_product_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteCashierRequestProduct(c *gin.Context) {

	var cashierRequestProductId = c.Param("cashier_request_product_id")
	if !util.IsValidUUID(cashierRequestProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "cashierRequestProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseSupplierService().CashierRequestProduct().DeleteCashierRequestProduct(
		context.Background(),
		&storehouse_supplier_service.CashierRequestProductPrimaryKey{Id: cashierRequestProductId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
