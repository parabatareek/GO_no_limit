package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timeStr := now.Format(time.ANSIC)
	fmt.Println(timeStr)
}
