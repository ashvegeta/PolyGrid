import unittest
from unittest.mock import patch, MagicMock
from generated import user_pb2
from main import call_say_hello

class TestUserComm(unittest.TestCase):
    @patch('grpc.insecure_channel')
    @patch('user_pb2_grpc.UserStub')
    def test_call_say_hello(self, MockUserStub, MockInsecureChannel):
        # Create a mock response
        mock_response = MagicMock()
        mock_response.message = 'Hello, ash!'
        
        # Set up the mock stub
        mock_stub = MockUserStub.return_value
        mock_stub.SayHello.return_value = mock_response
        
        # Set up the mock channel
        mock_channel = MagicMock()
        MockInsecureChannel.return_value = mock_channel
        
        # Call the function
        response_message = call_say_hello('ash')
        
        # Assert the response
        self.assertEqual(response_message, 'Hello, ash!')
        mock_stub.SayHello.assert_called_once_with(user_pb2.HelloRequest(name='ash'))
        MockInsecureChannel.assert_called_once_with('localhost:8080')

if __name__ == '__main__':
    unittest.main()