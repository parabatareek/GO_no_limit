package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	truncTime := now.Truncate(24 * time.Hour)
	fmt.Println(truncTime)
}
