from generated import analytics_pb2_grpc
import grpc
from concurrent import futures
from server import AnalyticsServicer

# function that exposes/serves analytics server's root gRPC functionalities
def serve():
    # init gRPC server
    srv = grpc.server(futures.ThreadPoolExecutor(max_workers = 4))
    
    # add/init gRPC server to - basically sending the class that holds the proto's function definitions
    analytics_pb2_grpc.add_AnalyticsServiceServicer_to_server(
        servicer = AnalyticsServicer(),
        server = srv
    )

    # init addr
    addr = "localhost:8080"
    srv.add_insecure_port(address = addr)

    # start service
    print(f"gRPC server is listening at {addr}....")
    srv.start()
    srv.wait_for_termination()

if __name__ == "__main__":
    serve()