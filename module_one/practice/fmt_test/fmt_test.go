package fmt_test

import (
	"fmt"
	"testing"
)

const pai = 3.14

func TestFmt(t *testing.T) {
	var inter rune = 123
	var str []rune = []rune{'中', '国'}
	var str_two []byte = []byte{101, 64, 65, 66}
	fmt.Println(str)

	fmt.Printf("%s: \n", str_two)

	fmt.Printf("%v\n", inter)
	fmt.Printf("%+v\n", inter)
	fmt.Printf("%#v\n", inter)
	fmt.Printf("%T\n", inter)
	fmt.Printf("%d%%\n", inter) //%%并非占位符，只是对%转义 输出10%

	one := 160
	/**整数格式化**/
	fmt.Printf("二进制表示：%b\n", one)
	fmt.Printf("八进制表示：%o\n", one)
	fmt.Printf("十进制表示：%d\n", one)
	fmt.Printf("十六进制表示(a-f)：%x\n", one)
	fmt.Printf("十六进制表示(A-F)：%X\n", one)
	fmt.Printf("该值对应的Unicode码值：%c\n", one)
	//该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("单引号括起来的go语法字符字面值：%q\n", one)
	//Unicode格式，貌似为U+16进制的表示的数字
	fmt.Printf("Unicode格式(U+000F)：%U\n", one)

	large_num := 123000000.345

	//fmt.Printf("%p", large_num)
	fmt.Printf("%b %e %E %f %F %g %G", large_num, large_num, large_num, large_num, large_num, large_num, large_num)

}
