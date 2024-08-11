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

type CashierRequestComingProductRepo struct {
	db *pgxpool.Pool
}

func NewCashierRequestComingProductRepo(db *pgxpool.Pool) storage.CashierRequestComingProductRepoI {
	return &CashierRequestComingProductRepo{
		db: db,
	}
}

func (c *CashierRequestComingProductRepo) Create(ctx context.Context, req *storehouse_client_service.CreateCashierRequestComingProductRequest) (resp *storehouse_client_service.CashierRequestComingProductPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "cashier_request_coming_product" (
			id,
			date_time,
			price,
			total_price,
			currency,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			send_value,
			status,
			user_id,
			product_id,
			cashier_request_coming_id,
			metrics,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetDateTime(),
		req.GetPrice(),
		req.GetTotalPrice(),
		req.GetCurrency(),
		req.GetQuantity(),
		req.GetQuantityType(),
		req.GetSizeType(),
		req.GetSizeValue(),
		req.GetWeightType(),
		req.GetWeightValue(),
		req.GetSendValue(),
		req.GetStatus(),
		req.GetUserId(),
		req.GetProductId(),
		req.GetCashierRequestComingId(),
		req.GetMetrics(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_client_service.CashierRequestComingProductPrimaryKey{Id: id.String()}, nil
}

func (c *CashierRequestComingProductRepo) GetByPKey(ctx context.Context, req *storehouse_client_service.CashierRequestComingProductPrimaryKey) (resp *storehouse_client_service.CashierRequestComingProduct, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			date_time,
			price,
			total_price,
			currency,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			send_value,
			status,
			user_id,
			product_id,
			cashier_request_coming_id,
			metrics,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "cashier_request_coming_product"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		dateTime sql.NullString
		price sql.NullFloat64
		totalPrice sql.NullFloat64
		currency sql.NullString
		quantity sql.NullInt64
		quantityType sql.NullString
		sizeType sql.NullString
		sizeValue sql.NullFloat64
		weightType sql.NullString
		weightValue sql.NullFloat64
		sendValue sql.NullFloat64
		status sql.NullString
		userId sql.NullString
		productId sql.NullString
		cashierRequestComingId sql.NullString
		metrics sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&dateTime,
		&price,
		&totalPrice,
		&currency,
		&quantity,
		&quantityType,
		&sizeType,
		&sizeValue,
		&weightType,
		&weightValue,
		&sendValue,
		&status,
		&userId,
		&productId,
		&cashierRequestComingId,
		&metrics,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_client_service.CashierRequestComingProduct{
		Id:        id.String,
		DateTime: dateTime.String,
		Price: price.Float64,
		TotalPrice: totalPrice.Float64,
		Currency: currency.String,
		Quantity: quantity.Int64,
		QuantityType: quantityType.String,
		SizeType: sizeType.String,
		SizeValue: sizeValue.Float64,
		WeightType: weightType.String,
		WeightValue: weightValue.Float64,
		SendValue: sendValue.Float64,
		Status: status.String,
		UserId: userId.String,
		ProductId: productId.String,
		CashierRequestComingId: cashierRequestComingId.String,
		Metrics: metrics.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *CashierRequestComingProductRepo) GetAll(ctx context.Context, req *storehouse_client_service.GetListCashierRequestComingProductRequest) (resp *storehouse_client_service.GetListCashierRequestComingProductResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_client_service.GetListCashierRequestComingProductResponse{}

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
			price,
			total_price,
			currency,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			send_value,
			status,
			user_id,
			product_id,
			cashier_request_coming_id,
			metrics,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "cashier_request_coming_product"
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
		price sql.NullFloat64
		totalPrice sql.NullFloat64
		currency sql.NullString
		quantity sql.NullInt64
		quantityType sql.NullString
		sizeType sql.NullString
		sizeValue sql.NullFloat64
		weightType sql.NullString
		weightValue sql.NullFloat64
		sendValue sql.NullFloat64
		status sql.NullString
		userId sql.NullString
		productId sql.NullString
		cashierRequestComingId sql.NullString
		metrics sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&dateTime,
		&price,
		&totalPrice,
		&currency,
		&quantity,
		&quantityType,
		&sizeType,
		&sizeValue,
		&weightType,
		&weightValue,
		&sendValue,
		&status,
		&userId,
		&productId,
		&cashierRequestComingId,
		&metrics,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.CashierRequestComingProducts = append(resp.CashierRequestComingProducts, &storehouse_client_service.CashierRequestComingProduct{
			Id:        id.String,
		DateTime: dateTime.String,
		Price: price.Float64,
		TotalPrice: totalPrice.Float64,
		Currency: currency.String,
		Quantity: quantity.Int64,
		QuantityType: quantityType.String,
		SizeType: sizeType.String,
		SizeValue: sizeValue.Float64,
		WeightType: weightType.String,
		WeightValue: weightValue.Float64,
		SendValue: sendValue.Float64,
		Status: status.String,
		UserId: userId.String,
		ProductId: productId.String,
		CashierRequestComingId: cashierRequestComingId.String,
		Metrics: metrics.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *CashierRequestComingProductRepo) Update(ctx context.Context, req *storehouse_client_service.UpdateCashierRequestComingProductRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"cashier_request_coming_product"
		SET
			date_time = :date_time,
			price = :price,
			total_price = :total_price,
			currency = :currency,
			quantity = :quantity,
			quantity_type = :quantity_type,
			size_type = :size_type,
			size_value = :size_value,
			weight_type = :weight_type,
			weight_value = :weight_value,
			send_value = :send_value,
			status = :status,
			user_id = :user_id,
			product_id = :product_id,
			cashier_request_coming_id = :cashier_request_coming_id,
			metrics = :metrics,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"date_time": req.GetDateTime(),
		"price": req.GetPrice(),
		"total_price": req.GetTotalPrice(),
		"currency": req.GetCurrency(),
		"quantity": req.GetQuantity(),
		"quantity_type": req.GetQuantityType(),
		"size_type": req.GetSizeType(),
		"size_value": req.GetSizeValue(),
		"weight_type": req.GetWeightType(),
		"weight_value": req.GetWeightValue(),
		"send_value": req.GetSendValue(),
		"status": req.GetStatus(),
		"user_id": req.GetUserId(),
		"product_id": req.GetProductId(),
		"cashier_request_coming_id": req.GetCashierRequestComingId(),
		"metrics": req.GetMetrics(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *CashierRequestComingProductRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"cashier_request_coming_product"
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

func (c *CashierRequestComingProductRepo) Delete(ctx context.Context, req *storehouse_client_service.CashierRequestComingProductPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "cashier_request_coming_product" WHERE id = $1`, req.Id)
	return err
}
