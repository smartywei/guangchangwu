package queue

import (
	"github.com/streadway/amqp"
	"log"
	"fmt"
	"rabbitmq/tool"
)

var Conn *amqp.Connection

var Ch *amqp.Channel

var Que amqp.Queue

func getConn() *amqp.Connection {

	conn, err := amqp.Dial("amqp://admin:admin@127.0.0.1:5672/")

	Conn = conn

	FailOnError(err, "Failed to connect to RabbitMQ")

	return Conn
}

func getCh() *amqp.Channel {

	if Conn == nil{
		Conn = getConn()
	}

	ch, err := Conn.Channel()

	Ch = ch

	FailOnError(err, "Failed to open a channel")

	return Ch
}

func CloseConnAndCh() {
	Conn.Close()
	Ch.Close()
}

func getOrCreateQueue(name string) amqp.Queue {

	if Ch == nil{
		Ch = getCh()
	}

	que, err := Ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	Que = que

	FailOnError(err, "Failed to declare a queue")

	return Que
}

func PushToQueue(queueName string, body []byte) {

	if Ch == nil{
		Ch = getCh()
	}

	err := Ch.Publish(
		"",
		getOrCreateQueue(queueName).Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	FailOnError(err, "Fialed to publish a message")
}

func PullToQueue(queueName string) <-chan amqp.Delivery{
	msgs, err := Ch.Consume(
		getOrCreateQueue(queueName).Name,
		"",
		true,
		false,
		false,
		true,
		nil,
	)
	tool.FailOnError(err,"Failed to register a consumer")
	return msgs
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
