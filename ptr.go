package main

import (

	"fmt"
)


func main(){
	// 指针
	var i int = 10 
	//  i 的地址 通过 &i可得获得
	fmt.Println("i的地址=",&i,"\n")

	// 下面的var ptr *int =&i 
	//1.ptr 是一个指针变量
	//2.prt 的类型是*int
	//3.ptr 本身的值&i
	var ptr *int =&i 
	fmt.Printf("ptr=%v\n",ptr)
	fmt.Println("ptr的地址:",&ptr,"\n")

	//获取指针类型所指向的值，使用：* 
	fmt.Printf("ptr 指向的值=%v\n",*ptr)


	// 获取int变量的地址 并显示终端
	// 将nun地址给指针ptr 并通过ptr修改num的值

	var num int = 100
	fmt.Printf("num的内存地址%v\n",&num)

	var num_ptr *int = &num 
	*num_ptr = 10
	fmt.Printf("num的值:%v",num)

	

	

}