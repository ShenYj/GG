package main

import "fmt"

func main() {
	const {
		// 1.常亮自动生成器，每一行自动累加1
		// 2.iota给常亮赋值使用
		// 3. iota遇到const 重置为0
		// 4. 可以只写一个iota
		// 5. 如果在同一行， 值都一样
		a = iota // 0
		b = iota // 1
		c = iota // 2
		d, e = iota  // 3
	}
	const {
		// 1.常亮自动生成器，每一行自动累加1
		// 2.iota给常亮赋值使用
		// 3. iota遇到const 重置为0
		// 4. 可以只写一个iota
		a1 = iota   // 0
		b1 			// 1
		c1 			// 2
	}
	
	const d = iota // d = 0
}