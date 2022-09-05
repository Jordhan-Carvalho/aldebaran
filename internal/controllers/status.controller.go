package controllers


import (

	"github.com/jordhan-carvalho/aldebaran/internal/interfaces"
)


type StatusController struct {} 

func (controller *StatusController) GetAldebaranStatus(service interfaces.IStatusService) string {
  status := service.GetStatus()
  return status
}

