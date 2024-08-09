package client_handler

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_client_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateClientContract godoc
// @Security ApiKeyAuth
// @ID create_clientContract
// @Router /v1/client-contract [POST]
// @Summary Create ClientContract
// @Description Create ClientContract
// @Tags ClientContract
// @Accept json
// @Produce json
// @Param ClientContract body storehouse_client_service.CreateClientContractRequest true "CreateClientContractRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_client_service.ClientContract} "ClientContract data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateClientContract(c *gin.Context) {

	var clientContract storehouse_client_service.CreateClientContractRequest
	err := c.ShouldBindJSON(&clientContract)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().ClientContract().CreateClientContract(
		context.Background(),
		&clientContract,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleClientContract godoc
// @Security ApiKeyAuth
// @ID get_clientContract_by_id
// @Router /v1/client-contract/{client_contract_id} [GET]
// @Summary Get single ClientContract
// @Description Get single ClientContract
// @Tags ClientContract
// @Accept json
// @Produce json
// @Param client_contract_id path string true "client_contract_id"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.ClientContract} "ClientContractBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleClientContract(c *gin.Context) {

	var clientContractId = c.Param("client_contract_id")
	if !util.IsValidUUID(clientContractId) {
		h.HandleResponse(c, status_http.InvalidArgument, "clientContract id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().ClientContract().GetByIDClientContract(
		context.Background(),
		&storehouse_client_service.ClientContractPrimaryKey{Id: clientContractId},
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

// GetClientContractList godoc
// @Security ApiKeyAuth
// @ID get_clientContract_list
// @Router /v1/client-contract [GET]
// @Summary Get ClientContract list
// @Description Get ClientContract list
// @Tags ClientContract
// @Accept json
// @Produce json
// @Param filters query storehouse_client_service.GetListClientContractRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.GetListClientContractResponse} "ClientContractBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetClientContractList(c *gin.Context) {

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

	response, err := h.services.StorehouseClientService().ClientContract().GetListClientContract(
		context.Background(),
		&storehouse_client_service.GetListClientContractRequest{
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

// UpdateClientContract godoc
// @Security ApiKeyAuth
// @ID update_clientContract
// @Router /v1/client-contract [PUT]
// @Summary Update ClientContract
// @Description Update ClientContract
// @Tags ClientContract
// @Accept json
// @Produce json
// @Param ClientContract body storehouse_client_service.UpdateClientContractRequest true "UpdateClientContractRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_client_service.ClientContract} "ClientContract data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateClientContract(c *gin.Context) {

	var updateClientContract storehouse_client_service.UpdateClientContractRequest
	err := c.ShouldBindJSON(&updateClientContract)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseClientService().ClientContract().UpdateClientContract(
		context.Background(),
		&updateClientContract,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteClientContract godoc
// @Security ApiKeyAuth
// @ID delete_clientContract
// @Router /v1/client-contract/{client_contract_id} [DELETE]
// @Summary Delete ClientContract
// @Description Delete ClientContract
// @Tags ClientContract
// @Accept json
// @Produce json
// @Param client_contract_id path string true "client_contract_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteClientContract(c *gin.Context) {

	var clientContractId = c.Param("client_contract_id")
	if !util.IsValidUUID(clientContractId) {
		h.HandleResponse(c, status_http.InvalidArgument, "clientContract id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseClientService().ClientContract().DeleteClientContract(
		context.Background(),
		&storehouse_client_service.ClientContractPrimaryKey{Id: clientContractId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
