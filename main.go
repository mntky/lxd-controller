package main

import (
	"fmt"
	// json
	"encoding/json"
	"log"
	// RabbitMQ client library
	"github.com/streadway/amqp"

	// lxd client
	//"github.com/lxc/lxd/client"
	//"github.com/lxc/lxd/shared/api"
)

type ContainerInfo struct {
	Name	string	`json:"name"`
	Image	string	`json:"image"`
	VxlanId	int		`json:"vxlanid"`
}

type Message struct {
	Action	string	`json:"action"`
	Container ContainerInfo	`json:"container"`
}

func failOnError(err error, msg string){
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

	q, err := ch.QueueDeclare(
		"messages", // name
		false, // durable
		false, // delete when isused
		false, // exclusive
		false, //no-wait
		nil, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,	// queue
		"",	// consumer
		true, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func(){
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var messages []Message
			if err := json.Unmarshal([]byte(d.Body), &messages); err != nil {
				log.Fatal(err)
			}

			message := messages[0]
			fmt.Println(message.Action)
			switch message.Action {
				case "start": fmt.Println("Action: start")
				case "stop": fmt.Println("Action: stop")
				case "restart": fmt.Println("Action: restart")
				case "launch": fmt.Println("Action: launch")
				case "delete": fmt.Println("Action: delete")
				case "status": fmt.Println("Action: status")
			}
		}
	}()

	<-forever
}
