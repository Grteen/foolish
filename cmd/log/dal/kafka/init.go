package kafka

import (
	"be/pkg/kafka"
)

func Init() {
	kafka.LogInit()
	InitTopicLoggerMap()
}
