package main

import "fmt"

func suggest(fatRateMsg string) string {
	switch fatRateMsg {
	case "体脂率有误":
		return "请检验个人信息是否输入正确"
	case "偏瘦":
		return "偏瘦, 多吃蛋白"
	case "标准":
		return "标准, 注意保持哟"
	case "偏重":
		return "偏重, 需要进行适当运动哟"
	case "肥胖":
		return "肥胖, 你要多加运动哟"
	case "严重肥胖":
		return "严重肥胖, 需要看下医生了"
	default:
		return fatRateMsg
	}
}

func getSexWeight(sex string) int {
	if sex == "男" {
		return 1
	}

	return 0
}

func getBmi(weight, tall float64) float64 {
	return weight / (tall * tall)
}

func getFatRate(bmi float64, age, sexWeight int) float64 {
	return (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
}

func getFatRateResult(fatRate float64, age, sexWeight int) string {
	fatRateMsg := "没什么建议"
	if sexWeight == 1 {
		if age >= 18 && age <= 39 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.1 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.16 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.21 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.26 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		} else if age <= 59 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.11 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.17 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.22 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.27 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		} else if age >= 60 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.13 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.19 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.24 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.29 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		}
	} else {
		if age >= 18 && age <= 39 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.2 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.27 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.34 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.39 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		} else if age <= 59 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.21 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.28 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.35 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.4 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		} else if age >= 60 {
			if fatRate < 0 {
				fatRateMsg = "体脂率有误"
			} else if fatRate <= 0.22 {
				fatRateMsg = "偏瘦"
			} else if fatRate <= 0.29 {
				fatRateMsg = "标准"
			} else if fatRate <= 0.36 {
				fatRateMsg = "偏重"
			} else if fatRate <= 0.41 {
				fatRateMsg = "肥胖"
			} else {
				fatRateMsg = "严重肥胖"
			}
		}
	}

	return fatRateMsg
}

func main() {
	var goOn string = "yes"

	for goOn == "yes" {
		var weight float64
		var tall float64
		var age int
		var sex string

		fmt.Print("请输入您的年龄(岁): ")
		fmt.Scanf("%d", &age)

		for age < 18 {
			fmt.Print("只能计算18岁以上的呦，请再次输入年龄(岁): ")
			fmt.Scanf("%d", &age)
		}

		fmt.Print("请输入您的性别(男/女): ")
		fmt.Scanf("%s", &sex)

		for sex != "男" && sex != "女" {
			fmt.Print("请输入有效的性别(男/女): ")
			fmt.Scanf("%s", &sex)
		}

		fmt.Print("请输入您的体重(公斤): ")
		fmt.Scanf("%f", &weight)

		for weight < 30 {
			fmt.Print("请输入您的有效体重(公斤): ")
			fmt.Scanf("%f", &weight)
		}

		fmt.Print("请输入您的身高(米): ")
		fmt.Scanf("%f", &tall)

		for tall < 0.5 || tall > 3.0 {
			fmt.Print("请输入您的身高(米): ")
			fmt.Scanf("%f", &tall)
		}

		sexWeight := getSexWeight(sex)
		bmi := getBmi(weight, tall)
		fatRate := getFatRate(bmi, age, sexWeight)

		fatRateMsg := getFatRateResult(fatRate, age, sexWeight)
		suggestInfo := suggest(fatRateMsg)

		fmt.Println("我的体脂率:", fatRate)
		fmt.Println(suggestInfo)

		fmt.Println("")
		fmt.Print("请问您是否还要继续计算体脂率(yes/no): ")
		fmt.Scanf("%s", &goOn)

		err_num := 1
		for goOn != "yes" && goOn != "no" {
			if err_num > 3 {
				break
			}
			fmt.Print("请问您是否还要继续计算体脂率(yes/no): ")
			fmt.Scanf("%s", &goOn)
			err_num++
		}

		if goOn == "no" {
			break
		}
	}

	fmt.Println("已退出体脂计算！\n")
}

/**
实现单个人的体脂率计算
但是存在的问题是：输入信息时，如果输入的值类型不是定义的类型，则Scanf阻塞异常
*/
