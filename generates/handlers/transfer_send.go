package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateTransferSend godoc
// @Security ApiKeyAuth
// @ID create_transferSend
// @Router /transfer-send [POST]
// @Summary Create TransferSend
// @Description Create TransferSend
// @Tags TransferSend
// @Accept json
// @Produce json
// @Param TransferSend body storehouse_service.CreateTransferSendRequest true "CreateTransferSendRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.TransferSend} "TransferSend data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateTransferSend(c *gin.Context) {

	var transferSend storehouse_service.CreateTransferSendRequest
	err := c.ShouldBindJSON(&transferSend)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TransferSend().CreateTransferSend(
		context.Background(),
		&transferSend,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleTransferSend godoc
// @Security ApiKeyAuth
// @ID get_transferSend_by_id
// @Router /transfer-send/{transfer_send_id} [GET]
// @Summary Get single TransferSend
// @Description Get single TransferSend
// @Tags TransferSend
// @Accept json
// @Produce json
// @Param transfer_send_id path string true "transfer_send_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.TransferSend} "TransferSendBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleTransferSend(c *gin.Context) {

	var transferSendId = c.Param("transfer_send_id")
	if !util.IsValidUUID(transferSendId) {
		h.HandleResponse(c, status_http.InvalidArgument, "transferSend id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TransferSend().GetByIDTransferSend(
		context.Background(),
		&storehouse_service.TransferSendPrimaryKey{Id: transferSendId},
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

// GetTransferSendList godoc
// @Security ApiKeyAuth
// @ID get_transferSend_list
// @Router /transfer-send [GET]
// @Summary Get TransferSend list
// @Description Get TransferSend list
// @Tags TransferSend
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListTransferSendRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListTransferSendResponse} "TransferSendBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetTransferSendList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().TransferSend().GetListTransferSend(
		context.Background(),
		&storehouse_service.GetListTransferSendRequest{
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

// UpdateTransferSend godoc
// @Security ApiKeyAuth
// @ID update_transferSend
// @Router /transfer-send [PUT]
// @Summary Update TransferSend
// @Description Update TransferSend
// @Tags TransferSend
// @Accept json
// @Produce json
// @Param TransferSend body storehouse_service.UpdateTransferSendRequest true "UpdateTransferSendRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.TransferSend} "TransferSend data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateTransferSend(c *gin.Context) {

	var updateTransferSend storehouse_service.UpdateTransferSendRequest
	err := c.ShouldBindJSON(&updateTransferSend)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TransferSend().UpdateTransferSend(
		context.Background(),
		&updateTransferSend,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteTransferSend godoc
// @Security ApiKeyAuth
// @ID delete_transferSend
// @Router /transfer-send/{transfer_send_id} [DELETE]
// @Summary Delete TransferSend
// @Description Delete TransferSend
// @Tags TransferSend
// @Accept json
// @Produce json
// @Param transfer_send_id path string true "transfer_send_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteTransferSend(c *gin.Context) {

	var transferSendId = c.Param("transfer_send_id")
	if !util.IsValidUUID(transferSendId) {
		h.HandleResponse(c, status_http.InvalidArgument, "transferSend id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TransferSend().DeleteTransferSend(
		context.Background(),
		&storehouse_service.TransferSendPrimaryKey{Id: transferSendId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
