package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"

	"githubc.com/asadbekGo/generate-code/config"
	"githubc.com/asadbekGo/generate-code/generates/storage"
)

type Store struct {
	db *pgxpool.Pool
}

type Pool struct {
	db *pgxpool.Pool
}

func (b *Pool) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "pgx.QueryRow")
	defer dbSpan.Finish()

	dbSpan.SetTag("sql", sql)
	dbSpan.SetTag("args", args)

	return b.db.QueryRow(ctx, sql, args...)
}

func (b *Pool) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "pgx.Query")
	defer dbSpan.Finish()

	dbSpan.SetTag("sql", sql)
	dbSpan.SetTag("args", args)

	return b.db.Query(ctx, sql, args...)
}

func (b *Pool) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	dbSpan, ctx := opentracing.StartSpanFromContext(ctx, "pgx.Exec")
	defer dbSpan.Finish()

	dbSpan.SetTag("sql", sql)
	dbSpan.SetTag("args", arguments)

	return b.db.Exec(ctx, sql, arguments...)
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + level + msg
	args = append(args, level, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}
