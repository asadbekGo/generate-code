package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateUser godoc
// @Security ApiKeyAuth
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param User body storehouse_service.CreateUserRequest true "CreateUserRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.User} "User data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {

	var user storehouse_service.CreateUserRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().User().CreateUser(
		context.Background(),
		&user,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleUser godoc
// @Security ApiKeyAuth
// @ID get_user_by_id
// @Router /user/{user_id} [GET]
// @Summary Get single User
// @Description Get single User
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.User} "UserBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleUser(c *gin.Context) {

	var userId = c.Param("user_id")
	if !util.IsValidUUID(userId) {
		h.HandleResponse(c, status_http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().User().GetByIDUser(
		context.Background(),
		&storehouse_service.UserPrimaryKey{Id: userId},
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

// GetUserList godoc
// @Security ApiKeyAuth
// @ID get_user_list
// @Router /user [GET]
// @Summary Get User list
// @Description Get User list
// @Tags User
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListUserRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListUserResponse} "UserBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetUserList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().User().GetListUser(
		context.Background(),
		&storehouse_service.GetListUserRequest{
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

// UpdateUser godoc
// @Security ApiKeyAuth
// @ID update_user
// @Router /user [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param User body storehouse_service.UpdateUserRequest true "UpdateUserRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.User} "User data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {

	var updateUser storehouse_service.UpdateUserRequest
	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().User().UpdateUser(
		context.Background(),
		&updateUser,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteUser godoc
// @Security ApiKeyAuth
// @ID delete_user
// @Router /user/{user_id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {

	var userId = c.Param("user_id")
	if !util.IsValidUUID(userId) {
		h.HandleResponse(c, status_http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().User().DeleteUser(
		context.Background(),
		&storehouse_service.UserPrimaryKey{Id: userId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
