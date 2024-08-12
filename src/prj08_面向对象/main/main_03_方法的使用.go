package main

import "fmt"

type Car struct {
	Name string
}

// Drive 给Car结构体绑定方法，对象c 是值传递
func (c Car) Drive(meter int) {
	fmt.Println(c.Name, "跑了", meter, "米")
}

// Run 传递结构体指针
func (c *Car) Run(meter int) {
	fmt.Println(c.Name, "又跑了", meter, "米")
}

func main() {
	var car1 Car
	car1.Name = "兰博基尼"

	car1.Drive(500)
	(&car1).Drive(500) // 虽然用指针类型调用，但传递还是按照值传递的形式

	// 下面两个等价
	car1.Run(1000)
	(&car1).Drive(1000)
	// 综上，对于方法，用对象或指向对象的地址调用都可以
}
