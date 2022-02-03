package service

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, m *Monument) (*Monument, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Monument, error)
	FetchByID(ctx context.Context, id string) (*Monument, error)
	FetchByName(ctx context.Context, name string) (*Monument, error)
}
