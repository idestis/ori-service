############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
ENV GO111MODULE=on
RUN apk update && apk add --no-cache git build-base
WORKDIR app
COPY . .
# Build the binary.
RUN cd activity_server && go test && CGO_ENABLED=0 GOOS=linux go build -o /go/bin/activity_server
############################  
# STEP 2 Microservice images
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/activity_server /go/bin/activity_server
# Run the hello binary.
ENTRYPOINT ["/go/bin/activity_server"]