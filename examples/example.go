package main

import (
	"os"

	"github.com/sagargaikwad2000/logger/logger"
)

func main() {

	file, _ := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log := logger.New(file, "Info", false)

	log.Info(
		"user created",
		"user_id", 123,
		"email", "john@gmail.com",
	)

	log.Debug(
		"user created",
		"user_id", 123,
		"email", "john@gmail.com",
	)

	log.Warn(
		"user created",
		"user_id", 123,
		"email", "john@gmail.com",
	)

	log.Error(
		"user created",
		"user_id", 123,
		"email", "john@gmail.com",
	)

}
