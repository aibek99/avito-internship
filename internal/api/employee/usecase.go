package employee

import "context"

type UseCase interface {
	Ping(ctx context.Context) (*string, error)
}
