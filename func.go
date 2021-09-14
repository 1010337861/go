package main

import (
	"fmt"
	"strconv"
	"time"
)
// 声明变量及变量类型，及返回值类型
func Cal(num1 int ,num2 int,operator string) int {
	result := 0
	switch operator  {
		case "+":
			result = num1 + num2
		case "*" :
			result = num1 * num2
		case "-" :
			result = num1 - num2
		case "/" :
			result = num1 / num2
	}
	return result
} 

// 可以用多个返回值
func getSunAdd(num1 int ,num2 int)(int,int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum,sub
}

// 递归的艰难应用：斐波那契数列
func fbn(n int) int {
	if (n==1|| n==2){
		return 1 
	}else{
		return fbn(n-1)+fbn(n-2)
	}
}

func test02(n1 int){
	n1 = n1 + 10 
	fmt.Println("test02() n1= ",n1)
}

func test03(n1 *int){
	*n1 = *n1 + 10 
	fmt.Println("test03() n1= ",n1)
}
type myfunctype func(int,int) (int,int)
func myfun(funvar myfunctype, num1 int,num2 int) (int,int) {
	return funvar(num1,num2)
}

//go支持可变参数 支持0 到多个参数 //args是silce 切片 -- python中的list ，通过args[index] 可以访问到各个值
func sum_(n1 int ,args... int) int {
	sum :=n1
	// 遍历args 
	for i :=0;i<len(args);i++ {
		sum +=args[i] 
	}
	return sum
}

// init 函数，通常可以在init函数完成初始化的工作
// 如果一个文件同事包含全局变量定义，init函数和main函数 则执行流程为 全局变量 -> init函数 -> main函数

func init() {
	fmt.Println("init()....")
}

// 匿名函数 没有名字的函数 常用的使用方法： 1 直接调用；2 赋值给变量
// 全局匿名函数
var (
	lambda = func(n1 int,n2 int) int{
		return n1+n2
	}
)

// 闭包： 一个函数和与其相关引用环境组合的一个整体(实体)
func addupper() func (int) int{
		n := 10 
		return func (x int) int{
			n = n+x
			return n
		}
	}


// 函数defer - 延时机制 
// 按照先入后出的方式进出栈
func defef_ex(n1 int ,n2 int) int {
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("ok1 n2=",n2)
	res := n1 + n2 
	fmt.Println("res", res)
	return res
}


// 错误处理函数
func test(){
	defer func() {
		err := recover() //recover() 内置函数 可以捕获到异常
		if err != nil{
			fmt.Println("err=",err)
			fmt.Println("终止程序")
		}
	}()
	num1 := 10 
	num2 := 0 
	res := num1 /num2 
	fmt.Println("res=",res)
}

func main(){
	res := Cal(2,5,"+")
	fmt.Println("res =",res)

	//函数内变量为局部变量，只在函数中起作用 
	sum, sub := getSunAdd(6,3)
	fmt.Println("sum = ",sum,"sub =",sub)

	// 函数的递归
	fbn := fbn(5)
	fmt.Println("fbn = ",fbn)

	// 基本数据类型和数组默认都是值传递，即进行值拷贝。在函数内修改不会影响原来的值
	num := 20 
	test02(num)
	fmt.Println("main() num= ",num)

	// 如果希望函数内变量能够修改函数外的变量(值默认以值传递的方式数据类型)，可以传入变量的地址&，函数内以指针的方式操作变量
	test03(&num)
	fmt.Println("main() num= ",num)

	// Go 函数不支持函数重载，即相同函数名称 重复定义
	// Go 中函数也是一种数据类型，可以赋值给一个变量，变量类型就是一个函数类型的变量了，通过变量可以对函数进行调用
	// 函数既然是一种数据类型 在Go中 函数可以作为形参 并被调用 

	// Go支持自定义数据类型 但自定义数据类型不等于 原生数据类型 及自定义的 int64 不等于 Go中的int64 
	type myint int 

	var num1_  myint 
	var num2_  int
	num1_ = 40 
	num2_ = int(num1_)
	fmt.Println("num1_ = ",num1_,"num2_",num2_)

	res2,_:= myfun(getSunAdd,6,2)
	fmt.Println("res2=",res2) 
	
	res3 := sum_(13,0,1,90,10,100)
	fmt.Println("res3=",res3) 

	// 匿名函数
	res4 := func(n1 int,n2 int) int {
			return n1 + n2 	
	}(10,110)
	fmt.Println("res4=",res4)

	lambda_func := func(n1 int,n2 int) int {
			return n1 + n2 	
	}
	fmt.Println("res4=",lambda_func(12,21))

	res5 := lambda(1,3)
	fmt.Println("res5=",res5)

	// 闭包：应用
	c_pkg :=addupper()
	fmt.Println(c_pkg(1))
	fmt.Println(c_pkg(4))
	fmt.Println(c_pkg(7))
	
	// defer 应用
	def_ex := defef_ex(1,3)
	fmt.Println("def_ex",def_ex)

	// 字符串常用的系统函数
	str := "hello word"
	fmt.Println("len of str = ",len(str))

	// 字符串遍历 s := []rune(str)
	s :=[]rune(str)
	for i:=0 ;i< len(s);i++{
		fmt.Printf("字符=%c\n",s[i])
	}

	// 字符串转整数
	n,err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转化错误",err)
	}else{
		fmt.Println("转成的结果是",n)
	}

	//整数转字符串
	str = strconv.Itoa(1234)
	fmt.Printf("str = %v str=%T\n",str,str)

	// 查找子串是否在指定字符串中：strings.Contains("seafood","foo")//true 
	// 统计一个字符串有几个指定子串 strings.Count("ceheese","e") //4
	// 区分大小写比较 strings.Equalfold("abc","aBC") //True
	//strings.Index("abc_aa","a") //0 返回子串一次出现的索引值 没有返回-1  LastIndex 返回最后一次出现的索引值，没有返回-1

	num1_new := 100 
	fmt.Printf("num1_new的类型%T，num1_new的值=%v,num1_newd的地址%v\n", num1_new,num1_new,&num1_new)

	num2_new := new(int)
	*num2_new = 100
	fmt.Printf("num2_new的类型%T，num2_new的值=%v,num2_newd的地址%v，num2_new 这个指针指向的值=%v\n", num2_new,num2_new,&num2_new,*num2_new)

	// 错误处理
	test()
	for {
		fmt.Println("main() 下面的代码....")
		time.Sleep(time.Second)
	}


}

// 函数参数的传递方式
	//1、值传递 2、引用传递  一般来说 地址拷贝的效率高，因为数据量小；
	// 值类型 和引用类型
	// 值类型（默认值传递）： 基本数据类型，数组和结构体struct   引用类型（默认引用传递）： 指针，slice切片，map，管道chan ，interface；

// 变量作用域
// 函数内部声明/定义的变量叫局部变量，作用域仅限于函数内部
// 函数外部声明/定义的变量叫全局变量，作用域在整个包，如果首字母大写，则作用域在整个程序有效
// 如果变量是在一个代码块，比如for/if 那么这个变量的作用域就在改代码块中


// 内置函数 len 用于求长度， new 用于分配内存 make 用来分配内存 主要分配引用类型
// 错误处理函数 处理方式：panci defer panic recover  python： try except 
// 自定义错误类型 使用errors.New 和panic 内置函数  1： 使用errors.New("错误说明")，会返回一个error类型的值，表示一个错误，2 panic内置函数接受一个interface{}类型的值(任何值) 可以作为error类型的变量 
