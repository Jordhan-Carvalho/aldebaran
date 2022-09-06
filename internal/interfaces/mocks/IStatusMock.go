package interfaces

// Could mock the initializer aswell
type StatusServiceMock struct {}

func (serviceMock *StatusServiceMock) GetStatus() string {
   return "Aldebaran lives!!"
}
