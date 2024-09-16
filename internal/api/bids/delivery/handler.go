package delivery

import (
	"net/http"

	bids2 "avito-internship/internal/api/bids"
)

var _ bids2.Handlers = (*BidsHandler)(nil)

type BidsHandler struct {
	useCase bids2.UseCase
}

func (b BidsHandler) Create() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (b BidsHandler) SubmitDecision() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (b BidsHandler) List() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (b BidsHandler) MyList() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (b BidsHandler) Status() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (b BidsHandler) Edit() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (b BidsHandler) Rollback() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (b BidsHandler) Reviews() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (b BidsHandler) Feedback() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}
