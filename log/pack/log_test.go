package pack

import "testing"

func TestLog(t *testing.T) {
	LogInit()
	ALoger.Print("test of ALoger")
	ELoger.Print("test of Eloger")
}
