package main

import "fmt"

func main() {
	test2(1231, 3434, 5984)
}

func test1(a ...int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func test2(args ...int) {
	fmt.Print("test2 调用Test1")
	test1(args...)
}
