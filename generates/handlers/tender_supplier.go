package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateTenderSupplier godoc
// @Security ApiKeyAuth
// @ID create_tenderSupplier
// @Router /v1/tender-supplier [POST]
// @Summary Create TenderSupplier
// @Description Create TenderSupplier
// @Tags TenderSupplier
// @Accept json
// @Produce json
// @Param TenderSupplier body storehouse_service.CreateTenderSupplierRequest true "CreateTenderSupplierRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.TenderSupplier} "TenderSupplier data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateTenderSupplier(c *gin.Context) {

	var tenderSupplier storehouse_service.CreateTenderSupplierRequest
	err := c.ShouldBindJSON(&tenderSupplier)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TenderSupplier().CreateTenderSupplier(
		context.Background(),
		&tenderSupplier,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleTenderSupplier godoc
// @Security ApiKeyAuth
// @ID get_tenderSupplier_by_id
// @Router /v1/tender-supplier/{tender_supplier_id} [GET]
// @Summary Get single TenderSupplier
// @Description Get single TenderSupplier
// @Tags TenderSupplier
// @Accept json
// @Produce json
// @Param tender_supplier_id path string true "tender_supplier_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.TenderSupplier} "TenderSupplierBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleTenderSupplier(c *gin.Context) {

	var tenderSupplierId = c.Param("tender_supplier_id")
	if !util.IsValidUUID(tenderSupplierId) {
		h.HandleResponse(c, status_http.InvalidArgument, "tenderSupplier id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TenderSupplier().GetByIDTenderSupplier(
		context.Background(),
		&storehouse_service.TenderSupplierPrimaryKey{Id: tenderSupplierId},
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

// GetTenderSupplierList godoc
// @Security ApiKeyAuth
// @ID get_tenderSupplier_list
// @Router /v1/tender-supplier [GET]
// @Summary Get TenderSupplier list
// @Description Get TenderSupplier list
// @Tags TenderSupplier
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListTenderSupplierRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListTenderSupplierResponse} "TenderSupplierBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetTenderSupplierList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().TenderSupplier().GetListTenderSupplier(
		context.Background(),
		&storehouse_service.GetListTenderSupplierRequest{
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

// UpdateTenderSupplier godoc
// @Security ApiKeyAuth
// @ID update_tenderSupplier
// @Router /v1/tender-supplier [PUT]
// @Summary Update TenderSupplier
// @Description Update TenderSupplier
// @Tags TenderSupplier
// @Accept json
// @Produce json
// @Param TenderSupplier body storehouse_service.UpdateTenderSupplierRequest true "UpdateTenderSupplierRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.TenderSupplier} "TenderSupplier data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateTenderSupplier(c *gin.Context) {

	var updateTenderSupplier storehouse_service.UpdateTenderSupplierRequest
	err := c.ShouldBindJSON(&updateTenderSupplier)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().TenderSupplier().UpdateTenderSupplier(
		context.Background(),
		&updateTenderSupplier,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteTenderSupplier godoc
// @Security ApiKeyAuth
// @ID delete_tenderSupplier
// @Router /v1/tender-supplier/{tender_supplier_id} [DELETE]
// @Summary Delete TenderSupplier
// @Description Delete TenderSupplier
// @Tags TenderSupplier
// @Accept json
// @Produce json
// @Param tender_supplier_id path string true "tender_supplier_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteTenderSupplier(c *gin.Context) {

	var tenderSupplierId = c.Param("tender_supplier_id")
	if !util.IsValidUUID(tenderSupplierId) {
		h.HandleResponse(c, status_http.InvalidArgument, "tenderSupplier id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().TenderSupplier().DeleteTenderSupplier(
		context.Background(),
		&storehouse_service.TenderSupplierPrimaryKey{Id: tenderSupplierId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
