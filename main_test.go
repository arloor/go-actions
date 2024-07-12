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
	string1:=string0
	string0="1"
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
