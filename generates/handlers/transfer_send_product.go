package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateTransferSendProduct godoc
// @Security ApiKeyAuth
// @ID create_transferSendProduct
// @Router /transfer-send-product [POST]
// @Summary Create TransferSendProduct
// @Description Create TransferSendProduct
// @Tags TransferSendProduct
// @Accept json
// @Produce json
// @Param TransferSendProduct body storehouse_service.CreateTransferSendProductRequest true "CreateTransferSendProductRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.TransferSendProduct} "TransferSendProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateTransferSendProduct(c *gin.Context) {

	var transferSendProduct storehouse_service.CreateTransferSendProductRequest
	err := c.ShouldBindJSON(&transferSendProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TransferSendProduct().CreateTransferSendProduct(
		context.Background(),
		&transferSendProduct,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleTransferSendProduct godoc
// @Security ApiKeyAuth
// @ID get_transferSendProduct_by_id
// @Router /transfer-send-product/{transfer_send_product_id} [GET]
// @Summary Get single TransferSendProduct
// @Description Get single TransferSendProduct
// @Tags TransferSendProduct
// @Accept json
// @Produce json
// @Param transfer_send_product_id path string true "transfer_send_product_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.TransferSendProduct} "TransferSendProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleTransferSendProduct(c *gin.Context) {

	var transferSendProductId = c.Param("transfer_send_product_id")
	if !util.IsValidUUID(transferSendProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "transferSendProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TransferSendProduct().GetByIDTransferSendProduct(
		context.Background(),
		&storehouse_service.TransferSendProductPrimaryKey{Id: transferSendProductId},
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

// GetTransferSendProductList godoc
// @Security ApiKeyAuth
// @ID get_transferSendProduct_list
// @Router /transfer-send-product [GET]
// @Summary Get TransferSendProduct list
// @Description Get TransferSendProduct list
// @Tags TransferSendProduct
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListTransferSendProductRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListTransferSendProductResponse} "TransferSendProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetTransferSendProductList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().TransferSendProduct().GetListTransferSendProduct(
		context.Background(),
		&storehouse_service.GetListTransferSendProductRequest{
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

// UpdateTransferSendProduct godoc
// @Security ApiKeyAuth
// @ID update_transferSendProduct
// @Router /transfer-send-product [PUT]
// @Summary Update TransferSendProduct
// @Description Update TransferSendProduct
// @Tags TransferSendProduct
// @Accept json
// @Produce json
// @Param TransferSendProduct body storehouse_service.UpdateTransferSendProductRequest true "UpdateTransferSendProductRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.TransferSendProduct} "TransferSendProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateTransferSendProduct(c *gin.Context) {

	var updateTransferSendProduct storehouse_service.UpdateTransferSendProductRequest
	err := c.ShouldBindJSON(&updateTransferSendProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TransferSendProduct().UpdateTransferSendProduct(
		context.Background(),
		&updateTransferSendProduct,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteTransferSendProduct godoc
// @Security ApiKeyAuth
// @ID delete_transferSendProduct
// @Router /transfer-send-product/{transfer_send_product_id} [DELETE]
// @Summary Delete TransferSendProduct
// @Description Delete TransferSendProduct
// @Tags TransferSendProduct
// @Accept json
// @Produce json
// @Param transfer_send_product_id path string true "transfer_send_product_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteTransferSendProduct(c *gin.Context) {

	var transferSendProductId = c.Param("transfer_send_product_id")
	if !util.IsValidUUID(transferSendProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "transferSendProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TransferSendProduct().DeleteTransferSendProduct(
		context.Background(),
		&storehouse_service.TransferSendProductPrimaryKey{Id: transferSendProductId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
