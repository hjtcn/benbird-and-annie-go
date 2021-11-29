package main

import (
	"fmt"
	"testing"
)

func TestVar(t *testing.T)  {
	var hello string = "hello world"
	fmt.Println(hello)

	float_3 := 1.235
	fmt.Println(float_3)

	var int4, int5 = 44, 55
	fmt.Println(int4, int5)

	float4, float5 := 4.0, 5.0
	fmt.Println(float4, float5)

	var (
		int6, int7 = 6, 7
	)
	fmt.Println(int6, int7)

}
