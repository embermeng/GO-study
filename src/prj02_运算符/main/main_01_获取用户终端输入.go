package main

import "fmt"

func main() {
	//读取键盘输入
	//方式1：Scanln
	var age int
	var name string
	var score float32
	var isVIP bool

	fmt.Println("请录入学生的年龄：")
	//传入age的地址的目的：在Scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响了
	fmt.Scanln(&age)

	//方式2：Scanf
	fmt.Println("请录入学生的姓名，成绩，是否VIP：（使用空格隔开）")
	fmt.Scanf("%s%f%t", &name, &score, &isVIP)
	fmt.Printf("学生的年龄为：%v，姓名为：%v，成绩为：%v，是否VIP：%v\n", age, name, score, isVIP)
}
