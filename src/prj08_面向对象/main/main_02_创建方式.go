package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	// 创建结构体实例的方式
	var s1 Student = Student{Age: 18, Name: "ember", School: "黑龙江大学"}
	fmt.Println(s1)

	var s2 Student = Student{"马士兵", 18, "黑龙江大学"}
	fmt.Println(s2)

	// s3是指针，应该给这个地址指向的对象的字段赋值
	var s3 *Student = new(Student)
	(*s3).Name = "张三"
	(*s3).Age = 45
	// 为了符合程序员的编程习惯，go提供了简化的赋值方式
	// go编译器底层对s3.School 转化为 (*s3).School =
	s3.School = "清华大学"
	fmt.Println(*s3)

	var s4 *Student = &Student{} // 当然也可以在{}里赋值
	(*s4).Name = "李四"
	(*s4).Age = 20
	s4.School = "蓝翔技术学院"
	fmt.Println(*s4)
}
