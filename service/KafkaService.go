package service

import (
	"dataPlatform/model"
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
)

var Client sarama.SyncProducer

func InitKafka() {
	fmt.Println("====>INIT KAFKA!")
	//初始化配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//生产者
	kAddr:= model.ConfigParam["kafka-addr"]
	addr:=strings.Split(kAddr,",")
	client, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer close,err:", err)
	}
	Client = client
	//defer client.Close()
}

func SendMsg(client sarama.SyncProducer, msgs string) {
	//创建消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = model.ConfigParam["kafka-topic"]
	msg.Value = sarama.StringEncoder(msgs)
	//发送消息
	//pid, offset, err := client.SendMessage(msg)
	_, _, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}
	fmt.Println("发送消息：" + msgs)
	//fmt.Printf("pid:%v offset:%v\n,", pid, offset)
}
