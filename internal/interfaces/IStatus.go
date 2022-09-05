package interfaces

type IStatusService interface {
  GetStatus() string  
}

type IStatusController interface {
  GetAldebaranStatus(service IStatusService) string
}
