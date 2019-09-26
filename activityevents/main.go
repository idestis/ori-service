//go:generate go get -u github.com/golang/protobuf/protoc-gen-go
//go:generate protoc activityevents.proto --go_out=plugins=grpc:./
//go:generate protoc health.proto --go_out=plugins=grpc:./health/v1

package activityevents
