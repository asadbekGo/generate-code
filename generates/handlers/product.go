package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateProduct godoc
// @Security ApiKeyAuth
// @ID create_product
// @Router /product [POST]
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body storehouse_service.CreateProductRequest true "CreateProductRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.Product} "Product data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateProduct(c *gin.Context) {

	var product storehouse_service.CreateProductRequest
	err := c.ShouldBindJSON(&product)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Product().CreateProduct(
		context.Background(),
		&product,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleProduct godoc
// @Security ApiKeyAuth
// @ID get_product_by_id
// @Router /product/{product_id} [GET]
// @Summary Get single Product
// @Description Get single Product
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.Product} "ProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleProduct(c *gin.Context) {

	var productId = c.Param("product_id")
	if !util.IsValidUUID(productId) {
		h.HandleResponse(c, status_http.InvalidArgument, "product id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Product().GetByIDProduct(
		context.Background(),
		&storehouse_service.ProductPrimaryKey{Id: productId},
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

// GetProductList godoc
// @Security ApiKeyAuth
// @ID get_product_list
// @Router /product [GET]
// @Summary Get Product list
// @Description Get Product list
// @Tags Product
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListProductRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListProductResponse} "ProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetProductList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().Product().GetListProduct(
		context.Background(),
		&storehouse_service.GetListProductRequest{
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

// UpdateProduct godoc
// @Security ApiKeyAuth
// @ID update_product
// @Router /product [PUT]
// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body storehouse_service.UpdateProductRequest true "UpdateProductRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.Product} "Product data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateProduct(c *gin.Context) {

	var updateProduct storehouse_service.UpdateProductRequest
	err := c.ShouldBindJSON(&updateProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Product().UpdateProduct(
		context.Background(),
		&updateProduct,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteProduct godoc
// @Security ApiKeyAuth
// @ID delete_product
// @Router /product/{product_id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteProduct(c *gin.Context) {

	var productId = c.Param("product_id")
	if !util.IsValidUUID(productId) {
		h.HandleResponse(c, status_http.InvalidArgument, "product id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Product().DeleteProduct(
		context.Background(),
		&storehouse_service.ProductPrimaryKey{Id: productId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
