start:
	cd cmd/ordersystem && go run main.go wire_gen.go

generate_grpc:
	protoc --go_out=. --go-grpc_out=. ./internal/infra/grpc/protofiles/order.proto

generate_wire:
	cd cmd/ordersystem && wire gen

generate_graphql:
	go run github.com/99designs/gqlgen generate

generate:
	go generate ./...

start_evans:
	evans -r repl
