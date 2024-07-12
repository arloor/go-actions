package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const right = "right"

type Student struct {
	Name string
}

func GenStudent(stuObj *Student) {
	s := Student{
		Name: right,
	}
	// 正确的赋值方式
	*stuObj = s
}

func EditStudent(stuObj *Student) {
	if stuObj!=nil{
		stuObj.Name = right
	}
}

// 错误的
// golangci-lint应该有提示 “ineffectual assignment to stuObj (ineffassign)”
func GenStudentWrong(stuObj *Student) {
	s := Student{
		Name: right,
	}
	// 这样赋值，不会改变实参！！！
	stuObj = &s
}

func TestTS(t *testing.T) {
	var s1 Student
	GenStudent(&s1)
	assert.Equal(t, s1.Name, right)

	var s2 Student
	GenStudentWrong(&s2)
	assert.NotEqual(t, s2.Name, right)

	var s3 Student
	EditStudent(&s3)
	assert.Equal(t, s3.Name, right)
}
