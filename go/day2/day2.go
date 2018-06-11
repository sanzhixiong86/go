package main

import (
	"fmt"
)

func main() {
	//语句
	x := 1
	if x < 3 {
		fmt.Println("小于3")
	} else if x == 3 {
		fmt.Println("等于3")
	}

	//第二种比较特殊
	if y := 3; y < 3 { //首先副职然后进行操作
		fmt.Println("小于3")
	} else {
		fmt.Println("其他")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("11111")
	}

	a := 0
	//while
	for a < 5 {
		fmt.Println("其他")
		a++
	}
	//最后一种类似于foreach
	m := map[string]int{"e": 1, "f": 2}
	for k, v := range m {
		fmt.Println("k,v", k, v)
	}
}
