package main

import "fmt"

func main() {
	// 变量定义
	// var identifier type

	// 标准声明变量
	var basic string = "wudiguang"

	// 基础变量默认值
	// 数值 ---> 0
	// bool ---> false
	// 字符串 ---> 空串
	// 其他 ---> nil
	var mynumber int
	var mybool bool
	var mystring string
	// wudiguang 0 false
	fmt.Println(basic, mynumber, mybool, mystring)

	// 其他变量默认值
	var mypoint *int
	var myarray []int
	var mymap map[string]int
	var mychan chan int
	var myfunc func(string) int
	var myerror error
	// <nil> [] map[] <nil> <nil> <nil>
	fmt.Println(mypoint, myarray, mymap, mychan, myfunc, myerror)

	// 根据值执行判断变量类型
	var specialbool = true
	fmt.Println(specialbool)

	// 变量声明后，不允许使用 := 赋值
	// 下面代码会产生编译错误
	// var intVal int
	// intVal := 3
	// fmt.Println(intVal)

	// 多变量声明
	var multu1, multi2 = 1, 4
	fmt.Println(multu1, multi2)

	// 初始化并赋值，只能在函数体中出现
	transfer1, transfer2 := multi2, multu1
	fmt.Println(transfer1, transfer2)
}
