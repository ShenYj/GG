package main

import "fmt"

func main() {
	test2(1231, 3434, 5984, 34, 556, 43)
}

func test1(a ...int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func test2(args ...int) {
	fmt.Print("test2 调用Test1")
	//test1(args...) // 全部传给test1
	test1(args[2:]...) // 从args[2]始传递， 包括args[2]自身
	test1(args[:2]...) // 从args[0] ~args[2], 不包括args【2】
}
