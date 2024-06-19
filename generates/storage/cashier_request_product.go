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

type CashierRequestProductRepo struct {
	db *pgxpool.Pool
}

func NewCashierRequestProductRepo(db *pgxpool.Pool) storage.CashierRequestProductRepoI {
	return &CashierRequestProductRepo{
		db: db,
	}
}

func (c *CashierRequestProductRepo) Create(ctx context.Context, req *storehouse_supplier_service.CreateCashierRequestProductRequest) (resp *storehouse_supplier_service.CashierRequestProductPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "cashier_request_product" (
			id,
			barcode,
			product_number,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			price,
			metrics,
			status,
			product_id,
			cashier_request,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetBarcode(),
		req.GetProductNumber(),
		req.GetQuantity(),
		req.GetQuantityType(),
		req.GetSizeType(),
		req.GetSizeValue(),
		req.GetWeightType(),
		req.GetWeightValue(),
		req.GetPrice(),
		req.GetMetrics(),
		req.GetStatus(),
		req.GetProductId(),
		req.GetCashierRequest(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_supplier_service.CashierRequestProductPrimaryKey{Id: id.String()}, nil
}

func (c *CashierRequestProductRepo) GetByPKey(ctx context.Context, req *storehouse_supplier_service.CashierRequestProductPrimaryKey) (resp *storehouse_supplier_service.CashierRequestProduct, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			barcode,
			product_number,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			price,
			metrics,
			status,
			product_id,
			cashier_request,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "cashier_request_product"
		WHERE id = $1
	`

	var (
		id        sql.NullString
	barcode sql.NullString
	productNumber sql.NullString
	quantity sql.NullString
	quantityType sql.NullString
	sizeType sql.NullString
	sizeValue sql.NullString
	weightType sql.NullString
	weightValue sql.NullString
	price sql.NullString
	metrics sql.NullString
	status sql.NullString
	productId sql.NullString
	cashierRequest sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&barcode,
		&productNumber,
		&quantity,
		&quantityType,
		&sizeType,
		&sizeValue,
		&weightType,
		&weightValue,
		&price,
		&metrics,
		&status,
		&productId,
		&cashierRequest,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_supplier_service.CashierRequestProduct{
		Id:        id.String,
		Barcode: barcode.String,
		ProductNumber: productNumber.String,
		Quantity: quantity.String,
		QuantityType: quantityType.String,
		SizeType: sizeType.String,
		SizeValue: sizeValue.String,
		WeightType: weightType.String,
		WeightValue: weightValue.String,
		Price: price.String,
		Metrics: metrics.String,
		Status: status.String,
		ProductId: productId.String,
		CashierRequest: cashierRequest.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *CashierRequestProductRepo) GetAll(ctx context.Context, req *storehouse_supplier_service.GetListCashierRequestProductRequest) (resp *storehouse_supplier_service.GetListCashierRequestProductResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_supplier_service.GetListCashierRequestProductResponse{}

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
			barcode,
			product_number,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			price,
			metrics,
			status,
			product_id,
			cashier_request,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "cashier_request_product"
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
	barcode sql.NullString
	productNumber sql.NullString
	quantity sql.NullString
	quantityType sql.NullString
	sizeType sql.NullString
	sizeValue sql.NullString
	weightType sql.NullString
	weightValue sql.NullString
	price sql.NullString
	metrics sql.NullString
	status sql.NullString
	productId sql.NullString
	cashierRequest sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&barcode,
		&productNumber,
		&quantity,
		&quantityType,
		&sizeType,
		&sizeValue,
		&weightType,
		&weightValue,
		&price,
		&metrics,
		&status,
		&productId,
		&cashierRequest,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.CashierRequestProducts = append(resp.CashierRequestProducts, &storehouse_supplier_service.CashierRequestProduct{
			Id:        id.String,
		Barcode: barcode.String,
		ProductNumber: productNumber.String,
		Quantity: quantity.String,
		QuantityType: quantityType.String,
		SizeType: sizeType.String,
		SizeValue: sizeValue.String,
		WeightType: weightType.String,
		WeightValue: weightValue.String,
		Price: price.String,
		Metrics: metrics.String,
		Status: status.String,
		ProductId: productId.String,
		CashierRequest: cashierRequest.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *CashierRequestProductRepo) Update(ctx context.Context, req *storehouse_supplier_service.UpdateCashierRequestProductRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"cashier_request_product"
		SET
			barcode = :barcode,
			product_number = :product_number,
			quantity = :quantity,
			quantity_type = :quantity_type,
			size_type = :size_type,
			size_value = :size_value,
			weight_type = :weight_type,
			weight_value = :weight_value,
			price = :price,
			metrics = :metrics,
			status = :status,
			product_id = :product_id,
			cashier_request = :cashier_request,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"barcode": req.GetBarcode(),
		"product_number": req.GetProductNumber(),
		"quantity": req.GetQuantity(),
		"quantity_type": req.GetQuantityType(),
		"size_type": req.GetSizeType(),
		"size_value": req.GetSizeValue(),
		"weight_type": req.GetWeightType(),
		"weight_value": req.GetWeightValue(),
		"price": req.GetPrice(),
		"metrics": req.GetMetrics(),
		"status": req.GetStatus(),
		"product_id": req.GetProductId(),
		"cashier_request": req.GetCashierRequest(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *CashierRequestProductRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"cashier_request_product"
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

func (c *CashierRequestProductRepo) Delete(ctx context.Context, req *storehouse_supplier_service.CashierRequestProductPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "cashier_request_product" WHERE id = $1`, req.Id)
	return err
}
