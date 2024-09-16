package server

import (
	employeeDelivery "avito-internship/internal/api/employee/delivery"
	employeeRepo "avito-internship/internal/api/employee/repository"
	employeeUseCase "avito-internship/internal/api/employee/usecase"
	"avito-internship/internal/connection"

	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
)

func (s *Server) mapHandlers() error {
	db, err := connection.NewDB(s.config)
	if err != nil {
		return err
	}
	trManager := manager.Must(trmsqlx.NewDefaultFactory(db.Db))
	employeeRepository := employeeRepo.NewEmployeeRepository(trmsqlx.DefaultCtxGetter, db)
	employeeUsecase := employeeUseCase.NewEmployeeUseCase(trManager, employeeRepository)
	employeeHandlers := employeeDelivery.NewEmployeeHandler(employeeUsecase)
	MapRoutes(s.chi, employeeHandlers)
	return nil
}
