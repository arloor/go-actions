package main

import (
	"fmt"
	"time"
)

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
	k := c
	println(fmt.Sprintf("fullPath: %s", k))
}
