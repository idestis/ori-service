# Activity Events

The service will recieve and store event messages for deployments by ID

**Known issue:** Server will save duplicated records as well :(

## Server

The server was built on top of gRPC

In gRPC we have two methods

This method allows you to create event with message for any deployment.
In the response you will get "execution_ time" as calculated value from both sides (client-and-server)
```
rpc CreateEvent(EventRequest) returns (EventResponse) {}
```

ListEvents methods  allow you to retreive stream of events for requested deployment ID.

```
rpc ListEvents(Deployment) returns (stream EventRequest) {}
```

## Client

Add new event method call from client

```
go run activity_client/client.go -add '{"id":1, "deployment_id": 1, "message":"Deployment has beed created successfully"}'
```

Retreive stream of the events by deployment ID

```
go run activity_client/client.go -deployment 1
```

### Build

In the folder with server you can find already compiled binary for Darwins, but in any case you can build



## Rebuilding the generated code

```
protoc -I activityevents activityevents/activityevents.proto --go_out=plugins=grpc:activityevents
```