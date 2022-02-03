package repository

import (
	"context"
	"time"
)

type Repository interface {
	Create(ctx context.Context, m *Monument) error
	List(ctx context.Context) ([]*Monument, error)
	GetByID(ctx context.Context, id string) (*Monument, error)
	GetByName(ctx context.Context, name string) (*Monument, error)
	Delete(ctx context.Context, id string, deletedAt time.Time) error
}

type Monument struct {
	ID                 string `gorm:"primaryKey;type:uuid"`
	Name               string `gorm:"uniqueIndex"`
	Dating             string
	Type               string
	ArchitecturalPlant string
	Model3D            string
	Country            string
	Region             string
	Latitude           float64
	Longitude          float64
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}
