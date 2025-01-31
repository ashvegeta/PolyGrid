PROTO_DIR=./proto
OUT_DIR=./generated

gen:
	@protoc \
	--proto_path=$(PROTO_DIR) $(PROTO_DIR)/analytics.proto \
	--js_out=import_style=commonjs,binary:$(OUT_DIR) 

run:
	@npm run start