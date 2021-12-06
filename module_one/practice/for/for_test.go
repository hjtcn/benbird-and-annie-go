package _for

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T)  {
	var str string = "hello golang"

	for i := 1; i <= 100; i++ {
		fmt.Println(i, str)
	}
}
