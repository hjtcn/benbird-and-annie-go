package scope

import "testing"

func TestScope(t *testing.T) {
	tall, weight := 1.7, 1.7
	t.Log(tall, weight)
}

func TestIfScope(t *testing.T) {
	if v := giveRes(); v == 0 {
		t.Log(v)
	} else {
		t.Log(v)
	}

	//t.Log(v) v不存在
}

func giveRes() int {
	return 1
}

/**
作用域：
	1.花括号中定义的变量，不能为花括号外部使用

全局变量
局部变量
*/
