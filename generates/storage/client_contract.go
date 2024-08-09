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

type ClientContractRepo struct {
	db *pgxpool.Pool
}

func NewClientContractRepo(db *pgxpool.Pool) storage.ClientContractRepoI {
	return &ClientContractRepo{
		db: db,
	}
}

func (c *ClientContractRepo) Create(ctx context.Context, req *storehouse_client_service.CreateClientContractRequest) (resp *storehouse_client_service.ClientContractPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "client_contract" (
			id,
			from_date,
			to_date,
			total_amount,
			file,
			description,
			client_id,
			cashier_request_id,
			status,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetFromDate(),
		req.GetToDate(),
		req.GetTotalAmount(),
		req.GetFile(),
		req.GetDescription(),
		req.GetClientId(),
		req.GetCashierRequestId(),
		req.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_client_service.ClientContractPrimaryKey{Id: id.String()}, nil
}

func (c *ClientContractRepo) GetByPKey(ctx context.Context, req *storehouse_client_service.ClientContractPrimaryKey) (resp *storehouse_client_service.ClientContract, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			from_date,
			to_date,
			total_amount,
			file,
			description,
			client_id,
			cashier_request_id,
			status,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "client_contract"
		WHERE id = $1
	`

	var (
		id        sql.NullString
	fromDate sql.NullString
	toDate sql.NullString
	totalAmount sql.NullString
	file sql.NullString
	description sql.NullString
	clientId sql.NullString
	cashierRequestId sql.NullString
	status sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&fromDate,
		&toDate,
		&totalAmount,
		&file,
		&description,
		&clientId,
		&cashierRequestId,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_client_service.ClientContract{
		Id:        id.String,
		FromDate: fromDate.String,
		ToDate: toDate.String,
		TotalAmount: totalAmount.String,
		File: file.String,
		Description: description.String,
		ClientId: clientId.String,
		CashierRequestId: cashierRequestId.String,
		Status: status.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *ClientContractRepo) GetAll(ctx context.Context, req *storehouse_client_service.GetListClientContractRequest) (resp *storehouse_client_service.GetListClientContractResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_client_service.GetListClientContractResponse{}

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
			from_date,
			to_date,
			total_amount,
			file,
			description,
			client_id,
			cashier_request_id,
			status,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "client_contract"
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
	fromDate sql.NullString
	toDate sql.NullString
	totalAmount sql.NullString
	file sql.NullString
	description sql.NullString
	clientId sql.NullString
	cashierRequestId sql.NullString
	status sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&fromDate,
		&toDate,
		&totalAmount,
		&file,
		&description,
		&clientId,
		&cashierRequestId,
		&status,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.ClientContracts = append(resp.ClientContracts, &storehouse_client_service.ClientContract{
			Id:        id.String,
		FromDate: fromDate.String,
		ToDate: toDate.String,
		TotalAmount: totalAmount.String,
		File: file.String,
		Description: description.String,
		ClientId: clientId.String,
		CashierRequestId: cashierRequestId.String,
		Status: status.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ClientContractRepo) Update(ctx context.Context, req *storehouse_client_service.UpdateClientContractRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"client_contract"
		SET
			from_date = :from_date,
			to_date = :to_date,
			total_amount = :total_amount,
			file = :file,
			description = :description,
			client_id = :client_id,
			cashier_request_id = :cashier_request_id,
			status = :status,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"from_date": req.GetFromDate(),
		"to_date": req.GetToDate(),
		"total_amount": req.GetTotalAmount(),
		"file": req.GetFile(),
		"description": req.GetDescription(),
		"client_id": req.GetClientId(),
		"cashier_request_id": req.GetCashierRequestId(),
		"status": req.GetStatus(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *ClientContractRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"client_contract"
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

func (c *ClientContractRepo) Delete(ctx context.Context, req *storehouse_client_service.ClientContractPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "client_contract" WHERE id = $1`, req.Id)
	return err
}
