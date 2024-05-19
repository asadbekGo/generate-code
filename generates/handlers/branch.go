package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateBranch godoc
// @Security ApiKeyAuth
// @ID create_branch
// @Router /branch [POST]
// @Summary Create Branch
// @Description Create Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param Branch body storehouse_service.CreateBranchRequest true "CreateBranchRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.Branch} "Branch data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateBranch(c *gin.Context) {

	var branch storehouse_service.CreateBranchRequest
	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Branch().CreateBranch(
		context.Background(),
		&branch,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleBranch godoc
// @Security ApiKeyAuth
// @ID get_branch_by_id
// @Router /branch/{branch_id} [GET]
// @Summary Get single Branch
// @Description Get single Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.Branch} "BranchBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleBranch(c *gin.Context) {

	var branchId = c.Param("branch_id")
	if !util.IsValidUUID(branchId) {
		h.HandleResponse(c, status_http.InvalidArgument, "branch id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Branch().GetByIDBranch(
		context.Background(),
		&storehouse_service.BranchPrimaryKey{Id: branchId},
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

// GetBranchList godoc
// @Security ApiKeyAuth
// @ID get_branch_list
// @Router /branch [GET]
// @Summary Get Branch list
// @Description Get Branch list
// @Tags Branch
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListBranchRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListBranchResponse} "BranchBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetBranchList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().Branch().GetListBranch(
		context.Background(),
		&storehouse_service.GetListBranchRequest{
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

// UpdateBranch godoc
// @Security ApiKeyAuth
// @ID update_branch
// @Router /branch [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param Branch body storehouse_service.UpdateBranchRequest true "UpdateBranchRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.Branch} "Branch data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateBranch(c *gin.Context) {

	var updateBranch storehouse_service.UpdateBranchRequest
	err := c.ShouldBindJSON(&updateBranch)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Branch().UpdateBranch(
		context.Background(),
		&updateBranch,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteBranch godoc
// @Security ApiKeyAuth
// @ID delete_branch
// @Router /branch/{branch_id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteBranch(c *gin.Context) {

	var branchId = c.Param("branch_id")
	if !util.IsValidUUID(branchId) {
		h.HandleResponse(c, status_http.InvalidArgument, "branch id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Branch().DeleteBranch(
		context.Background(),
		&storehouse_service.BranchPrimaryKey{Id: branchId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
