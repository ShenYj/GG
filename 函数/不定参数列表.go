package main

import "fmt"

func main() {
	test(10, 230, 4, 235, 6543, 23)
}

func test(a int, args ...int) {
	fmt.Println(a)
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
