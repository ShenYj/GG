package main

import "fmt"

func main() {
	a := myFunc()
	fmt.Println(a)

	b, c, d := myFunc3()
	fmt.Print(b)
	fmt.Print(c)
	fmt.Print(d)
}

// 有一个返回值函数
func myFunc() int {
	return 10
}

// 有一个返回值， 给返回值起一个变量名（GO语言推荐写法）
func myFunc1() (result int) {
	return 10
}

// 有一个返回值， 给返回值起一个变量名（GO语言推荐写法）
func myFunc2() (result int) {
	result = 10 // 常用写法
	return
}

// 有多个返回值
func myFunc3() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}
