package main

import (
	"fmt"
	"github.com/hjtcn/go-base-homework/fatRate"
)

//第一次作业
func main() {
	for {
		fatRate.MainFatRateBody()

		if cont := fatRate.WhetherContinue(); !cont {
			break
		}
	}

	fmt.Println("已退出体脂计算！\n")
}