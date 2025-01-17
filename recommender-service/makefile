.PHONY: test

gen:
	@python -m grpc_tools.protoc -I../shared/proto \
	--python_out=./generated  \
	--grpc_python_out=./generated \
	user.proto 	
	@python utils/fix_imports.py 

run:
	@python src/main.py

test:
	@python -m unittest discover -s test