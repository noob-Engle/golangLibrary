package main

import "fmt"

func main(){
	fmt.Print("请输入xxxxx:")
	var name string //声明一个变量字符串类型
	fmt.Scan(&name)//接收用户输入存入name变量所在的地址
	
}