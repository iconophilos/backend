package service

import (
	"time"

	"github.com/iconophilos/backend/internal/pkg/monuments/repository"
)

type Monument struct {
	ID                 string
	Name               string
	Type               string
	Dating             string
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

func newStoreModel(m *Monument) *repository.Monument {
	return &repository.Monument{
		ID:                 m.ID,
		Name:               m.Name,
		Dating:             m.Dating,
		Type:               m.Type,
		ArchitecturalPlant: m.ArchitecturalPlant,
		Model3D:            m.Model3D,
		Country:            m.Country,
		Region:             m.Region,
		Latitude:           m.Latitude,
		Longitude:          m.Longitude,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
		DeletedAt:          m.DeletedAt,
	}
}

func newMonumentModel(mnmtStore *repository.Monument) *Monument {
	return &Monument{
		ID:                 mnmtStore.ID,
		Name:               mnmtStore.Name,
		Dating:             mnmtStore.Dating,
		Type:               mnmtStore.Type,
		ArchitecturalPlant: mnmtStore.ArchitecturalPlant,
		Model3D:            mnmtStore.Model3D,
		Country:            mnmtStore.Country,
		Region:             mnmtStore.Region,
		Latitude:           mnmtStore.Latitude,
		Longitude:          mnmtStore.Longitude,
		CreatedAt:          mnmtStore.CreatedAt,
		UpdatedAt:          mnmtStore.UpdatedAt,
	}
}
