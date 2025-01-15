.PHONY: test

gen:
	@protoc \
	--proto_path=. "./user.proto" \
	--go_out=generated \
	--go_opt=paths=source_relative \
    --go-grpc_out=generated \
	--go-grpc_opt=paths=source_relative
		
build:
	@go build -o cmd/server.exe .

run:
	@./cmd/server.exe

test:
	@go test -v ./test/...