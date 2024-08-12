package main

import "fmt"

func main() {
	//基本数据类型相互转换
	//Go 在不同类型的变量之间赋值时需要显式转换，不能自动转换
	var i int32 = 100
	var n1 float64 = float64(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)

	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n", i, n1, n2, n3)
}
