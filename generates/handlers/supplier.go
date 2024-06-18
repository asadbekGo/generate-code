package handlers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"warehouse/warehouse_go_api_gateway/api/status_http"
	"warehouse/warehouse_go_api_gateway/genproto/storehouse_service"
	"warehouse/warehouse_go_api_gateway/pkg/util"
)

// CreateSupplier godoc
// @Security ApiKeyAuth
// @ID create_supplier
// @Router /v1/supplier [POST]
// @Summary Create Supplier
// @Description Create Supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param Supplier body storehouse_service.CreateSupplierRequest true "CreateSupplierRequestBody"
// @Success 201 {object} status_http.Response{data=storehouse_service.Supplier} "Supplier data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) CreateSupplier(c *gin.Context) {

	var supplier storehouse_service.CreateSupplierRequest
	err := c.ShouldBindJSON(&supplier)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Supplier().CreateSupplier(
		context.Background(),
		&supplier,
	)
	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Created, response)
}

// GetSingleSupplier godoc
// @Security ApiKeyAuth
// @ID get_supplier_by_id
// @Router /v1/supplier/{supplier_id} [GET]
// @Summary Get single Supplier
// @Description Get single Supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param supplier_id path string true "supplier_id"
// @Success 200 {object} status_http.Response{data=storehouse_service.Supplier} "SupplierBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSingleSupplier(c *gin.Context) {

	var supplierId = c.Param("supplier_id")
	if !util.IsValidUUID(supplierId) {
		h.HandleResponse(c, status_http.InvalidArgument, "supplier id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Supplier().GetByIDSupplier(
		context.Background(),
		&storehouse_service.SupplierPrimaryKey{Id: supplierId},
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

// GetSupplierList godoc
// @Security ApiKeyAuth
// @ID get_supplier_list
// @Router /v1/supplier [GET]
// @Summary Get Supplier list
// @Description Get Supplier list
// @Tags Supplier
// @Accept json
// @Produce json
// @Param filters query storehouse_service.GetListSupplierRequest true "filters"
// @Success 200 {object} status_http.Response{data=storehouse_service.GetListSupplierResponse} "SupplierBody"
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) GetSupplierList(c *gin.Context) {

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

	response, err := h.services.StorehouseService().Supplier().GetListSupplier(
		context.Background(),
		&storehouse_service.GetListSupplierRequest{
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

// UpdateSupplier godoc
// @Security ApiKeyAuth
// @ID update_supplier
// @Router /v1/supplier [PUT]
// @Summary Update Supplier
// @Description Update Supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param Supplier body storehouse_service.UpdateSupplierRequest true "UpdateSupplierRequestBody"
// @Success 200 {object} status_http.Response{data=storehouse_service.Supplier} "Supplier data"
// @Response 400 {object} status_http.Response{data=string} "Bad Request"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) UpdateSupplier(c *gin.Context) {

	var updateSupplier storehouse_service.UpdateSupplierRequest
	err := c.ShouldBindJSON(&updateSupplier)
	if err != nil {
		h.HandleResponse(c, status_http.BadRequest, err.Error())
		return
	}

	response, err := h.services.StorehouseService().Supplier().UpdateSupplier(
		context.Background(),
		&updateSupplier,
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.Accepted, response)
}

// DeleteSupplier godoc
// @Security ApiKeyAuth
// @ID delete_supplier
// @Router /v1/supplier/{supplier_id} [DELETE]
// @Summary Delete Supplier
// @Description Delete Supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param supplier_id path string true "supplier_id"
// @Success 204
// @Response 400 {object} status_http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} status_http.Response{data=string} "Server Error"
func (h *Handler) DeleteSupplier(c *gin.Context) {

	var supplierId = c.Param("supplier_id")
	if !util.IsValidUUID(supplierId) {
		h.HandleResponse(c, status_http.InvalidArgument, "supplier id is an invalid uuid")
		return
	}

	response, err := h.services.StorehouseService().Supplier().DeleteSupplier(
		context.Background(),
		&storehouse_service.SupplierPrimaryKey{Id: supplierId},
	)

	if err != nil {
		h.HandleResponse(c, status_http.GRPCError, err.Error())
		return
	}

	h.HandleResponse(c, status_http.NoContent, response)
}
