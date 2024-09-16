package delivery

import (
	"net/http"

	tender2 "avito-internship/internal/api/tender"
)

var _ tender2.Handlers = (*TenderHandler)(nil)

type TenderHandler struct {
	useCase tender2.UseCase
}

type Request struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ServiceType     string `json:"serviceType"`
	Status          string `json:"status"`
	OrganizationID  string `json:"organizationId"`
	CreatorUsername string `json:"creatorUsername"`
}

func (t TenderHandler) Create() http.HandlerFunc {
	//return func(writer http.ResponseWriter, request *http.Request) {
	//	log.Printf("[tender][delivery][CreateTender]")
	//
	//	var req Request
	//
	//	bodyStr, err := reqvalidator.ReadRequest(request, &req)
	//	if err != nil {
	//		http.Error(writer, fmt.Sprintf("Failed to read request body: %v", err), http.StatusBadRequest)
	//		return
	//	}
	//
	//	defer func(Body io.ReadCloser) {
	//		err = Body.Close()
	//		if err != nil {
	//			log.Printf("[tender][delivery][Create] Unable to close request body")
	//		}
	//	}(request.Body)
	//
	//	id, err := t.useCase.Create(request.Context(), req)
	//	if err != nil {
	//		if strings.Contains(err.Error(), errlst.ErrBoxAlreadyExists.Error()) {
	//			http.Error(writer, fmt.Sprintf("Failed to create: %v", err), http.StatusConflict)
	//			return
	//		}
	//
	//		http.Error(writer, fmt.Sprintf("Failed to create: %v", err), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	_, err = writer.Write([]byte(strconv.FormatInt(id, 10)))
	//	if err != nil {
	//		http.Error(writer, "Failed to write response", http.StatusInternalServerError)
	//		return
	//	}
	//}
	//TODO implement me
	panic("implement me")
}

func (t TenderHandler) List() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (t TenderHandler) MyList() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (t TenderHandler) Status() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (t TenderHandler) Edit() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (t TenderHandler) Rollback() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
