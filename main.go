package main

import (
	"github.com/streadway/amqp"
	"log"
	"my.test.ru/internal/sensors"
	"my.test.ru/internal/simulator"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	s1 := sensors.DefaultSensor{ID: 1, Type: "temperature", Name: "Датчик температуры 1", MinValue: 25, MaxValue: 30, Value: 25}
	s2 := sensors.DefaultSensor{ID: 2, Type: "humidity", Name: "Датчик влажности 1", MinValue: 60, MaxValue: 80, Value: 70}
	s3 := sensors.DefaultSensor{ID: 3, Type: "light", Name: "Датчик освещения 1", MinValue: 90, MaxValue: 100, Value: 95}
	sim := simulator.DefaultSimulator{Status: "idle", Sensors: []sensors.DefaultSensor{s1, s2, s3}}

	q, err := ch.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")
	log.Println("Simulation started")
	for {
		sim.Simulate(ch, q)
		log.Printf("%v sensors was simulated", len(sim.Sensors))
		time.Sleep(time.Second * 5)
	}
}
