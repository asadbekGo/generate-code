package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateMachine godoc
// @Security ApiKeyAuth
// @ID create_machine
// @Router /machine [POST]
// @Summary Create Machine
// @Description Create Machine
// @Tags Machine
// @Accept json
// @Produce json
// @Param Machine body storehouse_service.CreateMachineRequest true "CreateMachineRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.Machine} "Machine data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateMachine(c *gin.Context) {

	var machine storehouse_service.CreateMachineRequest
	err := c.ShouldBindJSON(&machine)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Machine().CreateMachine(
		context.Background(),
		&machine,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleMachine godoc
// @Security ApiKeyAuth
// @ID get_machine_by_id
// @Router /machine/{machine_id} [GET]
// @Summary Get single Machine
// @Description Get single Machine
// @Tags Machine
// @Accept json
// @Produce json
// @Param machine_id path string true "machine_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.Machine} "MachineBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleMachine(c *gin.Context) {

	var machineId = c.Param("machine_id")
	if !util.IsValidUUID(machineId) {
		h.HandleResponse(c, status_http.InvalidArgument, "machine id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Machine().GetByIDMachine(
		context.Background(),
		&storehouse_service.MachinePrimaryKey{Id: machineId},
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

// GetMachineList godoc
// @Security ApiKeyAuth
// @ID get_machine_list
// @Router /machine [GET]
// @Summary Get Machine list
// @Description Get Machine list
// @Tags Machine
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListMachineRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListMachineResponse} "MachineBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetMachineList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().Machine().GetListMachine(
		context.Background(),
		&storehouse_service.GetListMachineRequest{
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

// UpdateMachine godoc
// @Security ApiKeyAuth
// @ID update_machine
// @Router /machine [PUT]
// @Summary Update Machine
// @Description Update Machine
// @Tags Machine
// @Accept json
// @Produce json
// @Param Machine body storehouse_service.UpdateMachineRequest true "UpdateMachineRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.Machine} "Machine data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateMachine(c *gin.Context) {

	var updateMachine storehouse_service.UpdateMachineRequest
	err := c.ShouldBindJSON(&updateMachine)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Machine().UpdateMachine(
		context.Background(),
		&updateMachine,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteMachine godoc
// @Security ApiKeyAuth
// @ID delete_machine
// @Router /machine/{machine_id} [DELETE]
// @Summary Delete Machine
// @Description Delete Machine
// @Tags Machine
// @Accept json
// @Produce json
// @Param machine_id path string true "machine_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteMachine(c *gin.Context) {

	var machineId = c.Param("machine_id")
	if !util.IsValidUUID(machineId) {
		h.HandleResponse(c, status_http.InvalidArgument, "machine id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Machine().DeleteMachine(
		context.Background(),
		&storehouse_service.MachinePrimaryKey{Id: machineId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
