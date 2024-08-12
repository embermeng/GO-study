package main

import "fmt"

// getSum返回一个函数，且这个函数的参数是一个int，返回值是一个int
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	f := getSum()
	fmt.Println(f(1)) // 1
	fmt.Println(f(2)) // 3
	fmt.Println(f(3)) // 6
}
