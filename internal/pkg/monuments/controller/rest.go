package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iconophilos/backend/internal/pkg/monuments/service"
	"go.uber.org/zap"
)

type REST struct {
	logger  *zap.Logger
	service service.Service
}

func NewRESTCtrl(logger *zap.Logger, svc service.Service) *REST {
	return &REST{
		logger:  logger,
		service: svc,
	}
}

func (ctrl *REST) Create(c *gin.Context) {
	var req CreateMonumentRequest
	if err := c.BindJSON(&req); err != nil {
		ctrl.logger.Warn("could not bind request", zap.Error(err))
		writeError(c, ErrInvalidPayload)
		return
	}

	mnmt, err := ctrl.service.Create(c.Request.Context(), req.toDomainModel())
	if err != nil {
		ctrl.logger.Warn("could not create monument", zap.Error(err))
		writeError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CreateMonumentResponse{
		Monument: newMonumentResponse(mnmt),
	})
}

func (ctrl *REST) Delete(c *gin.Context) {
	if err := ctrl.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		ctrl.logger.Warn("could not delete monument", zap.Error(err))
		writeError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (ctrl *REST) List(c *gin.Context) {
	if c.Query("name") != "" {
		mnmt, err := ctrl.service.FetchByName(c.Request.Context(), c.Query("name"))
		if err != nil {
			ctrl.logger.Warn("could not fetch monument", zap.Error(err))
			writeError(c, err)
			return
		}

		c.JSON(http.StatusOK, FetchMonumentByIDResponse{
			Monument: newMonumentResponse(mnmt),
		})
		return
	}

	mnmts, err := ctrl.service.List(c.Request.Context())
	if err != nil {
		ctrl.logger.Warn("could not list monuments", zap.Error(err))
		writeError(c, err)
		return
	}

	res := make([]*MonumentResponse, len(mnmts))
	for i, mnmt := range mnmts {
		res[i] = newMonumentResponse(mnmt)
	}

	c.JSON(http.StatusOK, ListMonumentsResponse{
		Monuments: res,
	})
}

func (ctrl *REST) FetchByID(c *gin.Context) {
	mnmt, err := ctrl.service.FetchByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		ctrl.logger.Warn("could not fetch monument", zap.Error(err))
		writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, FetchMonumentByIDResponse{
		Monument: newMonumentResponse(mnmt),
	})
}
