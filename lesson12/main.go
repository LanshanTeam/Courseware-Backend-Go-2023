package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	reader *kafka.Reader
	topic  = "test"
)

func main() {
	CreateTopic()
	ctx := context.Background()
	writeKafka(ctx)
	// go listenSignal()
	readKafka(ctx)
}

// 创建 topic
func CreateTopic() {
	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  topic,
		AllowAutoTopicCreation: true, // 自动创建topic
	}

	messages := []kafka.Message{
		{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	}

	var err error
	const retries = 3
	// 重试3次
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = w.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Fatalf(" create Topic error %v", err)
		}
		break
	}

	// 关闭Writer
	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

// 生产消息
func writeKafka(ctx context.Context) {
	writer := kafka.Writer{
		Addr:         kafka.TCP("127.0.0.1:9092"),
		Topic:        topic,
		Balancer:     &kafka.Hash{}, //负载均衡算法
		WriteTimeout: 1 * time.Second,
		//kafka操作不应该影响正常服务的调用，所以设置响应限时
		RequiredAcks: kafka.RequireNone,
		// 最简单，但是最不安全的 不需要 acks

		//AllowAutoTopicCreation: true,
		// topic 没有创建的话可以设置为true让自动创建
		// 但是工作当中，你只是使用者，无权创建，此处只是演示
	}
	defer writer.Close()
	// 函数允许传入不定长的消息，原子性操作
	if err := writer.WriteMessages(ctx,
		kafka.Message{Key: []byte("1"), Value: []byte("h")},
		kafka.Message{Key: []byte("2"), Value: []byte("e")},
		kafka.Message{Key: []byte("3"), Value: []byte("l")},
		kafka.Message{Key: []byte("4"), Value: []byte("l")},
		kafka.Message{Key: []byte("5"), Value: []byte("o")},
	); err != nil {
		// 一开始 topic 没有创建，写入肯定失败， 意料之中 ，让循环三次尝试
		fmt.Printf("批量写入kafka失败:%v\n", err)
	} else {
		fmt.Printf("success")
	}
}

func readKafka(ctx context.Context) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		// 一个 broker 就是一个服务器上运行的kafka实例
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		// 每一秒上报一次当前读取的位置
		CommitInterval: 1 * time.Second,
		// 每个group只能消费一份确定的topic的数据，消费者需要用group id说明是哪个group在消费
		GroupID: "server_team",
		// 历史第一条消息开始消费 还是  该消费者上线后的最新消息开始
		StartOffset: kafka.FirstOffset, // 历史第一条
	})
	//  defer reader.Close() 这一行根本执行不到，与writer不同

	// 死循环一直读取消息
	for {
		if message, err := reader.ReadMessage(ctx); err != nil {
			fmt.Printf("读取kafka失败:%v\n", err)
			break
		} else {
			fmt.Printf("topic=%s,partition=%s,offset=%d,key=%d,value=%d", message.Topic, message.Partition, message.Offset, message.Key, message.Value)
		}
	}
}

// // 当系统杀进程的时候，就需要把reader主动停下
// func listenSignal() {
// 	c := make(chan os.Signal)
// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
// 	sig := <-c
// 	fmt.Printf("接受到信号:%s", sig.String())
// 	if reader != nil {
// 		reader.Close()
// 	}
// 	os.Exit(0)
// }
