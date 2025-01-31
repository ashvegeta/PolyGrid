const { Kafka } = require("kafkajs");
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const path = require("path");

// Load the protobuf
const PROTO_PATH = path.join(__dirname, "./proto/analytics.proto");
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
const analyticsProto = grpc.loadPackageDefinition(packageDefinition).analytics;

// Create gRPC client
const client = new analyticsProto.AnalyticsService(
  "localhost:8080",
  grpc.credentials.createInsecure()
);

async function produce() {
  try {
    // Define Kafka client configuration
    const kafka = new Kafka({
      clientId: "polygrid-producer",
      brokers: ["localhost:9092"],
    });

    // Create a producer instance
    const producer = kafka.producer();

    // Connect to the Kafka broker
    await producer.connect();

    // Send messages to a topic
    let counter = 1;

    setInterval(async () => {
      const message = `Hello Kafka from Node.js! ${counter}`;

      // send produced message to kafka
      await producer.send({
        topic: "analytics",
        messages: [{ value: message }],
      });

      console.log(`Message sent to Kafka: ${message}`);

      // send message to logger - gRPC call
      const request = { message: message, senderType: "producer" };
      client.SendLog(request, (error, response) => {
        if (error) {
          console.error(`ERROR: failed to call gRPC server: ${error.message}`);
        } else {
          console.log(`Received gRPC ACK: ${response.message}\n`);
        }
      });

      // increment msg counter
      counter += 1;
    }, 1000);

    // Disconnect from the Kafka broker
    // await producer.disconnect();
  } catch (error) {
    console.error(`Error producing message: ${error.message}`);
  }
}

produce();
