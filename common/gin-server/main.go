package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	r := gin.Default()

	// rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	exchangeName := "ex-logs"
	queueName := "qu-logs"

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	_ = ch.QueueBind(
		queueName,    // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)

	r.GET("/ping/:id", func(c *gin.Context) {
		mes := c.Param("id")

		ch.PublishWithContext(
			c,
			exchangeName,
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(mes),
			})
		fmt.Printf("sent %v success ! \n", mes)
		c.JSON(200, gin.H{
			"message": mes,
		})
	})
	r.GET("/getPing", func(c *gin.Context) {
		consume, _ := ch.Consume(
			queueName,
			"",    // consumer
			true,  // auto-ack
			false, // exclusive
			false, // no-local
			false, // no-wait
			nil,   // args
		)

		var wg sync.WaitGroup
		var res []string

		wg.Add(1)
		go func() {
			for i := range consume {
				fmt.Printf("Received a message: %s \n", i.Body)
				res = append(res, string(i.Body))

			}
			wg.Done()
		}()
		wg.Wait()
		c.JSON(200, gin.H{
			"message": res,
		})
	})
	r.Run(":8090")

}
