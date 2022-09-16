package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timeStr := now.Format("Mon 02 Jan 2006 15:04:05 MST")
	fmt.Println(timeStr)
}
