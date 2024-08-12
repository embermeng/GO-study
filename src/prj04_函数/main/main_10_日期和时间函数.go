package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	fmt.Printf("%v, 对应的类型: %T\n", nowTime, nowTime)
	fmt.Printf("年: %v\n", nowTime.Year())
	fmt.Printf("月: %v\n", nowTime.Month())
	fmt.Printf("月: %v\n", int(nowTime.Month()))
	fmt.Printf("日: %v\n", nowTime.Day())
	fmt.Printf("时: %v\n", nowTime.Hour())
	fmt.Printf("分: %v\n", nowTime.Minute())
	fmt.Printf("秒: %v\n", nowTime.Second())
	fmt.Println("------------------------------------")
	// 日期格式化
	fmt.Printf("%d-%d-%d %d:%d:%d\n", nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute(), nowTime.Second())
	// Sprintf能得到字符串
	dateStr := fmt.Sprintf("%d-%d-%d %d:%d:%d\n", nowTime.Year(), nowTime.Month(), nowTime.Day(), nowTime.Hour(), nowTime.Minute(), nowTime.Second())

	fmt.Println(dateStr)
	// Format中参数的各个数字是固定的
	fmt.Println("Format指定格式：", nowTime.Format("2006-01-02 15:04:05"))
}
