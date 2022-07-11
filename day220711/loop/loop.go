package main

import "fmt"

func main() {
	// 重复遍历

	// for
	// for init; condition; post { }
	// for condition { }
	// for { }

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum=", sum)

	sum2 := 1
	for sum2 <= 10 {
		sum2 += sum2
	}
	fmt.Println("sum2=", sum2)

	// sum3 := 0
	// for {
	// 	sum3++
	// }

	// For-each range 循环:可以对字符串、数组、切片等进行迭代输出元素

	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

	// for嵌套：同上

	// break

	// continue

	// goto

	// goto label;
	// ..
	// .
	// label: statement;
}
