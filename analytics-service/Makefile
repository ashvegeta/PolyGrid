gen:
	@python -m grpc_tools.protoc -I./ \
	--python_out=./generated  \
	--grpc_python_out=./generated \
	analytics.proto 

run:
	@python analytics_server.py