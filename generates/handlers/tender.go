package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateTender godoc
// @Security ApiKeyAuth
// @ID create_tender
// @Router /v1/tender [POST]
// @Summary Create Tender
// @Description Create Tender
// @Tags Tender
// @Accept json
// @Produce json
// @Param Tender body storehouse_service.CreateTenderRequest true "CreateTenderRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.Tender} "Tender data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateTender(c *gin.Context) {

	var tender storehouse_service.CreateTenderRequest
	err := c.ShouldBindJSON(&tender)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Tender().CreateTender(
		context.Background(),
		&tender,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleTender godoc
// @Security ApiKeyAuth
// @ID get_tender_by_id
// @Router /v1/tender/{tender_id} [GET]
// @Summary Get single Tender
// @Description Get single Tender
// @Tags Tender
// @Accept json
// @Produce json
// @Param tender_id path string true "tender_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.Tender} "TenderBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleTender(c *gin.Context) {

	var tenderId = c.Param("tender_id")
	if !util.IsValidUUID(tenderId) {
		h.HandleResponse(c, status_http.InvalidArgument, "tender id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Tender().GetByIDTender(
		context.Background(),
		&storehouse_service.TenderPrimaryKey{Id: tenderId},
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

// GetTenderList godoc
// @Security ApiKeyAuth
// @ID get_tender_list
// @Router /v1/tender [GET]
// @Summary Get Tender list
// @Description Get Tender list
// @Tags Tender
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListTenderRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListTenderResponse} "TenderBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetTenderList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().Tender().GetListTender(
		context.Background(),
		&storehouse_service.GetListTenderRequest{
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

// UpdateTender godoc
// @Security ApiKeyAuth
// @ID update_tender
// @Router /v1/tender [PUT]
// @Summary Update Tender
// @Description Update Tender
// @Tags Tender
// @Accept json
// @Produce json
// @Param Tender body storehouse_service.UpdateTenderRequest true "UpdateTenderRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.Tender} "Tender data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateTender(c *gin.Context) {

	var updateTender storehouse_service.UpdateTenderRequest
	err := c.ShouldBindJSON(&updateTender)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Tender().UpdateTender(
		context.Background(),
		&updateTender,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteTender godoc
// @Security ApiKeyAuth
// @ID delete_tender
// @Router /v1/tender/{tender_id} [DELETE]
// @Summary Delete Tender
// @Description Delete Tender
// @Tags Tender
// @Accept json
// @Produce json
// @Param tender_id path string true "tender_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteTender(c *gin.Context) {

	var tenderId = c.Param("tender_id")
	if !util.IsValidUUID(tenderId) {
		h.HandleResponse(c, status_http.InvalidArgument, "tender id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Tender().DeleteTender(
		context.Background(),
		&storehouse_service.TenderPrimaryKey{Id: tenderId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
