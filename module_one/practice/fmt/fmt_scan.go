package main

import (
	"fmt"
	"os"
)

const pai = 3.14

func main()  {
	var r_one float64
	var r_two float64

	fmt.Println("输入第一个圆的半径:")
	fmt.Scan(&r_one)

	fmt.Println("输入第二个圆的半径:")
	fmt.Scan(&r_two)

	fmt.Println("两圆面积和:", r_one * r_one * pai + r_two * r_two * pai)

	os.Exit(1)
}
