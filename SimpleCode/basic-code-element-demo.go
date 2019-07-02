package main // 指定当前源文件所在的包名

import (
	"fmt"
	"math/rand"
) // 引入一个标准库包

const MaxRand = 16 // 声明一个有名整型常量

// 函数声明
func StatRandomNumbers(numRands int) (int, int) {
	// 声明两个变量（类型都为int，初始值为0）
	var a, b int
	// 一个for循环代码块
	for i := 0; i < numRands; i++ {
		// 一个if-else条件控制代码块
		if rand.Intn(MaxRand) < MaxRand/2 {
			a += 1
		} else {
			b++ // 等价于：b = b + 1
		}
	}
	return a, b // 此函数返回两个结果
}

// main函数，或主函数，是一个程序执行的入口函数
func main() {
	var num = 100
	// 调用上面声明的StatRandomNumber函数
	// 并将结果赋给使用短声明语句声明的两个变量
	x, y := StatRandomNumbers(num)
	// 调用两个内置函数（fmt.Printf()和fmt.Println()）
	fmt.Printf("Result: %d + %d = %d ? ", x, y, num)
	fmt.Println(x+y == num)
}
