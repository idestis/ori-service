# Simple Activity Events service

[![Build Status](https://cloud.drone.io/api/badges/idestis/ori-service/status.svg)](https://cloud.drone.io/idestis/ori-service)

The service will recieve and store event messages for deployments by ID

**Known issue:** Server will save duplicated records as well :(

- **CI/CD**: Drone.io (tests)
- **Coverage**: TODO
- **Code Quality**: TODO
- **Orchestration**: Kubernetes
- **Config Management**: Helm
- **Local Development**: protoc, grpcurl

This microservice was created from scratch for some reason. We all have a reason, right?
There is a lot of methods to generate microservices, for example there is a tool called [Lile](https://github.com/lileio/lile) which will create microservice skeleton with best practices like metrics (eg Prometheus), Service Discovery and Tracing

## Usage

Refer to [Install](#Install) for getting `activity_server` and `activity_client`

First, run `activity_server` somewhere (eg you can run it on your localhost)

```bash
./activity_server
```

Then, we can query our server using the client. There is two possible actions here:

- create event

```bash
./activity_client -add '{"id":1, "deployment_id": 1, "message":"Deployment has beed created successfully"}'
INFO[0000] An event created with id [1] for deployment_id [1] in 4 ms
```

- fetch events

```bash
./activity_client -deployment 1
INFO[0000] Event: id:1 message:"Deployment has beed created successfully" timestamp:<send:1569500037894488000 receive:1569500037898843000 > deployment_id:1
```

Default output if no params was given will show you help information

```bash
-add string
    Event to add
  -address string
    Address of the server for connection (default "127.0.0.1")
  -deployment int
    Stream deployment events by given ID
  -port int
    Port of the server for connection (default 50051)
```

## Install

### Docker image

For easy usage I provided Docker Image at the Docker Hub

```bash
docker run -e PORT=50051 -p 50051:50051 destis/activity_server:latest
INFO[0000] Server has been started on 50051
```

### Binaries

Binaries `activity_server` and `activity_client` are available on the Github Release Page

### Kubernetes & Helm

```bash
helm install ./k8s/helm/activity_server --name activity_server --namespace activity_server --set image.tag=latest
```

## Develop

```bash
git clone https://github.com/idestis/ori-service.git ori-service && cd $_

go generate ./...     # only if you canged .proto files
```
