load-env-unix:
	@echo 1. Loading environment variables from .env file (Unix)
	@set -a && . .env && set +a

load-env-windows:
	@echo 1. Loading environment variables from .env file (Windows)
	@for /f "tokens=1,2 delims==" %%i in (.env) do set %%i=%%j

ifeq ($(OS),Windows_NT)
load-env: load-env-windows
else
load-env: load-env-unix
endif

setup-kafka: load-env
	@echo 2. setting up kafka and creating topics
	@docker run -d -p 9092:9092 --name broker apache/kafka:3.9.0
	@docker exec --workdir /opt/kafka/bin/ -it broker sh -c "./kafka-topics.sh --bootstrap-server localhost:9092 --create --topic analytics"

run: load-env
	@echo 2. starting kafka broker
	@docker start broker