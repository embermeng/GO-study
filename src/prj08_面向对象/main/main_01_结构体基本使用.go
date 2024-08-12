package main

import "fmt"

type Teacher struct {
	// 变量名字大写，外界就可以访问
	Name   string
	Age    int
	School string
}

func main() {
	// 创建老师结构体的实例(对象)
	var t1 Teacher  // var a int
	fmt.Println(t1) // 默认值：{ 0 }
	t1.Name = "马士兵"
	t1.Age = 45
	t1.School = "清华大学"
	fmt.Println(t1)
}
