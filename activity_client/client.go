package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"time"

	log "github.com/sirupsen/logrus"

	pb "github.com/idestis/ori-service/activityevents"
	"google.golang.org/grpc"
)

func createEvent(client pb.ActivityEventsClient, event *pb.EventRequest) {
	// Calling the streaming API on our server
	// Set the time when package was sent from client side
	resp, err := client.CreateEvent(context.Background(), event)
	if err != nil {
		log.Fatalf("Could not create Event: %v", err)
	}
	if resp.Success {
		log.Infof("An event created with id [%d] for deployment_id [%d] in %d ms", resp.Id, event.DeploymentId, resp.ExecutionTime)
	}
}

func listEvents(client pb.ActivityEventsClient, deployment *pb.Deployment) {
	stream, err := client.ListEvents(context.Background(), deployment)
	if err != nil {
		log.Fatalf("Error on get events: %v", err)
	}

	for {
		// Receiving the stream of data
		event, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("%v.ListEvents(_) = _, %v", client, err)
		}
		log.Infof("Event: %v", event)
	}

}

func main() {
	addPtr := flag.String("add", "", "Event to add")
	streamPtr := flag.Int("deployment", 0, "Stream deployment events by given ID")
	addressPtr := flag.String("address", "127.0.0.1", "Address of the server for connection")
	portPtr := flag.Int("port", 500051, "Port of the server for connection")
	flag.Parse()
	connectionString := fmt.Sprintf("%s:%d", *addressPtr, *portPtr)
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(connectionString, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewActivityEventsClient(conn)
	if *streamPtr != 0 {
		deployment := &pb.Deployment{Id: int32(*streamPtr)}
		listEvents(client, deployment)
	}
	if *addPtr != "" {
		event := &pb.EventRequest{}
		err := json.Unmarshal([]byte(*addPtr), event)
		event.Timestamp = &pb.Timestamp{Send: time.Now().UnixNano()}
		if err != nil {
			log.Errorf("Unable to parse %s as json data", *addPtr)
			return
		}

		createEvent(client, event)
	}
}
