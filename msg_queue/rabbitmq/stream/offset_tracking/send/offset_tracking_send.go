package main

import (
	"fmt"
	"os"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

// 这个发送示例的作用，是先往 RabbitMQ Stream 里放一批测试消息，
// 供接收端示例演示“偏移量追踪（offset tracking）”。
//
// 可以先把它理解成：
// “先往一条有序日志里连续写入 100 条记录，后面的消费者再来按顺序读取。”
//
// 这个文件本身不负责保存 offset，但它准备的数据里有一条特殊消息 marker，
// 接收端正是靠它来判断“这一轮示例数据已经读完了”。
func main() {
	// 创建 Stream 环境。
	// 注意这里连接的是 RabbitMQ Stream 协议端口 5552，
	// 不是普通 AMQP 常见的 5672。
	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetHost("localhost").
			SetPort(5552).
			SetUser("guest").
			SetPassword("guest"))
	CheckErrSend(err)

	// 发送端和接收端必须使用同一个 streamName，
	// 否则接收端就不会消费到这里发送的数据。
	streamName := "stream-offset-tracking-go"

	// 如果 Stream 不存在，就创建它。
	// 这里把最大容量设置成 2GB，只是示例用途。
	err = env.DeclareStream(streamName,
		&stream.StreamOptions{
			MaxLengthBytes: stream.ByteCapacity{}.GB(2),
		},
	)
	CheckErrSend(err)

	// 创建生产者，后续所有消息都通过它写入 Stream。
	producer, err := env.NewProducer(streamName, stream.NewProducerOptions())
	CheckErrSend(err)

	// 本例固定发送 100 条消息。
	messageCount := 100

	// RabbitMQ Stream 的发送是异步确认的：
	// Send() 成功，只表示客户端库已经接受了这条消息；
	// 真正写入并被服务端确认，需要监听发布确认通知。
	chPublishConfirm := producer.NotifyPublishConfirmation()

	// 这个 channel 只是用于通知 main：
	// “所有消息都已经被确认，可以安全结束程序了”。
	ch := make(chan bool)
	handlePublishConfirm(chPublishConfirm, messageCount, ch)

	fmt.Printf("Publishing %d messages...\n", messageCount)
	for i := 0; i < messageCount; i++ {
		var body string

		// 最后一条消息使用 marker 作为标记。
		// 接收端读到它时，会保存最终 offset 并关闭消费者。
		if i == messageCount-1 {
			body = "marker"
		} else {
			body = "hello"
		}

		// 构造消息并发送。
		// 这里消息体非常简单，只有纯文本内容。
		err = producer.Send(amqp.NewMessage([]byte(body)))
		CheckErrSend(err)
	}

	// 等待所有消息都被 RabbitMQ 确认。
	// 如果这里不等待，程序可能在消息尚未真正写入时就提前退出。
	_ = <-ch
	fmt.Println("Messages confirmed.")

	// 关闭生产者，释放连接与内部资源。
	err = producer.Close()
	CheckErrSend(err)
}

func handlePublishConfirm(confirms stream.ChannelPublishConfirm, messageCount int, ch chan bool) {
	go func() {
		// 统计已经收到确认的消息数量。
		confirmedCount := 0

		// confirms 是一个 channel。
		// 每次收到的 confirmed 可能是一批确认结果，而不是单条。
		for confirmed := range confirms {
			for _, msg := range confirmed {
				// IsConfirmed() 表示该消息已经被服务端确认接收。
				if msg.IsConfirmed() {
					confirmedCount++

					// 当确认总数达到预期值时，通知主协程继续往下执行。
					if confirmedCount == messageCount {
						ch <- true
					}
				}
			}
		}
	}()
}

func CheckErrSend(err error) {
	if err != nil {
		// 示例程序采用最直接的方式处理错误：
		// 只要出错，就打印并退出，避免干扰学习主逻辑。
		fmt.Printf("%s ", err)
		os.Exit(1)
	}
}
