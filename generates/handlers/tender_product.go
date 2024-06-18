package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateTenderProduct godoc
// @Security ApiKeyAuth
// @ID create_tenderProduct
// @Router /v1/tender-product [POST]
// @Summary Create TenderProduct
// @Description Create TenderProduct
// @Tags TenderProduct
// @Accept json
// @Produce json
// @Param TenderProduct body storehouse_service.CreateTenderProductRequest true "CreateTenderProductRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.TenderProduct} "TenderProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateTenderProduct(c *gin.Context) {

	var tenderProduct storehouse_service.CreateTenderProductRequest
	err := c.ShouldBindJSON(&tenderProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TenderProduct().CreateTenderProduct(
		context.Background(),
		&tenderProduct,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleTenderProduct godoc
// @Security ApiKeyAuth
// @ID get_tenderProduct_by_id
// @Router /v1/tender-product/{tender_product_id} [GET]
// @Summary Get single TenderProduct
// @Description Get single TenderProduct
// @Tags TenderProduct
// @Accept json
// @Produce json
// @Param tender_product_id path string true "tender_product_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.TenderProduct} "TenderProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleTenderProduct(c *gin.Context) {

	var tenderProductId = c.Param("tender_product_id")
	if !util.IsValidUUID(tenderProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "tenderProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TenderProduct().GetByIDTenderProduct(
		context.Background(),
		&storehouse_service.TenderProductPrimaryKey{Id: tenderProductId},
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

// GetTenderProductList godoc
// @Security ApiKeyAuth
// @ID get_tenderProduct_list
// @Router /v1/tender-product [GET]
// @Summary Get TenderProduct list
// @Description Get TenderProduct list
// @Tags TenderProduct
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListTenderProductRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListTenderProductResponse} "TenderProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetTenderProductList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().TenderProduct().GetListTenderProduct(
		context.Background(),
		&storehouse_service.GetListTenderProductRequest{
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

// UpdateTenderProduct godoc
// @Security ApiKeyAuth
// @ID update_tenderProduct
// @Router /v1/tender-product [PUT]
// @Summary Update TenderProduct
// @Description Update TenderProduct
// @Tags TenderProduct
// @Accept json
// @Produce json
// @Param TenderProduct body storehouse_service.UpdateTenderProductRequest true "UpdateTenderProductRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.TenderProduct} "TenderProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateTenderProduct(c *gin.Context) {

	var updateTenderProduct storehouse_service.UpdateTenderProductRequest
	err := c.ShouldBindJSON(&updateTenderProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TenderProduct().UpdateTenderProduct(
		context.Background(),
		&updateTenderProduct,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteTenderProduct godoc
// @Security ApiKeyAuth
// @ID delete_tenderProduct
// @Router /v1/tender-product/{tender_product_id} [DELETE]
// @Summary Delete TenderProduct
// @Description Delete TenderProduct
// @Tags TenderProduct
// @Accept json
// @Produce json
// @Param tender_product_id path string true "tender_product_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteTenderProduct(c *gin.Context) {

	var tenderProductId = c.Param("tender_product_id")
	if !util.IsValidUUID(tenderProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "tenderProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TenderProduct().DeleteTenderProduct(
		context.Background(),
		&storehouse_service.TenderProductPrimaryKey{Id: tenderProductId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
