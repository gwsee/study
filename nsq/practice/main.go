package practice

import (
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

//写个消费者
//消费者实现的目标是 能够读取到nsq的消息，且不同的消费者 能够读取不同的消息，未读的消息有明确的标识
//s 分别使用2 或者3 来代表不同的消费者，在没有被消费或者在服务断掉 重启的时候 能够再次监听成功
func customer(s int) {
	topic := "test"
	channel := "testgo"
	lookupAddr := "localhost:4161"
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatalln(err)
	}
	//定义如何处理消息
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		i, _ := strconv.Atoi(string(message.Body))
		if i%s == 0 {
			log.Println(i)
			message.Finish()
			return nil
		}
		log.Println(string(message.Body), "尚且不能被消费")
		return errors.New(string(message.Body)) //不被消费的数据会
	}))
	err = consumer.ConnectToNSQLookupd(lookupAddr)
	if err != nil {
		log.Fatalln(err)
	}
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	consumer.Stop()
}

//测试所有的消费者都能访问同样的topic数据--消息只能执行一次 在没有返回错误的时候，所以想给整体发送消息应该在外部去实现
func customerAll() {
	topic := "test"
	channel := "testgo"
	lookupAddr := "localhost:4161"
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatalln(err)
	}
	mp := make(map[string]int)
	//定义如何处理消息
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		str := string(message.Body)
		fmt.Println(str)
		mp[str] = mp[str] + 1
		if mp[str] >= 2 {
			message.Finish()
		}
		return nil
	}))
	err = consumer.ConnectToNSQLookupd(lookupAddr)
	if err != nil {
		log.Fatalln(err)
	}
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
	consumer.Stop()
}

//写个生成者
//1：生成制定主题
//2：给主题里面添加消息
//3：查询主题里面的消息有哪些

func producer() {
	topic := "test"
	addr := "localhost:4150"
	config := nsq.NewConfig()
	productor, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 10)
		log.Println(i)
		productor.Publish(topic, []byte(fmt.Sprintf("%d", i)))
	}
	productor.Stop()
}
