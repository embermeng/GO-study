package main

import (
	"GO-study/src/mypkg/dbutils"
	"fmt"
)

func main() {
	dbutils.GetConn()
	fmt.Println("hh")
}
