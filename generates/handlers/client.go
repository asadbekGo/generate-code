package client_handler

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_client_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateClient godoc
// @Security ApiKeyAuth
// @ID create_client
// @Router /v1/client [POST]
// @Summary Create Client
// @Description Create Client
// @Tags Client
// @Accept json
// @Produce json
// @Param Client body storehouse_client_service.CreateClientRequest true "CreateClientRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_client_service.Client} "Client data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateClient(c *gin.Context) {

	var client storehouse_client_service.CreateClientRequest
	err := c.ShouldBindJSON(&client)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().Client().CreateClient(
		context.Background(),
		&client,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleClient godoc
// @Security ApiKeyAuth
// @ID get_client_by_id
// @Router /v1/client/{client_id} [GET]
// @Summary Get single Client
// @Description Get single Client
// @Tags Client
// @Accept json
// @Produce json
// @Param client_id path string true "client_id"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.Client} "ClientBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleClient(c *gin.Context) {

	var clientId = c.Param("client_id")
	if !util.IsValidUUID(clientId) {
		h.HandleResponse(c, status_http.InvalidArgument, "client id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().Client().GetByIDClient(
		context.Background(),
		&storehouse_client_service.ClientPrimaryKey{Id: clientId},
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

// GetClientList godoc
// @Security ApiKeyAuth
// @ID get_client_list
// @Router /v1/client [GET]
// @Summary Get Client list
// @Description Get Client list
// @Tags Client
// @Accept json
// @Produce json
// @Param filters query storehouse_client_service.GetListClientRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.GetListClientResponse} "ClientBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetClientList(c *gin.Context) {

	page, err := h.GetPageParam(c)
	if err != nil {
		h.HandleResponse(c, status_http.InvalidArgument, err.Error())
		return
	}


	limit, err := h.GetLimitParam(c)
	if err != nil {
		h.HandleResponse(c, status_http.InvalidArgument, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().Client().GetListClient(
		context.Background(),
		&storehouse_client_service.GetListClientRequest{
			Limit:  int32(limit),
			Page:   int32(page),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.OK, response)
}

// UpdateClient godoc
// @Security ApiKeyAuth
// @ID update_client
// @Router /v1/client [PUT]
// @Summary Update Client
// @Description Update Client
// @Tags Client
// @Accept json
// @Produce json
// @Param Client body storehouse_client_service.UpdateClientRequest true "UpdateClientRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.Client} "Client data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateClient(c *gin.Context) {

	var updateClient storehouse_client_service.UpdateClientRequest
	err := c.ShouldBindJSON(&updateClient)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().Client().UpdateClient(
		context.Background(),
		&updateClient,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteClient godoc
// @Security ApiKeyAuth
// @ID delete_client
// @Router /v1/client/{client_id} [DELETE]
// @Summary Delete Client
// @Description Delete Client
// @Tags Client
// @Accept json
// @Produce json
// @Param client_id path string true "client_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteClient(c *gin.Context) {

	var clientId = c.Param("client_id")
	if !util.IsValidUUID(clientId) {
		h.HandleResponse(c, status_http.InvalidArgument, "client id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().Client().DeleteClient(
		context.Background(),
		&storehouse_client_service.ClientPrimaryKey{Id: clientId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
