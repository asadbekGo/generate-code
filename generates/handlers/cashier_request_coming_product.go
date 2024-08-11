package client_handler

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_client_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateCashierRequestComingProduct godoc
// @Security ApiKeyAuth
// @ID create_cashierRequestComingProduct
// @Router /v1/cashier-request-coming-product [POST]
// @Summary Create CashierRequestComingProduct
// @Description Create CashierRequestComingProduct
// @Tags CashierRequestComingProduct
// @Accept json
// @Produce json
// @Param CashierRequestComingProduct body storehouse_client_service.CreateCashierRequestComingProductRequest true "CreateCashierRequestComingProductRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_client_service.CashierRequestComingProduct} "CashierRequestComingProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateCashierRequestComingProduct(c *gin.Context) {

	var cashierRequestComingProduct storehouse_client_service.CreateCashierRequestComingProductRequest
	err := c.ShouldBindJSON(&cashierRequestComingProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().CashierRequestComingProduct().CreateCashierRequestComingProduct(
		context.Background(),
		&cashierRequestComingProduct,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleCashierRequestComingProduct godoc
// @Security ApiKeyAuth
// @ID get_cashierRequestComingProduct_by_id
// @Router /v1/cashier-request-coming-product/{cashier_request_coming_product_id} [GET]
// @Summary Get single CashierRequestComingProduct
// @Description Get single CashierRequestComingProduct
// @Tags CashierRequestComingProduct
// @Accept json
// @Produce json
// @Param cashier_request_coming_product_id path string true "cashier_request_coming_product_id"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.CashierRequestComingProduct} "CashierRequestComingProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleCashierRequestComingProduct(c *gin.Context) {

	var cashierRequestComingProductId = c.Param("cashier_request_coming_product_id")
	if !util.IsValidUUID(cashierRequestComingProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "cashierRequestComingProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().CashierRequestComingProduct().GetByIDCashierRequestComingProduct(
		context.Background(),
		&storehouse_client_service.CashierRequestComingProductPrimaryKey{Id: cashierRequestComingProductId},
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.OK, response)
}

// GetCashierRequestComingProductList godoc
// @Security ApiKeyAuth
// @ID get_cashierRequestComingProduct_list
// @Router /v1/cashier-request-coming-product [GET]
// @Summary Get CashierRequestComingProduct list
// @Description Get CashierRequestComingProduct list
// @Tags CashierRequestComingProduct
// @Accept json
// @Produce json
// @Param filters query storehouse_client_service.GetListCashierRequestComingProductRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.GetListCashierRequestComingProductResponse} "CashierRequestComingProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetCashierRequestComingProductList(c *gin.Context) {

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

	response, err := h.services.StorehouseClientService().CashierRequestComingProduct().GetListCashierRequestComingProduct(
		context.Background(),
		&storehouse_client_service.GetListCashierRequestComingProductRequest{
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

// UpdateCashierRequestComingProduct godoc
// @Security ApiKeyAuth
// @ID update_cashierRequestComingProduct
// @Router /v1/cashier-request-coming-product [PUT]
// @Summary Update CashierRequestComingProduct
// @Description Update CashierRequestComingProduct
// @Tags CashierRequestComingProduct
// @Accept json
// @Produce json
// @Param CashierRequestComingProduct body storehouse_client_service.UpdateCashierRequestComingProductRequest true "UpdateCashierRequestComingProductRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.CashierRequestComingProduct} "CashierRequestComingProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateCashierRequestComingProduct(c *gin.Context) {

	var updateCashierRequestComingProduct storehouse_client_service.UpdateCashierRequestComingProductRequest
	err := c.ShouldBindJSON(&updateCashierRequestComingProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().CashierRequestComingProduct().UpdateCashierRequestComingProduct(
		context.Background(),
		&updateCashierRequestComingProduct,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteCashierRequestComingProduct godoc
// @Security ApiKeyAuth
// @ID delete_cashierRequestComingProduct
// @Router /v1/cashier-request-coming-product/{cashier_request_coming_product_id} [DELETE]
// @Summary Delete CashierRequestComingProduct
// @Description Delete CashierRequestComingProduct
// @Tags CashierRequestComingProduct
// @Accept json
// @Produce json
// @Param cashier_request_coming_product_id path string true "cashier_request_coming_product_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteCashierRequestComingProduct(c *gin.Context) {

	var cashierRequestComingProductId = c.Param("cashier_request_coming_product_id")
	if !util.IsValidUUID(cashierRequestComingProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "cashierRequestComingProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().CashierRequestComingProduct().DeleteCashierRequestComingProduct(
		context.Background(),
		&storehouse_client_service.CashierRequestComingProductPrimaryKey{Id: cashierRequestComingProductId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
