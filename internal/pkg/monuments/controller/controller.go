package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iconophilos/backend/internal/pkg/monuments/service"
)

type Controller interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	List(c *gin.Context)
	FetchByID(c *gin.Context)
}

type CreateMonumentRequest struct {
	Name               string  `json:"name"`
	Type               string  `json:"type"`
	Dating             string  `json:"dating"`
	ArchitecturalPlant string  `json:"architectural_plant"`
	Model3D            string  `json:"3d_model"`
	Country            string  `json:"country"`
	Region             string  `json:"region"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
}

func (c *CreateMonumentRequest) toDomainModel() *service.Monument {
	return &service.Monument{
		Name:               c.Name,
		Type:               c.Type,
		Dating:             c.Dating,
		ArchitecturalPlant: c.ArchitecturalPlant,
		Model3D:            c.Model3D,
		Country:            c.Country,
		Region:             c.Region,
		Latitude:           c.Latitude,
		Longitude:          c.Longitude,
	}
}

type CreateMonumentResponse struct {
	Monument *MonumentResponse `json:"data"`
}

type MonumentResponse struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Dating             string    `json:"dating"`
	Type               string    `json:"type"`
	ArchitecturalPlant string    `json:"architectural_plant"`
	Model3D            string    `json:"3d_model"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Country            string    `json:"country"`
	Region             string    `json:"region"`
	Latitude           float64   `json:"latitude"`
	Longitude          float64   `json:"longitude"`
}

func newMonumentResponse(m *service.Monument) *MonumentResponse {
	return &MonumentResponse{
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
	}
}

type ListMonumentsResponse struct {
	Monuments []*MonumentResponse `json:"data"`
}

type FetchMonumentByIDResponse CreateMonumentResponse
