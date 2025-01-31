import grpc
from generated import analytics_pb2_grpc, analytics_pb2
from dotenv import load_dotenv
import os

# Load environment variables from .env file
load_dotenv()

with grpc.insecure_channel(os.getenv("ANALYTICS_GRPC_ADDR")) as channel:
    stub = analytics_pb2_grpc.AnalyticsServiceStub(channel)
    message = "ashvegeta"
    res = stub.SendLog(analytics_pb2.SendLogRequest(message = message))
    assert res.message == f"message {message} received"
    