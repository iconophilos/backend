package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iconophilos/backend/internal/pkg/monuments/service"
)

var (
	// Enumerate transport errors

	ErrInvalidPayload error = errors.New("invalid payload")
)

func writeError(c *gin.Context, err error) {
	e := unwrapErr(err)

	switch e {
	case service.ErrMonumentNotFound:
		c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
	case
		service.ErrMonumentAlreadyExists,
		service.ErrMissingName,
		service.ErrMissingType,
		service.ErrInvalidType,
		service.ErrMissingDating,
		service.ErrMissingArchitecturalPlant,
		service.ErrMissingModel3D,
		service.ErrMissingLocation,
		service.ErrMissingCountry,
		service.ErrMissingRegion,
		service.ErrMissingLatitude,
		service.ErrMissingLongitude,
		service.ErrMonumentAlreadyExists,
		ErrInvalidPayload:
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
	}
}

func unwrapErr(err error) error {
	if e := errors.Unwrap(err); e != nil {
		return e
	}
	return err
}
