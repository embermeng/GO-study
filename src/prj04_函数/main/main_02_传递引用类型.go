package main

import "fmt"

// 参数的类型为指针
func testPt(num *int) {
	*num = 30
	fmt.Println("test---", *num)
}
func main() {
	var num int = 10
	testPt(&num) //传num的地址
	fmt.Println("main---", num)
}
