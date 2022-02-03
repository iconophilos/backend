package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iconophilos/backend/internal/pkg/monuments/repository"
	"go.uber.org/zap"
)

var _ Service = (*DefaultService)(nil)

type DefaultService struct {
	logger *zap.Logger
	repo   repository.Repository
}

func NewDefaultService(logger *zap.Logger, repo repository.Repository) *DefaultService {
	return &DefaultService{
		logger: logger,
		repo:   repo,
	}
}

func (s *DefaultService) Create(ctx context.Context, m *Monument) (*Monument, error) {
	if err := validateCreateInput(m); err != nil {
		return nil, fmt.Errorf("could not validate input for monument creation: %w", err)
	}

	m.ID = uuid.New().String()
	now := time.Now().UTC()
	m.CreatedAt, m.UpdatedAt = now, now

	// Store new monument
	store := newStoreModel(m)
	if err := s.repo.Create(ctx, store); err != nil {
		if err == repository.ErrDuplicateRecord {
			return nil, ErrMonumentAlreadyExists
		}
		return nil, fmt.Errorf("could not create monument: %w", err)
	}

	return m, nil
}

func (s *DefaultService) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id, time.Now().UTC()); err != nil {
		if err == repository.ErrRecordNotFound {
			return ErrMonumentNotFound
		}
		return fmt.Errorf("could not delete monument: %w", err)
	}
	return nil
}

func (s *DefaultService) List(ctx context.Context) ([]*Monument, error) {
	mnmtsStore, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get monuments: %w", err)
	}

	mnmts := make([]*Monument, len(mnmtsStore))
	for i, m := range mnmtsStore {
		mnmts[i] = newMonumentModel(m)
	}
	return mnmts, nil
}

func (s *DefaultService) FetchByID(ctx context.Context, id string) (*Monument, error) {
	m, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return nil, ErrMonumentNotFound
		}
		return nil, fmt.Errorf("could not get monument: %w", err)
	}
	return newMonumentModel(m), nil
}

func (s *DefaultService) FetchByName(ctx context.Context, name string) (*Monument, error) {
	fmt.Println("NAME: ", name)

	m, err := s.repo.GetByName(ctx, name)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return nil, ErrMonumentNotFound
		}
		return nil, fmt.Errorf("could not get monument: %w", err)
	}
	return newMonumentModel(m), nil
}
