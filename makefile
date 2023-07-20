.PHONY: pb

pb:
	@protoc --go_out=./generated --go_opt=paths=source_relative \
    --go-grpc_out=./generated --go-grpc_opt=paths=source_relative \
    ./*.proto
run:
	@go run server/main.go
test:
	@go test ./config --cover
	@go test ./context --cover
	@go test ./services --cover
grpcui:
	@grpcui -plaintext localhost:8000