package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateTransferReceive godoc
// @Security ApiKeyAuth
// @ID create_transferReceive
// @Router /transfer-receive [POST]
// @Summary Create TransferReceive
// @Description Create TransferReceive
// @Tags TransferReceive
// @Accept json
// @Produce json
// @Param TransferReceive body storehouse_service.CreateTransferReceiveRequest true "CreateTransferReceiveRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.TransferReceive} "TransferReceive data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateTransferReceive(c *gin.Context) {

	var transferReceive storehouse_service.CreateTransferReceiveRequest
	err := c.ShouldBindJSON(&transferReceive)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TransferReceive().CreateTransferReceive(
		context.Background(),
		&transferReceive,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleTransferReceive godoc
// @Security ApiKeyAuth
// @ID get_transferReceive_by_id
// @Router /transfer-receive/{transfer_receive_id} [GET]
// @Summary Get single TransferReceive
// @Description Get single TransferReceive
// @Tags TransferReceive
// @Accept json
// @Produce json
// @Param transfer_receive_id path string true "transfer_receive_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.TransferReceive} "TransferReceiveBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleTransferReceive(c *gin.Context) {

	var transferReceiveId = c.Param("transfer_receive_id")
	if !util.IsValidUUID(transferReceiveId) {
		h.HandleResponse(c, status_http.InvalidArgument, "transferReceive id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TransferReceive().GetByIDTransferReceive(
		context.Background(),
		&storehouse_service.TransferReceivePrimaryKey{Id: transferReceiveId},
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

// GetTransferReceiveList godoc
// @Security ApiKeyAuth
// @ID get_transferReceive_list
// @Router /transfer-receive [GET]
// @Summary Get TransferReceive list
// @Description Get TransferReceive list
// @Tags TransferReceive
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListTransferReceiveRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListTransferReceiveResponse} "TransferReceiveBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetTransferReceiveList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().TransferReceive().GetListTransferReceive(
		context.Background(),
		&storehouse_service.GetListTransferReceiveRequest{
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

// UpdateTransferReceive godoc
// @Security ApiKeyAuth
// @ID update_transferReceive
// @Router /transfer-receive [PUT]
// @Summary Update TransferReceive
// @Description Update TransferReceive
// @Tags TransferReceive
// @Accept json
// @Produce json
// @Param TransferReceive body storehouse_service.UpdateTransferReceiveRequest true "UpdateTransferReceiveRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.TransferReceive} "TransferReceive data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateTransferReceive(c *gin.Context) {

	var updateTransferReceive storehouse_service.UpdateTransferReceiveRequest
	err := c.ShouldBindJSON(&updateTransferReceive)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TransferReceive().UpdateTransferReceive(
		context.Background(),
		&updateTransferReceive,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteTransferReceive godoc
// @Security ApiKeyAuth
// @ID delete_transferReceive
// @Router /transfer-receive/{transfer_receive_id} [DELETE]
// @Summary Delete TransferReceive
// @Description Delete TransferReceive
// @Tags TransferReceive
// @Accept json
// @Produce json
// @Param transfer_receive_id path string true "transfer_receive_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteTransferReceive(c *gin.Context) {

	var transferReceiveId = c.Param("transfer_receive_id")
	if !util.IsValidUUID(transferReceiveId) {
		h.HandleResponse(c, status_http.InvalidArgument, "transferReceive id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TransferReceive().DeleteTransferReceive(
		context.Background(),
		&storehouse_service.TransferReceivePrimaryKey{Id: transferReceiveId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
