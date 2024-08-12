// 声明整个文件所在的包，每个go文件必须有归属的包
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 使用变量的三种方式
	// 1：指定变量类型，不赋值则使用默认值
	var a int = 1
	fmt.Println("a=", a)

	// 2：类型推导
	var b = 10.11
	fmt.Println("b=", b)

	// 3：省略var，:=左边变量不能重复声明
	c := "jerry"
	fmt.Println("c=", c)

	// 多变量声明
	var d1, d2, d3 string
	fmt.Println("d1=", d1, "d2=", d2, "d3=", d3)
	var e1, e2, e3 = 100, "tom", "jerry"
	fmt.Println("e1=", e1, "e2=", e2, "e3=", e3)
	f1, f2, f3 := 100.11, "tom", "jerry"
	fmt.Println("f1=", f1, "f2=", f2, "f3=", f3)

	// 一次性定义多个全局变量
	var (
		g1 = 300
		g2 = 100
		g3 = "marry"
	)
	fmt.Println("g1=", g1, "g2=", g2, "g3=", g3)
	fmt.Printf("g3的类型：%T，g3占用的字节数：%d\n", g3, unsafe.Sizeof(g3))
}
