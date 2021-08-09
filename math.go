package main

import (
	"fmt"
)

func main(){
	// 运算符的使用
	// int 除 int 结果为int
	fmt.Println(10/4)
	// 浮点数参与运算结果会保留浮点数
	fmt.Println(10.0/4)


	// 取模
	// 取模的公式 a/b = a-a/b*b
	fmt.Println(10%3)

	// ++ 和 -- 的使用
	var i int =10 
	i++  //等价 i = i+1
	fmt.Println("i=",i)

	i-- //等价 i= i-1
	fmt.Println("i=",i)
	


}
