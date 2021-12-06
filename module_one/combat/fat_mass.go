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

func main() {
	var weight float64 = 58
	var tall float64 = 1.63
	var bmi float64 = weight / (tall * tall)

	var age int = 24
	var sexWeight int
	var sex string = "女"

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}

	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("我的体脂率:", fatRate)

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

	suggestInfo := suggest(fatRateMsg)

	fmt.Println(suggestInfo)
}
