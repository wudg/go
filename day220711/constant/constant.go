package main

import (
	"fmt"
)

func main() {
	// 常量定义(显式和隐式)
	// const identifier [type] = value

	const LENGTH int = 10
	const WIDTH int = 5

	fmt.Printf("面积为：%d", LENGTH*WIDTH)

	// 常量还可以用作枚举
	const (
		UNKNOWN = 0
		FEMALE  = 1
		MALE    = 2
	)
	fmt.Println(UNKNOWN, FEMALE, MALE)

	// iota：特殊的常量，可以被认为是一个可以被编译器修改的常量。const中每新增一行常量声明就使iota计数一次
	const (
		fitst  = iota // first = 0
		second = iota // second = first + 1
		three  = iota // three = second + 1
	)

	const (
		iota_a = iota
		iota_b
		iota_c
		iota_d = "ha"
		iota_e
		iota_f = 100
		iota_g
		iota_h = iota
		iota_i
	)
	fmt.Println(iota_a, iota_b, iota_c, iota_d, iota_e, iota_f, iota_g, iota_h, iota_i)
}
