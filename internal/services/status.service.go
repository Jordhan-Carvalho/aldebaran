package services

type StatusService struct {}

func (service *StatusService) GetStatus() (statusMessage string) {
  statusMessage = "Aldebaran lives!!"
  return
}
