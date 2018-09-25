package main

import "fmt"

func main() {
	const a, b = 10, 11
	fmt.Println("a = ", a, " b = ", b)
	test(a)
	test2(b, 20, 20, 20, 40, 60, 8)
}
func test(a int) {
	fmt.Println(a)
}

// 多个不确定参数使用args， 必须放在最后
func test2(a int, args ...int) {
	fmt.Println(a)
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
