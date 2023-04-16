package kafka

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	LogInit()
	s := NewLogServer([]string{"127.0.0.1:9092"})
	defer func() {
		if err := s.Close(); err != nil {
			log.Println("Failed to close server", err)
		}
	}()

	s.AccessLog("Hello World")
	s.ErrorLog("Hello World Again")
}
