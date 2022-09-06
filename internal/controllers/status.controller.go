package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/jordhan-carvalho/aldebaran/internal/interfaces"
)

type StatusController struct {
	StatusService interfaces.IStatusService
}

func NewStatusController(service interfaces.IStatusService) StatusController {
	return StatusController{StatusService: service}
}

func (controller *StatusController) GetAldebaranStatus(c *gin.Context) {
	status := controller.StatusService.GetStatus()

	c.JSON(http.StatusOK, status)
}
