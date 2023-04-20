package kafka

import (
	"be/pkg/constants"
	"be/pkg/kafka"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

var topicLoggerMap map[string]*log.Logger
var pureALoger *log.Logger
var pureELoger *log.Logger

func InitTopicLoggerMap() {
	LogInit()
	topicLoggerMap = make(map[string]*log.Logger)
	topicLoggerMap[constants.KafkaAccessLogTopic] = pureALoger
	topicLoggerMap[constants.KafkaErrorLogTopic] = pureELoger
}

func LogInit() {
	alog, err := os.OpenFile(constants.ALogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	elog, err := os.OpenFile(constants.ELogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	pureALoger = log.New(alog, "", 0)
	pureELoger = log.New(elog, "", 0)
}

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			logger, ok := topicLoggerMap[message.Topic]
			if !ok {
				kafka.ELoger.Println("logger get fail")
			}
			logger.Println(string(message.Value))
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

func CreateConsumer(brokerList []string, topic []string, groupID string) {
	keepRunning := true
	version, err := sarama.ParseKafkaVersion(constants.KafkaConsumerVersion)
	if err != nil {
		kafka.ELoger.Panicf("Error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version

	consumer := Consumer{
		ready: make(chan bool),
	}
	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(brokerList, groupID, config)
	if err != nil {
		kafka.ELoger.Panicf("Error creating consumer group client: %v", err)
	}

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, topic, &consumer); err != nil {
				kafka.ELoger.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			kafka.ALoger.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			kafka.ALoger.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		kafka.ELoger.Panicf("Error closing client: %v", err)
	}
}
