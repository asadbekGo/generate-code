package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateStockProduct godoc
// @Security ApiKeyAuth
// @ID create_stockProduct
// @Router /stock-product [POST]
// @Summary Create StockProduct
// @Description Create StockProduct
// @Tags StockProduct
// @Accept json
// @Produce json
// @Param StockProduct body storehouse_service.CreateStockProductRequest true "CreateStockProductRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.StockProduct} "StockProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateStockProduct(c *gin.Context) {

	var stockProduct storehouse_service.CreateStockProductRequest
	err := c.ShouldBindJSON(&stockProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().StockProduct().CreateStockProduct(
		context.Background(),
		&stockProduct,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleStockProduct godoc
// @Security ApiKeyAuth
// @ID get_stockProduct_by_id
// @Router /stock-product/{stock_product_id} [GET]
// @Summary Get single StockProduct
// @Description Get single StockProduct
// @Tags StockProduct
// @Accept json
// @Produce json
// @Param stock_product_id path string true "stock_product_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.StockProduct} "StockProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleStockProduct(c *gin.Context) {

	var stockProductId = c.Param("stock_product_id")
	if !util.IsValidUUID(stockProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "stockProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().StockProduct().GetByIDStockProduct(
		context.Background(),
		&storehouse_service.StockProductPrimaryKey{Id: stockProductId},
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

// GetStockProductList godoc
// @Security ApiKeyAuth
// @ID get_stockProduct_list
// @Router /stock-product [GET]
// @Summary Get StockProduct list
// @Description Get StockProduct list
// @Tags StockProduct
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListStockProductRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListStockProductResponse} "StockProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetStockProductList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().StockProduct().GetListStockProduct(
		context.Background(),
		&storehouse_service.GetListStockProductRequest{
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

// UpdateStockProduct godoc
// @Security ApiKeyAuth
// @ID update_stockProduct
// @Router /stock-product [PUT]
// @Summary Update StockProduct
// @Description Update StockProduct
// @Tags StockProduct
// @Accept json
// @Produce json
// @Param StockProduct body storehouse_service.UpdateStockProductRequest true "UpdateStockProductRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.StockProduct} "StockProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateStockProduct(c *gin.Context) {

	var updateStockProduct storehouse_service.UpdateStockProductRequest
	err := c.ShouldBindJSON(&updateStockProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().StockProduct().UpdateStockProduct(
		context.Background(),
		&updateStockProduct,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteStockProduct godoc
// @Security ApiKeyAuth
// @ID delete_stockProduct
// @Router /stock-product/{stock_product_id} [DELETE]
// @Summary Delete StockProduct
// @Description Delete StockProduct
// @Tags StockProduct
// @Accept json
// @Produce json
// @Param stock_product_id path string true "stock_product_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteStockProduct(c *gin.Context) {

	var stockProductId = c.Param("stock_product_id")
	if !util.IsValidUUID(stockProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "stockProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().StockProduct().DeleteStockProduct(
		context.Background(),
		&storehouse_service.StockProductPrimaryKey{Id: stockProductId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
