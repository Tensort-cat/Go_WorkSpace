package main

import (
	"errors"
	"fmt"
	"os"
	"sync/atomic"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

// 这个接收示例展示的是 RabbitMQ Stream 的“偏移量追踪”。
//
// 你可以把 offset 理解成：
// “某条消息在 Stream 里的序号 / 位置”。
//
// 消费者如果把自己已经处理到的 offset 保存下来，
// 那么下次重启后，就能从上次停止的位置继续往后读，
// 而不是每次都从头开始消费。
func main() {
	// 创建 Stream 环境，连接到 RabbitMQ 的 Stream 端口 5552。
	env, err := stream.NewEnvironment(
		stream.NewEnvironmentOptions().
			SetHost("localhost").
			SetPort(5552).
			SetUser("guest").
			SetPassword("guest"))
	CheckErrReceive(err)

	// 这里必须和发送端使用同一个 Stream 名称。
	streamName := "stream-offset-tracking-go"

	// 声明 Stream。
	// 如果已经存在，则继续使用；如果不存在，则创建。
	err = env.DeclareStream(streamName,
		&stream.StreamOptions{
			MaxLengthBytes: stream.ByteCapacity{}.GB(2),
		},
	)
	CheckErrReceive(err)

	// 记录“本次运行接收到的第一条消息 offset”。
	// 初始值 -1 表示目前还没收到消息。
	var firstOffset int64 = -1

	// 统计当前进程已经处理了多少条消息。
	// 初始值设为 -1，是为了配合下面 AddInt64(..., 1) 的写法，
	// 让第 1、11、21... 条消息触发一次 StoreOffset()。
	var messageCount int64 = -1

	// 记录读到 marker 那条消息时的 offset。
	// 因为回调和主协程之间会共享这个值，所以用 atomic.Int64。
	var lastOffset atomic.Int64

	// 用于通知主协程：“消费已完成，可以退出了”。
	ch := make(chan bool)

	// 每收到一条消息，客户端库都会调用这个回调。
	messagesHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
		// 只在第一次收到消息时记录 firstOffset。
		// CompareAndSwapInt64 的意思是：
		// 只有当 firstOffset 仍然是 -1 时，才更新成当前 offset。
		if atomic.CompareAndSwapInt64(&firstOffset, -1, consumerContext.Consumer.GetOffset()) {
			fmt.Println("First message received.")
		}

		// 每处理 10 条消息，手动保存一次 offset。
		//
		// 为什么要手动保存？
		// 因为这里用了 SetManualCommit()，
		// 所以“消费到哪里”完全由我们自己决定何时持久化。
		//
		// 保存后，RabbitMQ 会把这个消费者名对应的 offset 记住。
		// 程序下次启动时，可以通过 QueryOffset() 取回来。
		if atomic.AddInt64(&messageCount, 1)%10 == 0 {
			_ = consumerContext.Consumer.StoreOffset()
		}

		// 发送端最后会放一条 marker 消息，表示“这批测试数据到头了”。
		if string(message.GetData()) == "marker" {
			// 保存最后一条消息的位置，方便打印观察。
			lastOffset.Store(consumerContext.Consumer.GetOffset())

			// 再显式保存一次 offset，确保最终消费位置落盘。
			_ = consumerContext.Consumer.StoreOffset()

			// 关闭消费者，停止继续拉取消息。
			_ = consumerContext.Consumer.Close()

			// 通知主协程收尾退出。
			ch <- true
		}
	}

	// 这个变量决定“消费者从哪里开始读”。
	var offsetSpecification stream.OffsetSpecification

	// consumerName 是 offset tracking 的关键。
	// RabbitMQ 保存 offset 时，是按“消费者名 + Stream 名”来记的。
	// 所以只要下次继续使用同一个 consumerName，
	// 就能恢复到上次保存的位置。
	consumerName := "offset-tracking-tutorial"

	// 查询这个消费者在该 Stream 上是否有已保存的 offset。
	storedOffset, err := env.QueryOffset(consumerName, streamName)
	if errors.Is(err, stream.OffsetNotFoundError) {
		// 如果没查到，说明这是第一次运行，或者之前从未保存过 offset。
		// 那就从第一条消息开始消费。
		offsetSpecification = stream.OffsetSpecification{}.First()
	} else {
		// 如果查到了 storedOffset，说明上次已经处理到这个位置。
		// 所以本次应从下一条开始读，避免重复消费最后一条已处理消息。
		offsetSpecification = stream.OffsetSpecification{}.Offset(storedOffset + 1)
	}

	// 创建消费者。
	//
	// SetManualCommit():
	// 表示 offset 不自动提交，而是由我们在回调里调用 StoreOffset()。
	//
	// SetConsumerName(consumerName):
	// 指定用于保存 / 查询 offset 的消费者名字。
	//
	// SetOffset(offsetSpecification):
	// 指定本次消费的起始位置。
	_, err = env.NewConsumer(streamName, messagesHandler,
		stream.NewConsumerOptions().
			SetManualCommit().
			SetConsumerName(consumerName).
			SetOffset(offsetSpecification))
	CheckErrReceive(err)

	fmt.Println("Started consuming...")

	// 阻塞等待，直到消息回调中读到 marker 并发送完成信号。
	_ = <-ch
	fmt.Printf("Done consuming, first offset %d, last offset %d.\n", firstOffset, lastOffset.Load())
}

func CheckErrReceive(err error) {
	if err != nil {
		// 示例代码中统一采用“打印错误并退出”的处理方式。
		fmt.Printf("%s ", err)
		os.Exit(1)
	}
}
