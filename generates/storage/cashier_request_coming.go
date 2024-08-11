package client_storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"

	"warehouse/warehouse_go_storehouse_service/genproto/storehouse_client_service"
	"warehouse/warehouse_go_storehouse_service/models"
	"warehouse/warehouse_go_storehouse_service/pkg/helper"
	"warehouse/warehouse_go_storehouse_service/storage"
)

type CashierRequestComingRepo struct {
	db *pgxpool.Pool
}

func NewCashierRequestComingRepo(db *pgxpool.Pool) storage.CashierRequestComingRepoI {
	return &CashierRequestComingRepo{
		db: db,
	}
}

func (c *CashierRequestComingRepo) Create(ctx context.Context, req *storehouse_client_service.CreateCashierRequestComingRequest) (resp *storehouse_client_service.CashierRequestComingPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "cashier_request_coming" (
			id,
			date_time,
			cashier_request_coming_number,
			term_payment,
			term_amount,
			currency,
			description,
			file,
			client_id,
			user_id,
			status,
			type_price,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetDateTime(),
		req.GetCashierRequestComingNumber(),
		req.GetTermPayment(),
		req.GetTermAmount(),
		req.GetCurrency(),
		req.GetDescription(),
		req.GetFile(),
		req.GetClientId(),
		req.GetUserId(),
		req.GetStatus(),
		req.GetTypePrice(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_client_service.CashierRequestComingPrimaryKey{Id: id.String()}, nil
}

func (c *CashierRequestComingRepo) GetByPKey(ctx context.Context, req *storehouse_client_service.CashierRequestComingPrimaryKey) (resp *storehouse_client_service.CashierRequestComing, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			date_time,
			cashier_request_coming_number,
			term_payment,
			term_amount,
			currency,
			description,
			file,
			client_id,
			user_id,
			status,
			type_price,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "cashier_request_coming"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		dateTime sql.NullString
		cashierRequestComingNumber sql.NullString
		termPayment sql.NullString
		termAmount sql.NullFloat64
		currency sql.NullString
		description sql.NullString
		file sql.NullString
		clientId sql.NullString
		userId sql.NullString
		status sql.NullString
		typePrice sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&dateTime,
		&cashierRequestComingNumber,
		&termPayment,
		&termAmount,
		&currency,
		&description,
		&file,
		&clientId,
		&userId,
		&status,
		&typePrice,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_client_service.CashierRequestComing{
		Id:        id.String,
		DateTime: dateTime.String,
		CashierRequestComingNumber: cashierRequestComingNumber.String,
		TermPayment: termPayment.String,
		TermAmount: termAmount.Float64,
		Currency: currency.String,
		Description: description.String,
		File: file.String,
		ClientId: clientId.String,
		UserId: userId.String,
		Status: status.String,
		TypePrice: typePrice.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *CashierRequestComingRepo) GetAll(ctx context.Context, req *storehouse_client_service.GetListCashierRequestComingRequest) (resp *storehouse_client_service.GetListCashierRequestComingResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_client_service.GetListCashierRequestComingResponse{}

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
			date_time,
			cashier_request_coming_number,
			term_payment,
			term_amount,
			currency,
			description,
			file,
			client_id,
			user_id,
			status,
			type_price,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "cashier_request_coming"
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
		dateTime sql.NullString
		cashierRequestComingNumber sql.NullString
		termPayment sql.NullString
		termAmount sql.NullFloat64
		currency sql.NullString
		description sql.NullString
		file sql.NullString
		clientId sql.NullString
		userId sql.NullString
		status sql.NullString
		typePrice sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&dateTime,
		&cashierRequestComingNumber,
		&termPayment,
		&termAmount,
		&currency,
		&description,
		&file,
		&clientId,
		&userId,
		&status,
		&typePrice,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.CashierRequestComings = append(resp.CashierRequestComings, &storehouse_client_service.CashierRequestComing{
			Id:        id.String,
		DateTime: dateTime.String,
		CashierRequestComingNumber: cashierRequestComingNumber.String,
		TermPayment: termPayment.String,
		TermAmount: termAmount.Float64,
		Currency: currency.String,
		Description: description.String,
		File: file.String,
		ClientId: clientId.String,
		UserId: userId.String,
		Status: status.String,
		TypePrice: typePrice.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *CashierRequestComingRepo) Update(ctx context.Context, req *storehouse_client_service.UpdateCashierRequestComingRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"cashier_request_coming"
		SET
			date_time = :date_time,
			cashier_request_coming_number = :cashier_request_coming_number,
			term_payment = :term_payment,
			term_amount = :term_amount,
			currency = :currency,
			description = :description,
			file = :file,
			client_id = :client_id,
			user_id = :user_id,
			status = :status,
			type_price = :type_price,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"date_time": req.GetDateTime(),
		"cashier_request_coming_number": req.GetCashierRequestComingNumber(),
		"term_payment": req.GetTermPayment(),
		"term_amount": req.GetTermAmount(),
		"currency": req.GetCurrency(),
		"description": req.GetDescription(),
		"file": req.GetFile(),
		"client_id": req.GetClientId(),
		"user_id": req.GetUserId(),
		"status": req.GetStatus(),
		"type_price": req.GetTypePrice(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *CashierRequestComingRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"cashier_request_coming"
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

func (c *CashierRequestComingRepo) Delete(ctx context.Context, req *storehouse_client_service.CashierRequestComingPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "cashier_request_coming" WHERE id = $1`, req.Id)
	return err
}
