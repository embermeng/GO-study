package main

import (
	"errors"
	"fmt"
)

func main() {
	err := testErr02(10, 0)
	if err != nil {
		fmt.Println("自定义错误：", err)
		// panic停止当前程序的执行
		panic(err)
	}
	fmt.Println("继续执行后续逻辑")
}

func testErr02(num1 int, num2 int) (err error) {
	if num2 == 0 {
		// 抛出自定义错误
		return errors.New("除数不能为0哦！")
	} else {
		res := num1 / num2
		fmt.Println(res)
		// 如果没有错误，返回nil
		return nil
	}
}
