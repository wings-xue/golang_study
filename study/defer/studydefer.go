package main

import "fmt"

func panicStudy() {
	panic("this is study panic")
}

// 为什么添加到函数就不行了?
// panic 能够改变程序的控制流，调用 panic 后会立刻停止执行当前函数的剩余代码，并在当前 Goroutine 中递归执行调用方的 defer；
// recover 可以中止 panic 造成的程序崩溃。它是一个只能在 defer 中发挥作用的函数，在其他作用域中调用不会发挥作用；
// func recoverStudy() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovered in f", r)
// 	}
// }

func main() {
	i := 0
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	panicStudy()
	i++
	return
}

// 总结
// panic可以在defer循环调用
// recorve只有在defer生效
// panic只作用当前gorou
