package main

import "fmt"

func danfen(){

	fmt.Print("请输入xxxxx:")
	var name string //声明一个变量字符串类型
	fmt.Scan(&name)//接收用户输入存入name变量所在的地址
	//单分支判断
	if name == "shabi"{ //判断用户输入如果name等于shabi就执行我是shabi，如果不是shabi的话直接打印我不是shabi
		//也就是判断字符串是否相等
		fmt.Println("我是shabi")

		return //结束当前函数的执行
	}
	fmt.Println("我不是shabi")
}

func shuangfenzhi(){
	fmt.Print("请输入你喜欢的人:")
	var name string
	fmt.Scan(&name)
	if name == "xiaojie" {
		fmt.Println("你不是我要的人")

	}else{
		fmt.Println("你是我喜欢的人")
		return//结束当前函数的执行就不打印下面的人生完蛋了

	}
	
	fmt.Println("人生完蛋！！！！！！！！！！！！！")
}


//多分枝结构
func main(){
	fmt.Println("请输入你的喜欢的女明星:")
	var name string
	fmt.Scan(&name)
	if name == "xiaohong"{
		fmt.Println("滚吧")
	}else if name == "杨幂"{
		fmt.Println("jio臭")
	}else if name == "女帝"{
		fmt.Println("i love")
		return
	}
	fmt.Println("你没有找到你喜欢的明星")
}