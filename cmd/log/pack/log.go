package pack

import (
	"be/pkg/constants"
	"log"
	"os"
)

var ALoger *log.Logger
var ELoger *log.Logger
var SLoger *log.Logger

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
