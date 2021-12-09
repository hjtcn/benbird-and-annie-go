package main

import "fmt"

func main()  {

	sourceArr := [...]int{23, 34, 45, 4 ,67, 89, 8}
	arrLen := len(sourceArr)

	if arrLen >= 1 {
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
	}

	fmt.Println(sourceArr)
}
