package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/idestis/ori-service/activityevents"
)

func TestCreateEvent(t *testing.T) {
	//  createEvent basic test
	s := server{}

	req := &pb.EventRequest{
		Id:           1,
		DeploymentId: 1,
		Message:      "Testing createEvent endpoint",
		Timestamp: &pb.Timestamp{
			Send: time.Now().UnixNano(),
		}}
	res, err := s.CreateEvent(context.Background(), req)

	if err != nil {
		t.Error("CreateEvent() got unexpected error")
	}
	// We expect to get true here, because our service is able to store any event
	if !res.Success {
		t.Errorf("Expected true but got %v createEvent function", res.Success)
	}
}
