package main

import (
	"fmt"
	"testing"
)

func TestUseVar(t *testing.T) {
	var hen int = 4
	var shu int = 5
	fmt.Println(hen * shu)

	var mianji = hen * shu
	fmt.Println(mianji)

	/* 不允许的操作，变量类型不同，不能进行运算
	var chang, kuan = 4, 5.1
	mianji = chang + kuan
	fmt.Println(mianji)*/

	var width, high int
	fmt.Println(width, high)

	var name string
	fmt.Println(name)

	var mapA map[string]string
	fmt.Println(mapA)
}
