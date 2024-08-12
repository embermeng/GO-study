package dbutils

import "fmt"

func GetConn() { //首字母大写才能被其他包访问
	fmt.Println("执行了dbutils包下的getConn")
}
