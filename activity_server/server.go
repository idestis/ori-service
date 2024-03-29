package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"

	pb "github.com/idestis/ori-service/activityevents"
	grpc_health_v1 "github.com/idestis/ori-service/activityevents/health/v1"
)

type server struct {
	// No database is presented in this service
	// In case if we will store this events to database, this part will be updated
	savedEvents []*pb.EventRequest
}

type HealthImpl struct {
	// This is our healthcheck implementation
}

func (h *HealthImpl) Check(ctx context.Context, args *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	// Health-check endpoint
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthImpl) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	// Keep-alive streaming
	return nil
}

func (s *server) CreateEvent(ctx context.Context, in *pb.EventRequest) (*pb.EventResponse, error) {
	// Set the time when package was received on the server side
	// in.Timestamp = &pb.Timestamp{Receive: time.Now().Unix()}
	in.Timestamp.Receive = time.Now().UnixNano()
	timeDiff := (in.Timestamp.Receive - in.Timestamp.Send) / 1000000
	s.savedEvents = append(s.savedEvents, in)
	log.Infof("Created event [%d] for deployment [%d] in %d ms", in.Id, in.DeploymentId, timeDiff)
	return &pb.EventResponse{Id: in.Id, Success: true, ExecutionTime: timeDiff}, nil
}

func (s *server) ListEvents(deployment *pb.Deployment, stream pb.ActivityEvents_ListEventsServer) error {
	for _, event := range s.savedEvents {
		if deployment.Id != 0 {
			if event.DeploymentId != deployment.Id {
				continue
			}
		}
		if err := stream.Send(event); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Infof("Server has been started on %s", port)

	// Creates a new gRPC server
	s := grpc.NewServer()
	defer s.Stop()
	pb.RegisterActivityEventsServer(s, &server{})
	grpc_health_v1.RegisterHealthServer(s, &HealthImpl{})
	s.Serve(lis)
}
