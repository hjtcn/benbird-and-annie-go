package main

import (
	fatRate "benbird-and-annie-go/module_two/practice/refactorFatRate"
	"fmt"
)

func main() {
	for {
		fatRate.MainFatRateBody()

		if cont := fatRate.WhetherContinue(); !cont {
			break
		}
	}

	fmt.Println("已退出体脂计算！\n")
}
