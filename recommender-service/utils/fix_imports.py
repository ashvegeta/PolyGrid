file_path = './generated/user_pb2_grpc.py'

with open(file_path, 'r') as file:
    content = file.read()

content = content.replace('import user_pb2 as user__pb2', 'import generated.user_pb2 as user__pb2')

with open(file_path, 'w') as file:
    file.write(content)