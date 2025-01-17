from flask import Flask, jsonify, request
import grpc
import generated.user_pb2_grpc as user_pb2_grpc
import generated.user_pb2 as user_pb2
 
app = Flask(__name__)

@app.route('/recommend', methods=['POST'])
def recommend():
    data = request.json
    user_name = data.get('name', 'default_user')
    recommendations = {"movies" : ["movie1", "movie2"]}  # Placeholder for recommendations
    return jsonify(recommendations)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)