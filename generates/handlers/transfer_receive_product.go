package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateTransferReceiveProduct godoc
// @Security ApiKeyAuth
// @ID create_transferReceiveProduct
// @Router /transfer-receive-product [POST]
// @Summary Create TransferReceiveProduct
// @Description Create TransferReceiveProduct
// @Tags TransferReceiveProduct
// @Accept json
// @Produce json
// @Param TransferReceiveProduct body storehouse_service.CreateTransferReceiveProductRequest true "CreateTransferReceiveProductRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.TransferReceiveProduct} "TransferReceiveProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateTransferReceiveProduct(c *gin.Context) {

	var transferReceiveProduct storehouse_service.CreateTransferReceiveProductRequest
	err := c.ShouldBindJSON(&transferReceiveProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TransferReceiveProduct().CreateTransferReceiveProduct(
		context.Background(),
		&transferReceiveProduct,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleTransferReceiveProduct godoc
// @Security ApiKeyAuth
// @ID get_transferReceiveProduct_by_id
// @Router /transfer-receive-product/{transfer_receive_product_id} [GET]
// @Summary Get single TransferReceiveProduct
// @Description Get single TransferReceiveProduct
// @Tags TransferReceiveProduct
// @Accept json
// @Produce json
// @Param transfer_receive_product_id path string true "transfer_receive_product_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.TransferReceiveProduct} "TransferReceiveProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleTransferReceiveProduct(c *gin.Context) {

	var transferReceiveProductId = c.Param("transfer_receive_product_id")
	if !util.IsValidUUID(transferReceiveProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "transferReceiveProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TransferReceiveProduct().GetByIDTransferReceiveProduct(
		context.Background(),
		&storehouse_service.TransferReceiveProductPrimaryKey{Id: transferReceiveProductId},
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

// GetTransferReceiveProductList godoc
// @Security ApiKeyAuth
// @ID get_transferReceiveProduct_list
// @Router /transfer-receive-product [GET]
// @Summary Get TransferReceiveProduct list
// @Description Get TransferReceiveProduct list
// @Tags TransferReceiveProduct
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListTransferReceiveProductRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListTransferReceiveProductResponse} "TransferReceiveProductBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetTransferReceiveProductList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().TransferReceiveProduct().GetListTransferReceiveProduct(
		context.Background(),
		&storehouse_service.GetListTransferReceiveProductRequest{
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

// UpdateTransferReceiveProduct godoc
// @Security ApiKeyAuth
// @ID update_transferReceiveProduct
// @Router /transfer-receive-product [PUT]
// @Summary Update TransferReceiveProduct
// @Description Update TransferReceiveProduct
// @Tags TransferReceiveProduct
// @Accept json
// @Produce json
// @Param TransferReceiveProduct body storehouse_service.UpdateTransferReceiveProductRequest true "UpdateTransferReceiveProductRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.TransferReceiveProduct} "TransferReceiveProduct data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateTransferReceiveProduct(c *gin.Context) {

	var updateTransferReceiveProduct storehouse_service.UpdateTransferReceiveProductRequest
	err := c.ShouldBindJSON(&updateTransferReceiveProduct)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TransferReceiveProduct().UpdateTransferReceiveProduct(
		context.Background(),
		&updateTransferReceiveProduct,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteTransferReceiveProduct godoc
// @Security ApiKeyAuth
// @ID delete_transferReceiveProduct
// @Router /transfer-receive-product/{transfer_receive_product_id} [DELETE]
// @Summary Delete TransferReceiveProduct
// @Description Delete TransferReceiveProduct
// @Tags TransferReceiveProduct
// @Accept json
// @Produce json
// @Param transfer_receive_product_id path string true "transfer_receive_product_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteTransferReceiveProduct(c *gin.Context) {

	var transferReceiveProductId = c.Param("transfer_receive_product_id")
	if !util.IsValidUUID(transferReceiveProductId) {
		h.HandleResponse(c, status_http.InvalidArgument, "transferReceiveProduct id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TransferReceiveProduct().DeleteTransferReceiveProduct(
		context.Background(),
		&storehouse_service.TransferReceiveProductPrimaryKey{Id: transferReceiveProductId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
