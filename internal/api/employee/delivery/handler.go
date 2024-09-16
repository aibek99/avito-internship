package delivery

import (
	"fmt"
	"net/http"

	employee2 "avito-internship/internal/api/employee"
)

var _ employee2.Handlers = (*EmployeeHandler)(nil)

type EmployeeHandler struct {
	useCase employee2.UseCase
}

func NewEmployeeHandler(useCase employee2.UseCase) *EmployeeHandler {
	return &EmployeeHandler{useCase: useCase}
}

func (e EmployeeHandler) Ping() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		response, err := e.useCase.Ping(request.Context())
		if err != nil {
			http.Error(writer, fmt.Sprintf("Failed to get: %v", err), http.StatusInternalServerError)
			return
		}
		_, err = writer.Write([]byte(*response))
		if err != nil {
			http.Error(writer, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

func (e EmployeeHandler) Tenders() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		response, err := e.useCase.Ping(request.Context())
		if err != nil {
			http.Error(writer, fmt.Sprintf("Failed to get: %v", err), http.StatusInternalServerError)
			return
		}
		_, err = writer.Write([]byte(*response))
		if err != nil {
			http.Error(writer, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}
