package main

// go run main/string_data_race_panic.go

import (
	"fmt"
	"time"
)

// 并发读写string，会panic
func main() {
	fullPath := "init"
	go func() { // goroutine 不断读取fullpath
		for i := 1; i < 10000; i++ {
			request(fullPath)
		}
	}()

	for { // main goroutine会不断修改fullPath.
		fullPath = ""
		time.Sleep(10 * time.Nanosecond)
		fullPath = "/test/test/test"
		time.Sleep(10 * time.Nanosecond)
	}

}

func request(c string) { // 这里传参，有一次拷贝，会做feild（string的poiner和len）的赋值
	fmt.Printf("fullPath: %s\n", c)
	// 或者下面的json Marshal，也会panic
	// _, _ = json.Marshal(c)
}
