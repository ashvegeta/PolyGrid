import grpc
from generated import analytics_pb2_grpc, analytics_pb2

with grpc.insecure_channel("localhost:8080") as channel:
    stub = analytics_pb2_grpc.AnalyticsServiceStub(channel)
    message = "ashvegeta"
    res = stub.SendLog(analytics_pb2.SendLogRequest(message = message))
    assert res.message == f"message {message} received"
    