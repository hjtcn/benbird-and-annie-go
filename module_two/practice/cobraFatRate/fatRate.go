package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	/*var (
		name   string
		sex    string
		tall   string
		weight string
		age    string
	)

	arguments := os.Args
	fmt.Println(arguments)

	name = arguments[1]
	sex = arguments[2]
	tall = arguments[3]
	weight = arguments[4]
	age = arguments[5]*/

	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	cmd := &cobra.Command{
		Use:   "fatTate",
		Short: "计算个人体脂",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			//获取信息
			fmt.Println("name: ", name)
			fmt.Println("sex: ", sex)
			fmt.Println("tall: ", tall)
			fmt.Println("weight: ", weight)
			fmt.Println("age: ", age)

			//计算

			//评估结果

		},
	}

	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "姓名")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.Execute()
}
