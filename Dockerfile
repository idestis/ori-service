############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
WORKDIR app
COPY . .
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git coreutils \
  && GO111MODULE=off go get github.com/grpc-ecosystem/grpc-health-probe \
  && GO111MODULE=on go get -v ./...
# Build the binary.
RUN GO111MODULE=on go install ./...
############################  
# STEP 2 Microservice images
############################
FROM alpine
WORKDIR /bin/
RUN apk add --no-cache bash ca-certificates
# Copy our static executable.
COPY --from=builder /go/bin/grpc-health-probe /go/bin/activity_server /go/bin/activity_client ./
ENV PORT 50051
EXPOSE 50051
# Run the activity_server binary.
ENTRYPOINT ["activity_server"]