package kafka

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	Init()
	defer func() {
		if err := logServer.Close(); err != nil {
			log.Println("Failed to close server", err)
		}
	}()

	AccessLog("Hello World")
	ErrorLog("Hello World Again")
}
