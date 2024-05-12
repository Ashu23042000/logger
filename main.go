package main

import (
	"github.com/Ashu23042000/logger/logger"
)

func main() {

	// file, _ := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log := logger.New(nil)

	log.Debug("Hello Bhai")
	log.Debugf("%s %d", "Hello bhai", 23)

	log.Info("Hello Bhai")
	log.Infof("%s %d", "Hello bhai", 23)

	log.Warn("Hello Bhai")
	log.Warnf("%s %d", "Hello bhai", 23)

	log.Error("Hello Bhai")
	log.Errorf("%s %d", "Hello bhai", 23)
}
