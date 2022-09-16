package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	timeStr := now.Format("1.2.06 3:4:5 -07 MST")
	fmt.Println(timeStr)
}
