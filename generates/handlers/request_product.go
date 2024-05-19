package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateRequestProduct godoc
// @Security ApiKeyAuth
// @ID create_requestProduct
// @Router /request-product [POST]
// @Summary Create RequestProduct
// @Description Create RequestProduct
// @Tags RequestProduct
// @Accept json
// @Produce json
// @Param RequestProduct body storehouse_service.CreateRequestProductRequest true "CreateRequestProductRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.RequestProduct} "RequestProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateRequestProduct(c *gin.Context) {

	var requestProduct storehouse_service.CreateRequestProductRequest
	err := c.ShouldBindJSON(&requestProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().RequestProduct().CreateRequestProduct(
		context.Background(),
		&requestProduct,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleRequestProduct godoc
// @Security ApiKeyAuth
// @ID get_requestProduct_by_id
// @Router /request-product/{request_product_id} [GET]
// @Summary Get single RequestProduct
// @Description Get single RequestProduct
// @Tags RequestProduct
// @Accept json
// @Produce json
// @Param request_product_id path string true "request_product_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.RequestProduct} "RequestProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleRequestProduct(c *gin.Context) {

	var requestProductId = c.Param("request_product_id")
	if !util.IsValidUUID(requestProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "requestProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().RequestProduct().GetByIDRequestProduct(
		context.Background(),
		&storehouse_service.RequestProductPrimaryKey{Id: requestProductId},
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

// GetRequestProductList godoc
// @Security ApiKeyAuth
// @ID get_requestProduct_list
// @Router /request-product [GET]
// @Summary Get RequestProduct list
// @Description Get RequestProduct list
// @Tags RequestProduct
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListRequestProductRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListRequestProductResponse} "RequestProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetRequestProductList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().RequestProduct().GetListRequestProduct(
		context.Background(),
		&storehouse_service.GetListRequestProductRequest{
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

// UpdateRequestProduct godoc
// @Security ApiKeyAuth
// @ID update_requestProduct
// @Router /request-product [PUT]
// @Summary Update RequestProduct
// @Description Update RequestProduct
// @Tags RequestProduct
// @Accept json
// @Produce json
// @Param RequestProduct body storehouse_service.UpdateRequestProductRequest true "UpdateRequestProductRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.RequestProduct} "RequestProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateRequestProduct(c *gin.Context) {

	var updateRequestProduct storehouse_service.UpdateRequestProductRequest
	err := c.ShouldBindJSON(&updateRequestProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().RequestProduct().UpdateRequestProduct(
		context.Background(),
		&updateRequestProduct,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteRequestProduct godoc
// @Security ApiKeyAuth
// @ID delete_requestProduct
// @Router /request-product/{request_product_id} [DELETE]
// @Summary Delete RequestProduct
// @Description Delete RequestProduct
// @Tags RequestProduct
// @Accept json
// @Produce json
// @Param request_product_id path string true "request_product_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteRequestProduct(c *gin.Context) {

	var requestProductId = c.Param("request_product_id")
	if !util.IsValidUUID(requestProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "requestProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().RequestProduct().DeleteRequestProduct(
		context.Background(),
		&storehouse_service.RequestProductPrimaryKey{Id: requestProductId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
