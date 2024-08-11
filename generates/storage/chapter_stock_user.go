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

type ChapterStockUserRepo struct {
	db *pgxpool.Pool
}

func NewChapterStockUserRepo(db *pgxpool.Pool) storage.ChapterStockUserRepoI {
	return &ChapterStockUserRepo{
		db: db,
	}
}

func (c *ChapterStockUserRepo) Create(ctx context.Context, req *storehouse_client_service.CreateChapterStockUserRequest) (resp *storehouse_client_service.ChapterStockUserPrimaryKey, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Create")
	defer dbSpan.Finish()

	var id = uuid.New()

	query := `
		INSERT INTO "chapter_stock_user" (
			id,
			chapter_stock_id,
			user_id,
			updated_at
		)
		VALUES ($1, $2, $3, now())
	`

	_, err = c.db.Exec(ctx,
		query,
		id,
		req.GetChapterStockId(),
		req.GetUserId(),
	)

	if err != nil {
		return nil, err
	}

	return &storehouse_client_service.ChapterStockUserPrimaryKey{Id: id.String()}, nil
}

func (c *ChapterStockUserRepo) GetByPKey(ctx context.Context, req *storehouse_client_service.ChapterStockUserPrimaryKey) (resp *storehouse_client_service.ChapterStockUser, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetByPKey")
	defer dbSpan.Finish()

	query := `
		SELECT
			id,
			chapter_stock_id,
			user_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "chapter_stock_user"
		WHERE id = $1
	`

	var (
		id        sql.NullString
	chapterStockId sql.NullString
	userId sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	err = c.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&chapterStockId,
		&userId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return resp, err
	}

	resp = &storehouse_client_service.ChapterStockUser{
		Id:        id.String,
		ChapterStockId: chapterStockId.String,
		UserId: userId.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}

	return
}

func (c *ChapterStockUserRepo) GetAll(ctx context.Context, req *storehouse_client_service.GetListChapterStockUserRequest) (resp *storehouse_client_service.GetListChapterStockUserResponse, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.GetAll")
	defer dbSpan.Finish()

	resp = &storehouse_client_service.GetListChapterStockUserResponse{}

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
			chapter_stock_id,
			user_id,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "chapter_stock_user"
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
	chapterStockId sql.NullString
	userId sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
		&chapterStockId,
		&userId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return resp, err
		}

		resp.ChapterStockUsers = append(resp.ChapterStockUsers, &storehouse_client_service.ChapterStockUser{
			Id:        id.String,
		ChapterStockId: chapterStockId.String,
		UserId: userId.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return
}

func (c *ChapterStockUserRepo) Update(ctx context.Context, req *storehouse_client_service.UpdateChapterStockUserRequest) (rowsAffected int64, err error) {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Update")
	defer dbSpan.Finish()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE 
			"chapter_stock_user"
		SET
			chapter_stock_id = :chapter_stock_id,
			user_id = :user_id,
			updated_at = now()
		WHERE
			id = :id
	`
	params = map[string]interface{}{
		"id":   req.GetId(),
		"chapter_stock_id": req.GetChapterStockId(),
		"user_id": req.GetUserId(),
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := c.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), nil
}

func (c *ChapterStockUserRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

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
			"chapter_stock_user"
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

func (c *ChapterStockUserRepo) Delete(ctx context.Context, req *storehouse_client_service.ChapterStockUserPrimaryKey) error {

	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "storage.Delete")
	defer dbSpan.Finish()

	_, err := c.db.Exec(ctx, `DELETE FROM "chapter_stock_user" WHERE id = $1`, req.Id)
	return err
}
