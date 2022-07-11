package main

import "fmt"

func main() {
	// 描述：条件语句控制代码执行流程

	// if
	var number1 int = 10
	fmt.Println("number1 的值为", number1)
	if number1 < 20 {
		fmt.Println("number1 小于 20")
	}
	// if...else
	if number1 < 20 {
		fmt.Println("number1 小于 20")
	} else {
		fmt.Println("number1 大于等于 20")
	}
	// if 嵌套
	if number1 < 20 {
		if number1 < 15 {
			fmt.Println("number1 小于 15")
		}
	}
	// switch
	switch number1 {
	case 10:
		fmt.Println("number1 为 10")
	case 15:
		fmt.Println("number1 为 10")
	case 20:
		fmt.Println("number1 为 10")
	default:
		fmt.Println("都不是")
	}
	// select
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
