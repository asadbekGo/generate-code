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

type ClientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) storage.ClientRepoI {
	return &ClientRepo{
		db: db,
	}
}

func (c *ClientRepo) Create(ctx context.Context, req *storehouse_client_service.CreateClientRequest) (resp *storehouse_client_service.ClientPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "client" (
			id,
			first_name,
			last_name,
			birthday,
			balance,
			currency,
			phone_number,
			address,
			status,
			description,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetFirstName(),
		req.GetLastName(),
		req.GetBirthday(),
		req.GetBalance(),
		req.GetCurrency(),
		req.GetPhoneNumber(),
		req.GetAddress(),
		req.GetStatus(),
		req.GetDescription(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_client_service.ClientPrimaryKey{Id: id.String()}, nil
}

func (c *ClientRepo) GetByPKey(ctx context.Context, req *storehouse_client_service.ClientPrimaryKey) (resp *storehouse_client_service.Client, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			first_name,
			last_name,
			birthday,
			balance,
			currency,
			phone_number,
			address,
			status,
			description,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "client"
		WHERE id = $1
	`

	var (
		id        sql.NullString
	firstName sql.NullString
	lastName sql.NullString
	birthday sql.NullString
	balance sql.NullString
	currency sql.NullString
	phoneNumber sql.NullString
	address sql.NullString
	status sql.NullString
	description sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&firstName,
		&lastName,
		&birthday,
		&balance,
		&currency,
		&phoneNumber,
		&address,
		&status,
		&description,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_client_service.Client{
		Id:        id.String,
		FirstName: firstName.String,
		LastName: lastName.String,
		Birthday: birthday.String,
		Balance: balance.String,
		Currency: currency.String,
		PhoneNumber: phoneNumber.String,
		Address: address.String,
		Status: status.String,
		Description: description.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *ClientRepo) GetAll(ctx context.Context, req *storehouse_client_service.GetListClientRequest) (resp *storehouse_client_service.GetListClientResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_client_service.GetListClientResponse{}

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
			first_name,
			last_name,
			birthday,
			balance,
			currency,
			phone_number,
			address,
			status,
			description,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "client"
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
	firstName sql.NullString
	lastName sql.NullString
	birthday sql.NullString
	balance sql.NullString
	currency sql.NullString
	phoneNumber sql.NullString
	address sql.NullString
	status sql.NullString
	description sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&firstName,
		&lastName,
		&birthday,
		&balance,
		&currency,
		&phoneNumber,
		&address,
		&status,
		&description,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Clients = append(resp.Clients, &storehouse_client_service.Client{
			Id:        id.String,
		FirstName: firstName.String,
		LastName: lastName.String,
		Birthday: birthday.String,
		Balance: balance.String,
		Currency: currency.String,
		PhoneNumber: phoneNumber.String,
		Address: address.String,
		Status: status.String,
		Description: description.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ClientRepo) Update(ctx context.Context, req *storehouse_client_service.UpdateClientRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"client"
		SET
			first_name = :first_name,
			last_name = :last_name,
			birthday = :birthday,
			balance = :balance,
			currency = :currency,
			phone_number = :phone_number,
			address = :address,
			status = :status,
			description = :description,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"first_name": req.GetFirstName(),
		"last_name": req.GetLastName(),
		"birthday": req.GetBirthday(),
		"balance": req.GetBalance(),
		"currency": req.GetCurrency(),
		"phone_number": req.GetPhoneNumber(),
		"address": req.GetAddress(),
		"status": req.GetStatus(),
		"description": req.GetDescription(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *ClientRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"client"
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

func (c *ClientRepo) Delete(ctx context.Context, req *storehouse_client_service.ClientPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "client" WHERE id = $1`, req.Id)
	return err
}
