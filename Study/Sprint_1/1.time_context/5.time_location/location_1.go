package main

import (
	"fmt"
	"time"
)

func main() {
	bithday := time.Date(2093, 11, 26, 0, 0, 0, 0, time.Local)
	duration := time.Until(bithday)
	days := int(duration.Hours() / 24)
	fmt.Println(days)
}
