from generated import analytics_pb2, analytics_pb2_grpc

# Wrapper for gRPC server to inherit and define gRPC methods
class AnalyticsServicer(analytics_pb2_grpc.AnalyticsServiceServicer):
    def __init__(self):
        pass

    def SendLog(self, request, context):
        return analytics_pb2.SendLogResponse(message = f"message {request.message} received")
    