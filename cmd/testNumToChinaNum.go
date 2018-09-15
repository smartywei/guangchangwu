package main

import (
	"fmt"
	"my_book/getResource"
)

func main()  {
	num1 := 100011020

	res1,_ := getResource.GetChinaNum(num1)

	fmt.Println(res1)
}