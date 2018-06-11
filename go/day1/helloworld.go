package main //和java一样就是是包路径也就是文件夹路径

import (
	"fmt" //类似于c语言的#include java的import 一样
)

//func 函数 主函数
func main() {
	//打印函数
	fmt.Println("helloword")
	//变量
	var a string
	a = "123455"
	fmt.Printf("a = %s\n", a)
	var b = "11111111111"
	fmt.Printf("b = %s\n", b)
	c := "您好" //这种是go的特殊变量方式必须要赋值
	fmt.Printf("c = %s\n", c)
	var (
		d = "1"
	)
	fmt.Printf("d = %s\n", d)
	//常量
	const HELLO = "hello"
	fmt.Printf("d = %s\n", HELLO)
	//复杂类型
	//数组
	var arr_1 [3]int = [3]int{1, 2, 3}
	fmt.Println("arr_1", arr_1)
	arr_2 := []int{4, 5, 6}
	fmt.Println("arr_2", arr_2)
	//切片
	var slice_1 []int = []int{1, 2, 3}
	fmt.Println("slice_1", slice_1)
	slice_2 := arr_1[0:2]
	fmt.Println("slice_2", slice_2)
	slice_3 := slice_1[0:1]
	fmt.Println("slice_3", slice_3)
	//map就是和所有的都类似
	var map_1 map[string]string
	map_1["one"] = "one"
}
