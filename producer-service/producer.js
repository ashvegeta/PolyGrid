const { Kafka } = require("kafkajs");

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
      await producer.send({
        topic: "analytics",
        messages: [{ value: "Hello Kafka from Node.js! " + counter }],
      });
      counter += 1;
    }, 1000);

    // Disconnect from the Kafka broker
    // await producer.disconnect();
  } catch (error) {
    console.error(`Error producing message: ${error.message}`);
  }
}

produce();
