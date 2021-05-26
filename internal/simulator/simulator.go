package simulator

import (
	"github.com/streadway/amqp"
	"log"
	"my.test.ru/internal/sensors"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Simulator struct {
	Status  string
	Sensors []sensors.Sensor
}

func (sim *Simulator) Simulate(ch *amqp.Channel, q amqp.Queue) {
	for i := range sim.Sensors {
		sim.Sensors[i].GenerateValue()
		err := ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(sim.Sensors[i].String()),
			})
		failOnError(err, "Failed to publish a message")
	}
}
