package employee

import "net/http"

type Handlers interface {
	Ping() http.HandlerFunc
}
