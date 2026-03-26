package main

import (
	"bytes"
	"log"
	"rabbitmq_base/util"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin123@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 发送端和接收端必须对这个队列的声明达成一致
	q, err := ch.QueueDeclare(
		"work_queue", // name
		true,         // durability
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)
	util.FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count (这个消费者最多同时持有 1 条未确认消息)
		0,     // prefetch size
		false, // global (false 表示这个限制是按当前消费者/当前消费上下文来起作用)
	)
	util.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			/*
				Ack() 的参数含义：
				false：只确认当前这一条消息
				true：一次性确认当前这条以及它之前尚未确认的多条消息
			*/
			err := d.Ack(false)
			util.FailOnError(err, "Failed to ack message")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
