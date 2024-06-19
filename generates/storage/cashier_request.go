package supplier_storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"

	"warehouse/warehouse_go_storehouse_service/genproto/storehouse_supplier_service"
	"warehouse/warehouse_go_storehouse_service/models"
	"warehouse/warehouse_go_storehouse_service/pkg/helper"
	"warehouse/warehouse_go_storehouse_service/storage"
)

type CashierRequestRepo struct {
	db *pgxpool.Pool
}

func NewCashierRequestRepo(db *pgxpool.Pool) storage.CashierRequestRepoI {
	return &CashierRequestRepo{
		db: db,
	}
}

func (c *CashierRequestRepo) Create(ctx context.Context, req *storehouse_supplier_service.CreateCashierRequestRequest) (resp *storehouse_supplier_service.CashierRequestPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "cashier_request" (
			id,
			cashier_request_number,
			term_payment,
			term_amount,
			currency,
			description,
			file,
			supplier_id,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetCashierRequestNumber(),
		req.GetTermPayment(),
		req.GetTermAmount(),
		req.GetCurrency(),
		req.GetDescription(),
		req.GetFile(),
		req.GetSupplierId(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_supplier_service.CashierRequestPrimaryKey{Id: id.String()}, nil
}

func (c *CashierRequestRepo) GetByPKey(ctx context.Context, req *storehouse_supplier_service.CashierRequestPrimaryKey) (resp *storehouse_supplier_service.CashierRequest, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			cashier_request_number,
			term_payment,
			term_amount,
			currency,
			description,
			file,
			supplier_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "cashier_request"
		WHERE id = $1
	`

	var (
		id        sql.NullString
	cashierRequestNumber sql.NullString
	termPayment sql.NullString
	termAmount sql.NullString
	currency sql.NullString
	description sql.NullString
	file sql.NullString
	supplierId sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&cashierRequestNumber,
		&termPayment,
		&termAmount,
		&currency,
		&description,
		&file,
		&supplierId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_supplier_service.CashierRequest{
		Id:        id.String,
		CashierRequestNumber: cashierRequestNumber.String,
		TermPayment: termPayment.String,
		TermAmount: termAmount.String,
		Currency: currency.String,
		Description: description.String,
		File: file.String,
		SupplierId: supplierId.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *CashierRequestRepo) GetAll(ctx context.Context, req *storehouse_supplier_service.GetListCashierRequestRequest) (resp *storehouse_supplier_service.GetListCashierRequestResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_supplier_service.GetListCashierRequestResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE  "
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			cashier_request_number,
			term_payment,
			term_amount,
			currency,
			description,
			file,
			supplier_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "cashier_request"
	`

	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}

	if req.GetPage() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = (req.Page - 1) * req.Limit
	}

	// if req.GetSearch() != "" {
	// 	filter += ` AND  (CONCAT(name::varchar) ILIKE '%' || :search || '%' )`
	// 	params["search"] = req.Search
	// }

	if req.GetWhereQuery() != "" {
		filter += req.WhereQuery
	}

	for key, val := range req.Filters.AsMap() {
		if !helper.CheckTypeAndEmpty(val) {
			filter += fmt.Sprintf(` AND  %s = :%s`, key, key)
			params[key] = val
		}
	}

	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        sql.NullString
	cashierRequestNumber sql.NullString
	termPayment sql.NullString
	termAmount sql.NullString
	currency sql.NullString
	description sql.NullString
	file sql.NullString
	supplierId sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&cashierRequestNumber,
		&termPayment,
		&termAmount,
		&currency,
		&description,
		&file,
		&supplierId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.CashierRequests = append(resp.CashierRequests, &storehouse_supplier_service.CashierRequest{
			Id:        id.String,
		CashierRequestNumber: cashierRequestNumber.String,
		TermPayment: termPayment.String,
		TermAmount: termAmount.String,
		Currency: currency.String,
		Description: description.String,
		File: file.String,
		SupplierId: supplierId.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *CashierRequestRepo) Update(ctx context.Context, req *storehouse_supplier_service.UpdateCashierRequestRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"cashier_request"
		SET
			cashier_request_number = :cashier_request_number,
			term_payment = :term_payment,
			term_amount = :term_amount,
			currency = :currency,
			description = :description,
			file = :file,
			supplier_id = :supplier_id,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"cashier_request_number": req.GetCashierRequestNumber(),
		"term_payment": req.GetTermPayment(),
		"term_amount": req.GetTermAmount(),
		"currency": req.GetCurrency(),
		"description": req.GetDescription(),
		"file": req.GetFile(),
		"supplier_id": req.GetSupplierId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *CashierRequestRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.UpdatePatch")
	defer dbSpan.Finish()

	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Fields) == 0 {
		err = errors.New("no updates provided")
		return
	}

	req.Fields["id"] = req.Id

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Fields)-1 {
			set += ", "
		}
		ind++
	}

	query = `
		UPDATE
			"cashier_request"
	` + set + ` , updated_at = now()
		WHERE
			id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), err
}

func (c *CashierRequestRepo) Delete(ctx context.Context, req *storehouse_supplier_service.CashierRequestPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "cashier_request" WHERE id = $1`, req.Id)
	return err
}
