# Simple Activity Events Microservice

[![Build Status](https://cloud.drone.io/api/badges/idestis/ori-service/status.svg)](https://cloud.drone.io/idestis/ori-service)
[![Test Coverage](https://api.codeclimate.com/v1/badges/3b7eaf418e30844ca273/test_coverage)](https://codeclimate.com/github/idestis/ori-service/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/3b7eaf418e30844ca273/maintainability)](https://codeclimate.com/github/idestis/ori-service/maintainability)

The service will recieve and store event messages for deployments by ID

**Known issue:** Server will save duplicated records as well :(

- **CI/CD**: Drone.io (tests, build & publish to DockerHub and Release on Github Releases)
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

Binaries `activity_server` and `activity_client` are available on the [Github Release Page](https://github.com/idestis/ori-service/releases)

### Kubernetes & Helm

```bash
$ helm install ./k8s/helm/activity_server --name activity-server --namespace activity-server --set image.tag=latest
```

Upgrade the version of server

```bash
$ helm upgrade activity-server ./k8s/helm/activity_server --set image.tag=0.0.3
```

## Develop

```bash
git clone https://github.com/idestis/ori-service.git ori-service && cd $_

go generate ./...     # only if you canged .proto files
```

You can test the service is running correctly by using `grpc-health-probe`

```bash
$ PORT=5000 go run activity_server/server.go &
$ go get github.com/grpc-ecosystem/grpc-health-probe
$ grpc-health-probe -addr=:5000
status: SERVING
```

Or from the docker container itself

```bash
$ docker run --rm -d --name=activity-server destis/activity-server:latest
$ docker exec -i activity-server grpc-health-probe -addr=:50051
status: SERVING
```

## 12 factor app practice

1. **Codebase**

Single codebase in trunk-based flow.

- master should be stable
- features are developed using pull requests (review practices, right now review is not possible, I'm solo developer)
- tags are identified releases

Note: Development using Trunk-based flow has been chosen based on ownership and trust. We keep master stable and trust in team review.

2. **Dependencies**

Go modules ease this process, everything is visible in go.sum, no hidden modules, no extra installations.

3. **Config**

Server have right now only one configuration which can be done using ``PORT`` environment variable, in future we can extend this based on requirements. Good practice to keep configuration in separate configuration service such as Consul.

4. **Backing services**

In the current state of this application is databaseless or eventless, so there are no backing services.

Ideally, we can use Consul for storing different configuration for example.

But if there is Configuration Service in the Stack, this logic should be in a separate library and designed on top of all microservices.

5. **Build, release, run**

As you can see, this three processes is strictly separated.

- build will create a docker image of the server
- release will be executed if there is any release planned using tags and will trigger next proccesses (eg build docker) and then release using helm. You can upgrade your current release using new image or rollout if something goes wrong.
- You can run this service in different methods (eg Docker, Kubernetes, Standalone) and use the client with it

6. **Processes**

This is not a stateless application, because I have a local storage for data, but ideally I need to keep it in database. 

Regarding resource limiting, I do not have limits or requests in this service. In real-case microservice, these resources will be tunned during circumstances and based on metrics.

7. Port biding

``PORT`` is available for setting the port.

You can change the port during deployment.

In client case you can choose the port where you want to connect

8. **Concurrency**

This service was built using gRPC. Each RPC handler attached to a registered server will be invoked in its own goroutine.

9. **Disposability**

Using Go we guaruntee fast startup. The reason there is no dependencies yet (eg DB or Consul), but even if we add this two dependencies, this will be on the same speed.

In Kubernetes side we can use ``terminationGracePeriodSeconds``, but we should catch interupt signal on the service side and using channels we can update the state of server and execute GracefulStop() of gRPC.

Unfortunately not done yet in this app. :(

10. **Dev/prod parity**

Using the Helm we can use the same codebase, but the deploy will be in the different namespaces, which will guarantee for us environments (eg dev, prod).

But the best option is to keep differents clusters.

11. **Logs**

The microservice print logs into stdout. As for me, is not good idea and bad implementation.

But actually, when I design application with microservices, I will suggest using centralized logging, tracing, service discovery and service mesh.

12. **Admin processes** 

Only one Admin process is available here, track the logs. There is not database oparations, only local storage, but we still don't have operation to wipe it or delete any record.


## Cloud-nativeness of this service

- Packaged as lightweight container
- Aligned to processes and lifecycle
- Automated capabilities using Kubernetes
- Defined, policy-driven resource allocation can be done using Kubernetes
- Centered around APIs for interaction and collaboration using gRPC

## Event Store

Event stores and event sourcing often also implies an event bus such as Cloud Pub/Sub (Google Cloud), RabbitMQ. 

We can use this in our stack to allow other services to be subscribed on topics.

For example we can use and update ``deployment-topic`` based on proccessed message of ``Activity_Server`` service when the event was added and than display in the activity page using websockets in ``Metrics`` service.
