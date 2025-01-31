gen:
	@protoc \
	--proto_path=./proto "./proto/analytics.proto" \
	--go_out=generated \
	--go_opt=paths=source_relative \
	--go-grpc_out=generated \
	--go-grpc_opt=paths=source_relative

run:
	@go run main.go