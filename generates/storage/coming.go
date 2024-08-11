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

type ComingRepo struct {
	db *pgxpool.Pool
}

func NewComingRepo(db *pgxpool.Pool) storage.ComingRepoI {
	return &ComingRepo{
		db: db,
	}
}

func (c *ComingRepo) Create(ctx context.Context, req *storehouse_client_service.CreateComingRequest) (resp *storehouse_client_service.ComingPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "coming" (
			id,
			name,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			price,
			total_price,
			currency,
			date_time,
			client_id,
			client_contract_id,
			product_id,
			cashier_request_coming_id,
			user_id,
			description,
			type,
			type_price,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetName(),
		req.GetQuantity(),
		req.GetQuantityType(),
		req.GetSizeType(),
		req.GetSizeValue(),
		req.GetWeightType(),
		req.GetWeightValue(),
		req.GetPrice(),
		req.GetTotalPrice(),
		req.GetCurrency(),
		req.GetDateTime(),
		req.GetClientId(),
		req.GetClientContractId(),
		req.GetProductId(),
		req.GetCashierRequestComingId(),
		req.GetUserId(),
		req.GetDescription(),
		req.GetType(),
		req.GetTypePrice(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_client_service.ComingPrimaryKey{Id: id.String()}, nil
}

func (c *ComingRepo) GetByPKey(ctx context.Context, req *storehouse_client_service.ComingPrimaryKey) (resp *storehouse_client_service.Coming, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			name,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			price,
			total_price,
			currency,
			date_time,
			client_id,
			client_contract_id,
			product_id,
			cashier_request_coming_id,
			user_id,
			description,
			type,
			type_price,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "coming"
		WHERE id = $1
	`

	var (
		id        sql.NullString
		name sql.NullString
		quantity sql.NullInt64
		quantityType sql.NullString
		sizeType sql.NullString
		sizeValue sql.NullFloat64
		weightType sql.NullString
		weightValue sql.NullFloat64
		price sql.NullFloat64
		totalPrice sql.NullFloat64
		currency sql.NullString
		dateTime sql.NullString
		clientId sql.NullString
		clientContractId sql.NullString
		productId sql.NullString
		cashierRequestComingId sql.NullString
		userId sql.NullString
		description sql.NullString
		type sql.NullString
		typePrice sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&quantity,
		&quantityType,
		&sizeType,
		&sizeValue,
		&weightType,
		&weightValue,
		&price,
		&totalPrice,
		&currency,
		&dateTime,
		&clientId,
		&clientContractId,
		&productId,
		&cashierRequestComingId,
		&userId,
		&description,
		&type,
		&typePrice,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_client_service.Coming{
		Id:        id.String,
		Name: name.String,
		Quantity: quantity.Int64,
		QuantityType: quantityType.String,
		SizeType: sizeType.String,
		SizeValue: sizeValue.Float64,
		WeightType: weightType.String,
		WeightValue: weightValue.Float64,
		Price: price.Float64,
		TotalPrice: totalPrice.Float64,
		Currency: currency.String,
		DateTime: dateTime.String,
		ClientId: clientId.String,
		ClientContractId: clientContractId.String,
		ProductId: productId.String,
		CashierRequestComingId: cashierRequestComingId.String,
		UserId: userId.String,
		Description: description.String,
		Type: type.String,
		TypePrice: typePrice.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *ComingRepo) GetAll(ctx context.Context, req *storehouse_client_service.GetListComingRequest) (resp *storehouse_client_service.GetListComingResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_client_service.GetListComingResponse{}

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
			name,
			quantity,
			quantity_type,
			size_type,
			size_value,
			weight_type,
			weight_value,
			price,
			total_price,
			currency,
			date_time,
			client_id,
			client_contract_id,
			product_id,
			cashier_request_coming_id,
			user_id,
			description,
			type,
			type_price,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "coming"
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
		name sql.NullString
		quantity sql.NullInt64
		quantityType sql.NullString
		sizeType sql.NullString
		sizeValue sql.NullFloat64
		weightType sql.NullString
		weightValue sql.NullFloat64
		price sql.NullFloat64
		totalPrice sql.NullFloat64
		currency sql.NullString
		dateTime sql.NullString
		clientId sql.NullString
		clientContractId sql.NullString
		productId sql.NullString
		cashierRequestComingId sql.NullString
		userId sql.NullString
		description sql.NullString
		type sql.NullString
		typePrice sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&name,
		&quantity,
		&quantityType,
		&sizeType,
		&sizeValue,
		&weightType,
		&weightValue,
		&price,
		&totalPrice,
		&currency,
		&dateTime,
		&clientId,
		&clientContractId,
		&productId,
		&cashierRequestComingId,
		&userId,
		&description,
		&type,
		&typePrice,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Comings = append(resp.Comings, &storehouse_client_service.Coming{
			Id:        id.String,
		Name: name.String,
		Quantity: quantity.Int64,
		QuantityType: quantityType.String,
		SizeType: sizeType.String,
		SizeValue: sizeValue.Float64,
		WeightType: weightType.String,
		WeightValue: weightValue.Float64,
		Price: price.Float64,
		TotalPrice: totalPrice.Float64,
		Currency: currency.String,
		DateTime: dateTime.String,
		ClientId: clientId.String,
		ClientContractId: clientContractId.String,
		ProductId: productId.String,
		CashierRequestComingId: cashierRequestComingId.String,
		UserId: userId.String,
		Description: description.String,
		Type: type.String,
		TypePrice: typePrice.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ComingRepo) Update(ctx context.Context, req *storehouse_client_service.UpdateComingRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"coming"
		SET
			name = :name,
			quantity = :quantity,
			quantity_type = :quantity_type,
			size_type = :size_type,
			size_value = :size_value,
			weight_type = :weight_type,
			weight_value = :weight_value,
			price = :price,
			total_price = :total_price,
			currency = :currency,
			date_time = :date_time,
			client_id = :client_id,
			client_contract_id = :client_contract_id,
			product_id = :product_id,
			cashier_request_coming_id = :cashier_request_coming_id,
			user_id = :user_id,
			description = :description,
			type = :type,
			type_price = :type_price,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"name": req.GetName(),
		"quantity": req.GetQuantity(),
		"quantity_type": req.GetQuantityType(),
		"size_type": req.GetSizeType(),
		"size_value": req.GetSizeValue(),
		"weight_type": req.GetWeightType(),
		"weight_value": req.GetWeightValue(),
		"price": req.GetPrice(),
		"total_price": req.GetTotalPrice(),
		"currency": req.GetCurrency(),
		"date_time": req.GetDateTime(),
		"client_id": req.GetClientId(),
		"client_contract_id": req.GetClientContractId(),
		"product_id": req.GetProductId(),
		"cashier_request_coming_id": req.GetCashierRequestComingId(),
		"user_id": req.GetUserId(),
		"description": req.GetDescription(),
		"type": req.GetType(),
		"type_price": req.GetTypePrice(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *ComingRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"coming"
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

func (c *ComingRepo) Delete(ctx context.Context, req *storehouse_client_service.ComingPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "coming" WHERE id = $1`, req.Id)
	return err
}
