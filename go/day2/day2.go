package main

import (
	"fmt"
)

//函数其他语言就四种 有参数，没参数，有返回和没有返回四种但是现在用的是go特殊的返回多个参数的函数
func sayHello(name string, age int) (sayName, sayAge string, askYou string) {
	sayName = "我叫" + name
	sayAge = "今年" + fmt.Sprintf("%d", age) + "岁" //看到这里是不是c++的同学很常见啊
	askYou = "您呢?"
	return sayName, sayAge, askYou
}

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
	// var (
	// 	getName string = "22222"
	// 	getAge  string = "33333"
	// 	ask     string = "1111"
	// )
	getName, getAge, ask := sayHello("周周", 31)
	fmt.Println("getName", getName, "getAge", getAge, "ask", ask)
}
