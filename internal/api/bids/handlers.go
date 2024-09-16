package bids

import "net/http"

type Handlers interface {
	Create() http.HandlerFunc
	SubmitDecision() http.HandlerFunc
	List() http.HandlerFunc
	MyList() http.HandlerFunc
	Status() http.HandlerFunc
	Edit() http.HandlerFunc
	Rollback() http.HandlerFunc
	Reviews() http.HandlerFunc
	Feedback() http.HandlerFunc
}
