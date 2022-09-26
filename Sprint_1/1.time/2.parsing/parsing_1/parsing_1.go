package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	currentTime, error := time.Parse(time.RFC3339, currentTimeStr)
	if error != nil {
		log.Fatalln(error)
	}
	fmt.Println(currentTime)
}
