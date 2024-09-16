package connection

import (
	"context"
	"database/sql"
	"fmt"

	"avito-internship/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // import the pq driver
)

// DBops is
var _ DB = (*Database)(nil)

// DB is
type DB interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Close() error
	GetPool() *sqlx.DB
}

// GenerateDsn is
func GenerateDsn(cfgs *config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfgs.Host, cfgs.Port, cfgs.User, cfgs.Password, cfgs.DBName)
}

// NewDB is
func NewDB(cfgs *config.Config) (*Database, error) {
	db, err := sqlx.Connect("postgres", GenerateDsn(cfgs))

	if err != nil {
		return nil, fmt.Errorf("sqlx.Connect: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return &Database{Db: db}, nil
}

// Database is
type Database struct {
	Db *sqlx.DB
}

// Get is
func (d *Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.Db.GetContext(ctx, dest, query, args...)
}

// QueryRow is
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.Db.QueryRowContext(ctx, query, args...)
}

// Query is
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.Db.QueryContext(ctx, query, args...)
}

// Select is
func (d *Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.Db.SelectContext(ctx, dest, query, args...)
}

// Execute is
func (d *Database) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.Db.ExecContext(ctx, query, args...)
}

// Close is
func (d *Database) Close() error {
	return d.Db.Close()
}

func (d *Database) GetPool() *sqlx.DB {
	return d.Db
}
