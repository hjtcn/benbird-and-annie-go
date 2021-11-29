package main

import "fmt"

func howUseMoney(money int) string {
	var res string
	if money <= 20 {
		res = "点个外卖"
	} else if money <= 200 {
		res = "下个馆子"
	} else if money <= 2000 {
		res = "去米其林"
	} else {
		res = "幸福的不知道干啥了"
	}

	return res
}

func switchUseMoney(money interface{}) string {
	switch newMoney := money.(type) {
	case int:
		fmt.Println(newMoney + 3.0)
		return "this is int"
	case string:
		return "this is string"
	case float64:
		return "this is float64"
	case rune:
		return "this is rune"
	default:
		return "this is nothing"
	}
}

func useFallthrouth(money int)  {
	switch {
	case money >= 0 && money <= 20:
		fmt.Println("点外卖")
	case money > 20 && money <= 200:
		fmt.Println("米其林")
		fallthrough
	case money > 200 && money <= 2000:
		fmt.Println("下馆子")
	default:
		fmt.Println("容我想想")
	}
}

func main()  {
	var money int = 200
	fmt.Println(howUseMoney(money))

	fmt.Println(switchUseMoney(money))

	useFallthrouth(money)
}
