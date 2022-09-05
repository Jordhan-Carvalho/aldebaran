package services

import "testing"

func TestGetStatus(t *testing.T) {
  expectedResp := "Aldebaran lives!!"

  statusService := StatusService{} 
  resp := statusService.GetStatus()

  if resp != expectedResp {
    t.Errorf("Expected %v but got %v\n", expectedResp, resp)
  }
}

