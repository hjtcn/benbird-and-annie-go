package main

import "fmt"

func printHello()  {
	fmt.Println("Hello Golang")
}

func printPoem()  {
	poem := []string{"鹅鹅鹅，", "曲项向天歌，", "白毛浮绿水，", "红掌拨清波。"}

	for _, val := range poem {
		fmt.Println(val)
	}
}

func printFivePointedStar()  {
	fmt.Println("            *")
	fmt.Println("           ***")
	fmt.Println("          *****")
	fmt.Println("         *******")
	fmt.Println("        *********")
	fmt.Println(" ************************")
	fmt.Println("   ********************")
	fmt.Println("     ****************")
	fmt.Println("       ************")
	fmt.Println("     ****************")
	fmt.Println("    ******      ******")
	fmt.Println("   ****            ****")
	fmt.Println(" **                   **")
}

func main()  {
	printHello()
	printPoem()
	printFivePointedStar()
}
