package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 数组赋值是值拷贝（value copy）
func TestArray(t *testing.T) {
	array := [1]int{1}
	array1 := array
	array[0] = 2
	assert.NotEqual(t, array[0], array1[0])

}

// 切片赋值是引用传递（reference assignment）
func TestSlice(t *testing.T) {
	slice := []int{1}
	slice1 := slice
	slice[0] = 2
	assert.Equal(t, slice[0], slice1[0])
}

// map赋值是引用传递（reference assignment）
func TestMap(t *testing.T) {
	map0 := map[int]int{0: 1}
	map1 := map0
	map0[0] = 2
	assert.Equal(t, map0[0], map1[0])
}

// string字符串赋值是值拷贝（value copy）
// golang的string是不可变类型
// 但是看到一些说并发读写string导致panic的情况，后续再研究
func TestString(t *testing.T) {
	string0 := "0"
	string1 := string0
	string0 = "1"
	assert.NotEqual(t, string0, string1)
}

type Some struct {
	field int
}

// struct赋值是值拷贝（value copy）
// 通过struct的指针可以访问和修改指针所指向的结构体的字段（用途，函数中传递struct指针以修改struct内容）
// 如果不需要修改struct内容，并且struct的大小不大（拷贝成本不大），则函数参数直接传struct可以避免“逃逸”，对gc更加友好
func TestStruct(t *testing.T) {
	some0 := Some{1}
	some1 := some0
	some0_pointer := &some0
	some0.field = 2
	assert.NotEqual(t, some0.field, some1.field)
	assert.Equal(t, some0.field, some0_pointer.field)
}

// channel是引用传递
func TestChannel(t *testing.T) {
	ch0 := make(chan int, 1) // 创建一个缓冲大小为1的通道
	ch1 := ch0               // 将ch0赋值给ch1

	ch0 <- 1      // 向通道ch0发送值1
	val0 := <-ch0 // 从通道ch0接收值
	assert.Equal(t, val0, 1)

	ch1 <- 2      // 向通道ch1发送值2
	val1 := <-ch0 // 从通道ch0接收值（因为ch0和ch1是同一个通道）
	assert.Equal(t, val1, 2)
}

// 数值类型、bool类型是值拷贝
func TestNumber(t *testing.T) {
	a := 0
	b := a
	a = 1
	assert.NotEqual(t, a, b)

	a_pointer := &a
	a = 3
	assert.Equal(t, a, *a_pointer)
}

// func类型是引用传递
// 但是！！！不共享值，反直觉，不太懂
func TestFunc(t *testing.T) {
	a := func() int { return 0 }
	b := a
	a = func() int { return 1 }

	assert.NotEqual(t, a(), b())
}
