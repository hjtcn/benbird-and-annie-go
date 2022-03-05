package main

import (
	"fmt"
)

const pai = 3.1415

func main() {
	var r_one float64
	var r_two float64

	fmt.Println("输入第一个圆的半径:")
	fmt.Scan(&r_one)

	fmt.Println("输入第二个圆的半径:")
	fmt.Scan(&r_two)

	fmt.Printf("两圆面积和:%.3f", r_one*r_one*pai+r_two*r_two*pai)
	/**
	1 宽度：整个数的最小宽度
	3 精度：小数点后的个数
	*/
}
