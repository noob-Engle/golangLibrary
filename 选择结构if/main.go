package main

import "fmt"

func main(){

	fmt.Print("请输入xxxxx:")
	var name string //声明一个变量字符串类型
	fmt.Scan(&name)//接收用户输入存入name变量所在的地址
	//单分支判断
	if name == "shabi"{ //判断用户输入如果name等于shabi就执行我是shabi，如果不是shabi的话直接打印我不是shabi
		fmt.Println("我是shabi")
	}
	fmt.Println("我不是shabi")
}