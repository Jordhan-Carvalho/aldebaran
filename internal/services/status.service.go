package services

type StatusService struct {}

func NewStatusService() StatusService {
  return StatusService{}
}

func (service *StatusService) GetStatus() (statusMessage string) {
  statusMessage = "Aldebaran lives!!"
  return
}
