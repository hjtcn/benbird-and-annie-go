package fmt_test

import (
	"fmt"
	"testing"
)

func TestFmt(t *testing.T)  {
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
}
