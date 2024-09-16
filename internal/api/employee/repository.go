package employee

import "context"

type Repository interface {
	Ping(ctx context.Context) (*string, error)
}
