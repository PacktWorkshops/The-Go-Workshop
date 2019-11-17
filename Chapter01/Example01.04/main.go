package main

import (
	"fmt"
	"time"
)

func main() {
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}
