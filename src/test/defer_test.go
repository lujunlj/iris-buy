package test

import (
	"fmt"
	"testing"
)

func ResultValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}

func ReturnResultValues() (result int) {
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return
}

func Test_ResultValues(t *testing.T) {
	//res1 := ResultValues()
	//fmt.Println(res1)
	res2 := ReturnResultValues()
	fmt.Println(res2)
}
