package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

const queueName = "Service1Queue"

func main() {
	// new connection
	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Open channel

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Queue declaration
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/send", func(c *gin.Context) {
		msg := c.Query("msg")
		if msg == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "message is required"})
		}

		// create message to publish
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(msg),
		}

		// publish a message to the queue
		err = ch.Publish("", queueName, false, false, message)
		if err != nil {
			log.Printf("failed to publish a message: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to publish a message"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": msg, "status": "success"})
		
	})

	log.Fatal(r.Run(":8080"))
}