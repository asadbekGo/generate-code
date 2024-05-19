package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateChapterMachine godoc
// @Security ApiKeyAuth
// @ID create_chapterMachine
// @Router /chapter-machine [POST]
// @Summary Create ChapterMachine
// @Description Create ChapterMachine
// @Tags ChapterMachine
// @Accept json
// @Produce json
// @Param ChapterMachine body storehouse_service.CreateChapterMachineRequest true "CreateChapterMachineRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.ChapterMachine} "ChapterMachine data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateChapterMachine(c *gin.Context) {

	var chapterMachine storehouse_service.CreateChapterMachineRequest
	err := c.ShouldBindJSON(&chapterMachine)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().ChapterMachine().CreateChapterMachine(
		context.Background(),
		&chapterMachine,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleChapterMachine godoc
// @Security ApiKeyAuth
// @ID get_chapterMachine_by_id
// @Router /chapter-machine/{chapter_machine_id} [GET]
// @Summary Get single ChapterMachine
// @Description Get single ChapterMachine
// @Tags ChapterMachine
// @Accept json
// @Produce json
// @Param chapter_machine_id path string true "chapter_machine_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.ChapterMachine} "ChapterMachineBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleChapterMachine(c *gin.Context) {

	var chapterMachineId = c.Param("chapter_machine_id")
	if !util.IsValidUUID(chapterMachineId) {
		h.HandleResponse(c, status_http.InvalidArgument, "chapterMachine id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().ChapterMachine().GetByIDChapterMachine(
		context.Background(),
		&storehouse_service.ChapterMachinePrimaryKey{Id: chapterMachineId},
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

// GetChapterMachineList godoc
// @Security ApiKeyAuth
// @ID get_chapterMachine_list
// @Router /chapter-machine [GET]
// @Summary Get ChapterMachine list
// @Description Get ChapterMachine list
// @Tags ChapterMachine
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListChapterMachineRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListChapterMachineResponse} "ChapterMachineBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetChapterMachineList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().ChapterMachine().GetListChapterMachine(
		context.Background(),
		&storehouse_service.GetListChapterMachineRequest{
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

// UpdateChapterMachine godoc
// @Security ApiKeyAuth
// @ID update_chapterMachine
// @Router /chapter-machine [PUT]
// @Summary Update ChapterMachine
// @Description Update ChapterMachine
// @Tags ChapterMachine
// @Accept json
// @Produce json
// @Param ChapterMachine body storehouse_service.UpdateChapterMachineRequest true "UpdateChapterMachineRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.ChapterMachine} "ChapterMachine data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateChapterMachine(c *gin.Context) {

	var updateChapterMachine storehouse_service.UpdateChapterMachineRequest
	err := c.ShouldBindJSON(&updateChapterMachine)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().ChapterMachine().UpdateChapterMachine(
		context.Background(),
		&updateChapterMachine,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteChapterMachine godoc
// @Security ApiKeyAuth
// @ID delete_chapterMachine
// @Router /chapter-machine/{chapter_machine_id} [DELETE]
// @Summary Delete ChapterMachine
// @Description Delete ChapterMachine
// @Tags ChapterMachine
// @Accept json
// @Produce json
// @Param chapter_machine_id path string true "chapter_machine_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteChapterMachine(c *gin.Context) {

	var chapterMachineId = c.Param("chapter_machine_id")
	if !util.IsValidUUID(chapterMachineId) {
		h.HandleResponse(c, status_http.InvalidArgument, "chapterMachine id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().ChapterMachine().DeleteChapterMachine(
		context.Background(),
		&storehouse_service.ChapterMachinePrimaryKey{Id: chapterMachineId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
