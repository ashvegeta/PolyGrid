setup-kafka:
	@docker run -d -p 9092:9092 --name broker apache/kafka:3.9.0
	@docker exec --workdir /opt/kafka/bin/ -it broker sh
	@./kafka-topics.sh --bootstrap-server localhost:9092 --create --topic analytics

run:
	@docker start broker
