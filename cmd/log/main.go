package main

import (
	"be/cmd/log/dal"
	"be/cmd/log/dal/kafka"
	"be/pkg/constants"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	// access log
	go kafka.CreateConsumer([]string{"127.0.0.1:9092"}, []string{constants.KafkaAccessLogTopic}, constants.KafkaAccessLogGroupID)
	go kafka.CreateConsumer([]string{"127.0.0.1:9092"}, []string{constants.KafkaErrorLogTopic}, constants.KafkaErrorLogGroupID)

	select {}
}
