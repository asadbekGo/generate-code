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

type TenderRepo struct {
	db *pgxpool.Pool
}

func NewTenderRepo(db *pgxpool.Pool) storage.TenderRepoI {
	return &TenderRepo{
		db: db,
	}
}

func (c *TenderRepo) Create(ctx context.Context, req *storehouse_supplier_service.CreateTenderRequest) (resp *storehouse_supplier_service.TenderPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "tender" (
			id,
			name,
			tender_number,
			date_time,
			status,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetName(),
		req.GetTenderNumber(),
		req.GetDateTime(),
		req.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_supplier_service.TenderPrimaryKey{Id: id.String()}, nil
}

func (c *TenderRepo) GetByPKey(ctx context.Context, req *storehouse_supplier_service.TenderPrimaryKey) (resp *storehouse_supplier_service.Tender, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			name,
			tender_number,
			date_time,
			status,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "tender"
		WHERE id = $1
	`

	var (
		id        sql.NullString
	name sql.NullString
	tenderNumber sql.NullString
	dateTime sql.NullString
	status sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&tenderNumber,
		&dateTime,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_supplier_service.Tender{
		Id:        id.String,
		Name: name.String,
		TenderNumber: tenderNumber.String,
		DateTime: dateTime.String,
		Status: status.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *TenderRepo) GetAll(ctx context.Context, req *storehouse_supplier_service.GetListTenderRequest) (resp *storehouse_supplier_service.GetListTenderResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_supplier_service.GetListTenderResponse{}

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
			tender_number,
			date_time,
			status,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "tender"
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
	tenderNumber sql.NullString
	dateTime sql.NullString
	status sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&name,
		&tenderNumber,
		&dateTime,
		&status,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.Tenders = append(resp.Tenders, &storehouse_supplier_service.Tender{
			Id:        id.String,
		Name: name.String,
		TenderNumber: tenderNumber.String,
		DateTime: dateTime.String,
		Status: status.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *TenderRepo) Update(ctx context.Context, req *storehouse_supplier_service.UpdateTenderRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"tender"
		SET
			name = :name,
			tender_number = :tender_number,
			date_time = :date_time,
			status = :status,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"name": req.GetName(),
		"tender_number": req.GetTenderNumber(),
		"date_time": req.GetDateTime(),
		"status": req.GetStatus(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *TenderRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"tender"
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

func (c *TenderRepo) Delete(ctx context.Context, req *storehouse_supplier_service.TenderPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "tender" WHERE id = $1`, req.Id)
	return err
}
