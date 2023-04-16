package kafka

import (
	"be/pkg/constants"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

var ALoger *log.Logger
var ELoger *log.Logger
var SLoger *log.Logger

// 初始化日志描述符
func LogInit() {
	alog, err := os.OpenFile(constants.ALogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	elog, err := os.OpenFile(constants.ELogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	slog, err := os.OpenFile(constants.SLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	ALoger = log.New(alog, "", log.LstdFlags|log.Llongfile|log.LUTC)
	ELoger = log.New(elog, "", log.LstdFlags|log.Llongfile|log.LUTC)
	SLoger = log.New(slog, "", log.LstdFlags|log.Llongfile|log.LUTC)
}

type LogServer struct {
	LogProducer sarama.AsyncProducer
}

func NewLogServer(brokerList []string) *LogServer {
	return &LogServer{
		LogProducer: newLogCollector(brokerList),
	}
}

func (s *LogServer) ErrorLog(err string) {
	s.LogProducer.Input() <- &sarama.ProducerMessage{
		Topic: constants.KafkaErrorLog,
		Key:   nil,
		Value: sarama.StringEncoder(err),
	}
}

func (s *LogServer) AccessLog(str string) {
	s.LogProducer.Input() <- &sarama.ProducerMessage{
		Topic: constants.KafkaAccessLog,
		Key:   nil,
		Value: sarama.StringEncoder(str),
	}
}

func (s *LogServer) Close() error {
	if err := s.LogProducer.Close(); err != nil {
		ELoger.Println("Failed to shut down data collector cleanly", err)
	}

	return nil
}

// 日志采用异步
func newLogCollector(brokerList []string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		ELoger.Fatalln("Failed to start Sarama producer:", err)
	}

	go func() {
		for err := range producer.Errors() {
			ELoger.Println("Failed to write access log entry:", err)
		}
	}()

	return producer
}
