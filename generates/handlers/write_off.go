package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateWriteOff godoc
// @Security ApiKeyAuth
// @ID create_writeOff
// @Router /write-off [POST]
// @Summary Create WriteOff
// @Description Create WriteOff
// @Tags WriteOff
// @Accept json
// @Produce json
// @Param WriteOff body storehouse_service.CreateWriteOffRequest true "CreateWriteOffRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.WriteOff} "WriteOff data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateWriteOff(c *gin.Context) {

	var writeOff storehouse_service.CreateWriteOffRequest
	err := c.ShouldBindJSON(&writeOff)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().WriteOff().CreateWriteOff(
		context.Background(),
		&writeOff,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleWriteOff godoc
// @Security ApiKeyAuth
// @ID get_writeOff_by_id
// @Router /write-off/{write_off_id} [GET]
// @Summary Get single WriteOff
// @Description Get single WriteOff
// @Tags WriteOff
// @Accept json
// @Produce json
// @Param write_off_id path string true "write_off_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.WriteOff} "WriteOffBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleWriteOff(c *gin.Context) {

	var writeOffId = c.Param("write_off_id")
	if !util.IsValidUUID(writeOffId) {
		h.HandleResponse(c, status_http.InvalidArgument, "writeOff id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().WriteOff().GetByIDWriteOff(
		context.Background(),
		&storehouse_service.WriteOffPrimaryKey{Id: writeOffId},
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

// GetWriteOffList godoc
// @Security ApiKeyAuth
// @ID get_writeOff_list
// @Router /write-off [GET]
// @Summary Get WriteOff list
// @Description Get WriteOff list
// @Tags WriteOff
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListWriteOffRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListWriteOffResponse} "WriteOffBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetWriteOffList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().WriteOff().GetListWriteOff(
		context.Background(),
		&storehouse_service.GetListWriteOffRequest{
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

// UpdateWriteOff godoc
// @Security ApiKeyAuth
// @ID update_writeOff
// @Router /write-off [PUT]
// @Summary Update WriteOff
// @Description Update WriteOff
// @Tags WriteOff
// @Accept json
// @Produce json
// @Param WriteOff body storehouse_service.UpdateWriteOffRequest true "UpdateWriteOffRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.WriteOff} "WriteOff data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateWriteOff(c *gin.Context) {

	var updateWriteOff storehouse_service.UpdateWriteOffRequest
	err := c.ShouldBindJSON(&updateWriteOff)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().WriteOff().UpdateWriteOff(
		context.Background(),
		&updateWriteOff,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteWriteOff godoc
// @Security ApiKeyAuth
// @ID delete_writeOff
// @Router /write-off/{write_off_id} [DELETE]
// @Summary Delete WriteOff
// @Description Delete WriteOff
// @Tags WriteOff
// @Accept json
// @Produce json
// @Param write_off_id path string true "write_off_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteWriteOff(c *gin.Context) {

	var writeOffId = c.Param("write_off_id")
	if !util.IsValidUUID(writeOffId) {
		h.HandleResponse(c, status_http.InvalidArgument, "writeOff id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().WriteOff().DeleteWriteOff(
		context.Background(),
		&storehouse_service.WriteOffPrimaryKey{Id: writeOffId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
