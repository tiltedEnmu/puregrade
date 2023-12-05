run:
	go run ./cmd/main.go

doc:
	godoc -http=localhost:6060 # http://localhost:6060/pkg/github.com/puregrade/puregrade-auth/?m=all

generate-proto:
	protoc --go_out=./internal/transport/grpc/proto --go-grpc_out=./internal/transport/grpc/proto ./internal/transport/grpc/proto/auth.proto