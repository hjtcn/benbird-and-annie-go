package main

import "fmt"

func main() {
	var hello string = "hello golang"
	var num int = 123
	var float_num float64 = 3.4
	fmt.Println(hello, num, float_num)

	hello = "hello 中国"
	fmt.Println(hello)

	/*	33 赋值失败，类型不同不能赋值
		hello = 33
		fmt.Println(hello)*/
}
