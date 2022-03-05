package main

import "fmt"

/**
翻转数组
*/
func flipArr(sourceArr []int) []int {
	arrLen := len(sourceArr)

	if arrLen <= 1 {
		return sourceArr
	}

	i := 0
	j := arrLen - 1

	for {
		if i == j || i > j {
			break
		}

		sourceArr[i], sourceArr[j] = sourceArr[j], sourceArr[i]

		i++
		j--
	}

	return sourceArr
}

func main() {
	var a [3][3]int

	a[0] = [3]int{11, 23, 23}
	a[1] = [3]int{23, 45, 56}
	a[2] = [3]int{34, 1, 5}

	fmt.Println(a)

	for i, val := range a {
		//fmt.Println(val)
		for j, item := range val {
			fmt.Printf("a[%d][%d] = %d\n", i, j, item)
		}
	}

	sourceArr := []int{23, 34, 45, 4, 67, 89, 8}

	arr := flipArr(sourceArr)
	fmt.Println(arr)
}
