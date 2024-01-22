package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	for true {
		min := 1
		max := 5
		currentTime := time.Now().Format("2006.01.02 15:04:05\n")
		randomMinute := rand.Intn(max-min) + min

		f, _ := os.OpenFile("/var/log/nginx/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.Write([]byte(currentTime))
		f.Close()

		time.Sleep(time.Duration(randomMinute) * time.Minute)
	}
}

// GOOS=linux GOARCH=amd64 go build log-genearator.go
