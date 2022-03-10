package main

import (
	"fmt"
)

func main() {
	//根据用户指定月份，打印该月份所属的季节。3，4，5 春季 6，7，8 夏季 9，10，11 秋季 12，1，2 冬季
	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	switch month {
	case 3, 4, 5:
		fmt.Println("spring")
	case 6, 7, 8:
		fmt.Println("summer")
	case 9, 10, 11:
		fmt.Println("autumn")
	case 12, 1, 2:
		fmt.Println("winter")
	default:
		fmt.Println("输入有误...")
	}
}
