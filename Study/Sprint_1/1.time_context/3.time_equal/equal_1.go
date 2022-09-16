package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	inputTime, err := time.Parse(time.RFC3339, "2021-09-19T15:59:41+03:00")
	if err != nil {
		log.Fatalln(err)
	}

	nowTime := time.Now()
	fmt.Println("Is", nowTime, "before", inputTime, "? Answer:", nowTime.Before(inputTime))
	fmt.Println("Is", nowTime, "after", inputTime, "? Answer:", nowTime.After(inputTime))
	fmt.Println("Is", nowTime, "equal", inputTime, "? Answer:", nowTime.Equal(inputTime))
}
