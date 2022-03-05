package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestInterVarType(t *testing.T) {
	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))

	a2 := 3
	fmt.Println(reflect.TypeOf(a2))

	a3 := 3.0
	fmt.Println(reflect.TypeOf(a3))

	a4 := &a3
	fmt.Println(reflect.TypeOf(a4))

	//变量类型转换及其代价
	var f1 float64 = 1.534
	var i1 int = int(f1)
	fmt.Println("f1:", f1, "i1:", i1)

	var a6 uint = math.MaxUint64
	var a7 int = int(a6)
	fmt.Println(a6, a7)

	var a8 uint = 1<<64 - 1
	fmt.Println(a8)
}
