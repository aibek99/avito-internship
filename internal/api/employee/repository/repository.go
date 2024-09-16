package repository

import (
	"context"

	employeeInterface "avito-internship/internal/api/employee"
	"avito-internship/internal/connection"

	"github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
)

type EmployeeRepository struct {
	txGetter *sqlx.CtxGetter
	db       connection.DB
}

var (
	_ employeeInterface.Repository = (*EmployeeRepository)(nil)
)

func NewEmployeeRepository(txGetter *sqlx.CtxGetter, db connection.DB) *EmployeeRepository {
	return &EmployeeRepository{
		txGetter: txGetter,
		db:       db,
	}
}

func (e *EmployeeRepository) Ping(ctx context.Context) (*string, error) {
	//query := `UPDATE user SET username = :username WHERE user_id = :user_id;`

	//tr := e.txGetter.DefaultTrOrDB(ctx, e.db.GetPool())
	//err := tr.GetContext(
	//	ctx,
	//	&employee, // struct
	//	query,
	//	args...,
	//)

	err := e.db.GetPool().Ping()
	if err != nil {
		return nil, err
	}
	response := "ok"
	return &response, nil
}
