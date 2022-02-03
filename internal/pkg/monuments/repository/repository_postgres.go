package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iconophilos/backend/pkg/db"
	"github.com/jackc/pgconn"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var _ Repository = (*PostgresRepository)(nil)

type PostgresRepository struct {
	logger *zap.Logger
	conn   *db.Conn
}

func NewPostgresRepository(logger *zap.Logger, conn *db.Conn) *PostgresRepository {
	return &PostgresRepository{
		logger: logger,
		conn:   conn,
	}
}

func (p *PostgresRepository) Create(ctx context.Context, m *Monument) error {
	tx := p.conn.WithContext(ctx).Create(&m)
	if tx.Error != nil {
		var pgErr *pgconn.PgError
		errors.As(tx.Error, &pgErr)
		if pq.ErrorCode(pgErr.Code) == uniqueViolationErr {
			return ErrDuplicateRecord
		}
		return fmt.Errorf("could not create monument: %w", tx.Error)
	}
	return nil
}

func (p *PostgresRepository) List(ctx context.Context) ([]*Monument, error) {
	var mnmts []*Monument
	if tx := p.conn.DB.WithContext(ctx).Where("deleted_at IS NULL").Find(&mnmts); tx.Error != nil {
		return nil, fmt.Errorf("could not get monument by id: %w", tx.Error)
	}
	return mnmts, nil
}

func (p *PostgresRepository) GetByID(ctx context.Context, id string) (*Monument, error) {
	var mnmt Monument
	if tx := p.conn.DB.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&mnmt); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("could not get monument by id: %w", tx.Error)
	}
	return &mnmt, nil
}

func (p *PostgresRepository) GetByName(ctx context.Context, name string) (*Monument, error) {
	var mnmt Monument
	if tx := p.conn.DB.WithContext(ctx).Where("name = ? AND deleted_at IS NULL", name).First(&mnmt); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, fmt.Errorf("could not get monument by id: %w", tx.Error)
	}
	return &mnmt, nil
}

func (p *PostgresRepository) Delete(ctx context.Context, id string, deletedAt time.Time) error {
	tx := p.conn.DB.WithContext(ctx).Exec("UPDATE monuments SET deleted_at = ? WHERE id = ?", deletedAt, id)
	if tx.Error != nil {
		return fmt.Errorf("could not remove monument: %w", tx.Error)
	}

	if tx.RowsAffected != 1 {
		return ErrRecordNotFound
	}
	return nil
}
