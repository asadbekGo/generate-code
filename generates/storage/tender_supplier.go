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

type TenderSupplierRepo struct {
	db *pgxpool.Pool
}

func NewTenderSupplierRepo(db *pgxpool.Pool) storage.TenderSupplierRepoI {
	return &TenderSupplierRepo{
		db: db,
	}
}

func (c *TenderSupplierRepo) Create(ctx context.Context, req *storehouse_supplier_service.CreateTenderSupplierRequest) (resp *storehouse_supplier_service.TenderSupplierPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "tender_supplier" (
			id,
			tender_supplier_number,
			price,
			supplier_balance,
			currency,
			metrics,
			type,
			description,
			supplier_id,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetTenderSupplierNumber(),
		req.GetPrice(),
		req.GetSupplierBalance(),
		req.GetCurrency(),
		req.GetMetrics(),
		req.GetType(),
		req.GetDescription(),
		req.GetSupplierId(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_supplier_service.TenderSupplierPrimaryKey{Id: id.String()}, nil
}

func (c *TenderSupplierRepo) GetByPKey(ctx context.Context, req *storehouse_supplier_service.TenderSupplierPrimaryKey) (resp *storehouse_supplier_service.TenderSupplier, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			tender_supplier_number,
			price,
			supplier_balance,
			currency,
			metrics,
			type,
			description,
			supplier_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "tender_supplier"
		WHERE id = $1
	`

	var (
		id        sql.NullString
	tenderSupplierNumber sql.NullString
	price sql.NullString
	supplierBalance sql.NullString
	currency sql.NullString
	metrics sql.NullString
	type sql.NullString
	description sql.NullString
	supplierId sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&tenderSupplierNumber,
		&price,
		&supplierBalance,
		&currency,
		&metrics,
		&type,
		&description,
		&supplierId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_supplier_service.TenderSupplier{
		Id:        id.String,
		TenderSupplierNumber: tenderSupplierNumber.String,
		Price: price.String,
		SupplierBalance: supplierBalance.String,
		Currency: currency.String,
		Metrics: metrics.String,
		Type: type.String,
		Description: description.String,
		SupplierId: supplierId.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *TenderSupplierRepo) GetAll(ctx context.Context, req *storehouse_supplier_service.GetListTenderSupplierRequest) (resp *storehouse_supplier_service.GetListTenderSupplierResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_supplier_service.GetListTenderSupplierResponse{}

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
			tender_supplier_number,
			price,
			supplier_balance,
			currency,
			metrics,
			type,
			description,
			supplier_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "tender_supplier"
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
	tenderSupplierNumber sql.NullString
	price sql.NullString
	supplierBalance sql.NullString
	currency sql.NullString
	metrics sql.NullString
	type sql.NullString
	description sql.NullString
	supplierId sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&tenderSupplierNumber,
		&price,
		&supplierBalance,
		&currency,
		&metrics,
		&type,
		&description,
		&supplierId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.TenderSuppliers = append(resp.TenderSuppliers, &storehouse_supplier_service.TenderSupplier{
			Id:        id.String,
		TenderSupplierNumber: tenderSupplierNumber.String,
		Price: price.String,
		SupplierBalance: supplierBalance.String,
		Currency: currency.String,
		Metrics: metrics.String,
		Type: type.String,
		Description: description.String,
		SupplierId: supplierId.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *TenderSupplierRepo) Update(ctx context.Context, req *storehouse_supplier_service.UpdateTenderSupplierRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"tender_supplier"
		SET
			tender_supplier_number = :tender_supplier_number,
			price = :price,
			supplier_balance = :supplier_balance,
			currency = :currency,
			metrics = :metrics,
			type = :type,
			description = :description,
			supplier_id = :supplier_id,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"tender_supplier_number": req.GetTenderSupplierNumber(),
		"price": req.GetPrice(),
		"supplier_balance": req.GetSupplierBalance(),
		"currency": req.GetCurrency(),
		"metrics": req.GetMetrics(),
		"type": req.GetType(),
		"description": req.GetDescription(),
		"supplier_id": req.GetSupplierId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *TenderSupplierRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"tender_supplier"
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

func (c *TenderSupplierRepo) Delete(ctx context.Context, req *storehouse_supplier_service.TenderSupplierPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "tender_supplier" WHERE id = $1`, req.Id)
	return err
}
