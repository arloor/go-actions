package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitStruct(t *testing.T) {
	var some Some
	assert.NotNil(t, some)

	some1 := Some{}
	assert.NotNil(t, some1)
	assert.Equal(t,some1.field,0)
}

func TestInitSlice(t *testing.T) {
	var nilSlice []string
	assert.Nil(t, nilSlice)

	slice := []string{} // 只能使用短格式赋值
	assert.NotNil(t, slice)
}

func TestInitMap(t *testing.T) {
	var nilMap map[int]int
	assert.Nil(t, nilMap)

	aMap := map[int]string{} // 只能使用短格式赋值
	assert.NotNil(t, aMap)
}
