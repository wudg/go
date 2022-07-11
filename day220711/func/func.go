package main

import (
	"fmt"
	"math"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

// func function_name( [parameter list] ) [return_types] {
// 	  函数体
//  }

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 函数返回多个值

func swap(x, y string) (string, string) {
	return y, x
}

// 值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数
// 引用传递：指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改将影响实际参数

// 函数闭包，创建函数getSequence()，返回另外一个函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	fmt.Println("最大值是：", max(10, 15))

	a, b := swap("world", "hello")
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	// 函数作为实参
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	// 闭包调用（类匿名方法）
	fmt.Println(getSquareRoot(9))

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	// 调用结构体的函数
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积=", c1.getArea())
}

// 定义结构体的内部函数
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
