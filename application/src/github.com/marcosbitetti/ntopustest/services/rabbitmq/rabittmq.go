package rabbitmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Broadcaster chan string

func run() {
	conn, err := amqp.Dial("amqp://" + os.Getenv("RABBITMQ_DEFAULT_USER") + ":" + os.Getenv("RABBITMQ_DEFAULT_PASS") + "@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ Channel")
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal("Failed to connect to declare queue")
	}

	for msg := range Broadcaster {
		log.Println(msg)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			})
	}

	conn.Close()
}

func Run() {
	Broadcaster = make(chan string)
	go run()
}
