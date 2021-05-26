package simulator

import (
	"github.com/streadway/amqp"
	"my.test.ru/internal/sensors"
	"testing"
)

func TestSimulator_Simulate(t *testing.T) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	testChannel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer testChannel.Close()
	testQueue, err := testChannel.QueueDeclare(
		"testSimulate", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")
	tS1 := sensors.Sensor{ID: 1, Type: "test", Name: "Тестовый датчик 1", MinValue: 25, MaxValue: 30, Value: 25}
	tS2 := sensors.Sensor{ID: 2, Type: "test", Name: "Тестовый датчик 2", MinValue: 60, MaxValue: 80, Value: 70}
	tS3 := sensors.Sensor{ID: 3, Type: "test", Name: "Тестовый датчик 3", MinValue: 90, MaxValue: 100, Value: 95}
	testSim := Simulator{Status: "idle", Sensors: []sensors.Sensor{tS1, tS2, tS3}}
	testSim.Simulate(testChannel, testQueue)
	msgs, err := testChannel.Consume(
		testQueue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	allMsgs := make([]amqp.Delivery, 0, 3)

	for d := range msgs {
		allMsgs = append(allMsgs, d)
		if len(allMsgs) == len(testSim.Sensors) {
			break
		}
	}
	for i, d := range allMsgs {
		if string(d.Body) != testSim.Sensors[i].String() {
			t.Errorf("TestSimulator_Simulate() recieved wrong data  = %s; want %s", string(d.Body), testSim.Sensors[i].String())
		}
	}
	_, err = testChannel.QueueDelete(testQueue.Name, false, false, false)
	failOnError(err, "Failed to delete a queue")
}
