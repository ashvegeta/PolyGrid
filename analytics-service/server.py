from generated import analytics_pb2, analytics_pb2_grpc
import os
from datetime import datetime
import grpc

# Wrapper for gRPC server to inherit and define gRPC methods
class AnalyticsServicer(analytics_pb2_grpc.AnalyticsServiceServicer):
    def __init__(self):
        if not os.path.exists("logs"):
            os.makedirs("logs")
        if not os.path.exists("logs/producer"):
            os.makedirs("logs/producer")
        if not os.path.exists("logs/consumer"):
            os.makedirs("logs/consumer")

    def SendLog(self, request, context):
        # return analytics_pb2.SendLogResponse(message = f"message {request.message} received")

        # create repo for storing logs and sub-folder for producer/consumer (if it doesnt exist)
        # create a log file based on current date (make sure each request falls in this date)
        # append logs to the file 

        print("received request:", request.message)

        date, time = datetime.now().strftime("%d-%m-%y %H:%M:%S").split()
        log_file = os.path.join(f"logs/{request.senderType}", f"{date}.txt")
        
        try:
            with open(log_file, "a") as fp:
                fp.write(f"{time} - {request.message}\n")
            return analytics_pb2.SendLogResponse(message="Logged Successfully")
        except Exception as e:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return analytics_pb2.SendLogResponse()
        