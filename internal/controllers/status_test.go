package controllers

import (
	"fmt"
	"testing"

	"github.com/jordhan-carvalho/aldebaran/internal/interfaces/mocks"
)

func TestGetAldebaranStatus(t *testing.T) {
  expectedResp := "Aldebaran lives!!"
	serviceMock := interfaces.StatusServiceMock{}

	statusController := NewStatusController(&serviceMock)


  fmt.Printf("Just %v %v", expectedResp, statusController)
  // it should test if its context.JSON is returning OK and the message content
	/* resp := statusController.GetAldebaranStatus()

  if resp != expectedResp {
    t.Errorf("Expected %v but got %v", expectedResp, resp)
  } */


}
