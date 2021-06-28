package main

import (
	"fmt"
)

const (
	Engle = iota + 1
)

func main() {

	fmt.Println("请输入今天星期几?1-7")
	var userinput string
	fmt.Println(&userinput)

	switch userinput {
	case "Monday":
		fmt.Println("心情很糟糕")
	case "tuesday":
		fmt.Println("心情很糟糕！！！！")
	case "wednesday":
		fmt.Println("心情很糟糕！！！！！！")
	case "thursday":
		fmt.Println("心情很糟糕! ! ! ! ! ! ! ! ! ! ! ! !")
	case "friday":
		fmt.Println("心情很糟糕! ! ! ! ! ! ! ! ! ! ! ! ! ! ! !")
	case "saturday":
		fmt.Println("心情很糟糕! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! !")
	case "sunday":
		fmt.Println("心情很糟糕! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! ! !")
	default:

	}

}
