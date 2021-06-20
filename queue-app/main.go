package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

var upgrader = websocket.Upgrader{} // use default options

var broadcaster chan string

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func echo(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Scoket error", err)
	}

	defer c.Close()
	for msg := range broadcaster {
		//err = c.WriteMessage(websocket.TextMessage, []byte(msg))
		var bmsg []byte = []byte(msg)
		bmsg = append(bmsg, 0)
		err = c.WriteMessage(websocket.TextMessage, bmsg)
	}
}

func main() {
	// check if environment is loaded. if not, assume that is on developer
	if os.Getenv("APP_MONITOR_PORT") == "" {
		err := godotenv.Load("./../.env")
		if err != nil {
			panic(err)
		}
	}

	// start websocket
	broadcaster = make(chan string)
	var addr string = ":" + string(os.Getenv("APP_MONITOR_PORT"))
	http.HandleFunc("/", index)
	http.HandleFunc("/echo", echo)
	go http.ListenAndServe(addr, nil)

	// start rabbitmq
	user := os.Getenv("RABBITMQ_DEFAULT_USER")
	pass := os.Getenv("RABBITMQ_DEFAULT_PASS")

	conn, err := amqp.Dial("amqp://" + user + ":" + pass + "@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("RABITMQ: %s", d.Body)
			broadcaster <- string(d.Body)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
