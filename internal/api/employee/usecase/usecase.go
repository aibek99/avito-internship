package usecase

import (
	"context"

	"avito-internship/internal/api/employee"

	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

type EmployeeUseCase struct {
	trManager *manager.Manager
	repo      employee.Repository
}

// employee.UseCase

func NewEmployeeUseCase(trManager *manager.Manager, repo employee.Repository) *EmployeeUseCase {
	return &EmployeeUseCase{
		trManager: trManager,
		repo:      repo,
	}
}

func (e *EmployeeUseCase) Ping(ctx context.Context) (*string, error) {
	var response *string
	err := e.trManager.Do(ctx, func(ctx context.Context) error {
		var err error
		response, err = e.repo.Ping(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
