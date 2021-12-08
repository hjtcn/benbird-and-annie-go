package main

import "fmt"

func main() {
	/**
	  问题分析：
	  	判断两直线平行 == 判断两直线斜率是否相等
	  	两点确定一条直线(x1, y1) (x2, y2)
	  	斜率 k = (y2 - y1) / (x2 - x1)
	  特殊情况:
	  	1.直线平行于y轴(x2 - x1 == 0)
	  	2.直线平行于x轴
	*/

	for {
		line_one := make([]float64, 4)
		line_two := make([]float64, 4)

		fmt.Print("请输入line_one的两个坐标(x1, y1, x2, y2)：")
		fmt.Scan(&line_one[0], &line_one[1], &line_one[2], &line_one[3])

		fmt.Print("请输入line_two的两个坐标(x1, y1, x2, y2)：")
		fmt.Scan(&line_two[0], &line_two[1], &line_two[2], &line_two[3])

		line_one_xd := line_one[2] - line_one[0]
		line_one_yd := line_one[3] - line_one[1]
		line_two_xd := line_two[2] - line_two[0]
		line_two_yd := line_two[3] - line_two[1]

		slopeMsg := ""
		if (line_one_xd == 0) || (line_two_xd == 0) {
			if line_one_xd == line_two_xd {
				slopeMsg = "直线line_one和line_two平行,并且都与y轴平行"
			} else {
				slopeMsg = "直线line_one和line_two不平行"
			}
		} else {
			k_one := line_one_yd / line_one_xd
			k_two := line_two_yd / line_two_xd

			fmt.Printf("line_one斜率:%f line_two斜率:%f\n", k_one, k_two)
			if k_two == k_one {
				if k_two == 0 {
					slopeMsg = "直线line_one和line_two平行, 并且与X轴平行"
				} else {
					slopeMsg = "直线line_one和line_two平行"
				}
			} else {
				slopeMsg = "直线line_one和line_two不平行"
			}
		}

		fmt.Println(slopeMsg)

		var goOn string
		fmt.Println("")
		fmt.Print("请问您是否还要继续判断两直线平行(yes/no): ")
		fmt.Scanf("%s", &goOn)

		err_num := 1
		for goOn != "yes" && goOn != "no" {
			if err_num >= 3 {
				break
			}
			fmt.Print("请问您是否还要继续判断两直线平行(yes/no): ")
			fmt.Scanf("%s", &goOn)
			err_num++
		}

		if goOn == "no" {
			break
		}
	}

	fmt.Println("已退出程序！")
}
