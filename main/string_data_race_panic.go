package main

// go run main/string_data_race_panic.go

import (
	"encoding/json"
	"time"
)

// 并发读写string，会panic
func main() {
	fullPath := "init"
	go func() {
		for i := 1; i < 10000; i++ {
			request(fullPath)
		}
	}()

	for {
		fullPath = ""
		time.Sleep(10 * time.Nanosecond)
		fullPath = "/test/test/test"
		time.Sleep(10 * time.Nanosecond)
	}

}

func request(c string) {
	_, _ = json.Marshal(c)
	// 或者下面的读取
	// println(fmt.Sprintf("fullPath: %s", c))
}
