package main

import (
	"github.com/sagargaikwad2000/logger/logger"
)

func main() {

	// file, _ := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	log := logger.New(nil, "Info ")

	log.Info(
		"user created",
		"user_id", "123",
		"email", "john@gmail.com",
	)
}
