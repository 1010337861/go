package main

import (
	"fmt"
)

// 数组的定义 ： var 数组名 [数组大小] 数据类型  ex： var a [5] int,赋初值 a[0] =1 a[1] = 30...  可以通过数据名获取数组地址&intarr[i]
func main() {
	var intarr [3] int 
	intarr[0] = 1 
	intarr[1] = 3
	intarr[2] = 5
	fmt.Println("intarr=", intarr)

	// 四种初始化数组的方式：
	var numarr01 [3] int= [3]int{1,2,3}
	fmt.Println("numarr01=", numarr01)

	var numarr02 = [3] int {5,6,7}
	fmt.Println("numarr02=", numarr02)

	// 这里的[...]是规定写法
	var numarr03 = [...] int {5,6,7}
	fmt.Println("numarr03=", numarr03)

	// 这里的[...]是规定写法
	var numarr04 = [...] int {1:5,0:6,2:7}
	fmt.Println("numarr04=", numarr04)

	// 类型推导
	strarr05 := [...]string{1:"tom",0:"jack",2:"mary"}
	fmt.Println("strarr05=", strarr05)

	//数组的变量 for-range 结构遍历   for index,value := rsange array01 {}

	for i ,v := range strarr05 {
		fmt.Printf("i=%v, v=%v\n",i,v)
		fmt.Printf("strarr01[%d]=%v\n",i,strarr05[i])
		fmt.Printf("v的值=%v\n", v)
	}

	//数组的使用细节  数组是多个相同类型数据的组合一旦声明了/定义了，其长度也是固定的，不能发生动态变化  即int类型的数组 不能存放float数据，数组创建后 没赋值 会有默认值
	var mychar [26] byte
	for i:=0;i<26;i++ {
		mychar[i] = 'A' + byte(i)
	}

	for i:=0;i<26;i++ {
		fmt.Printf("%c",mychar[i])
	}
	fmt.Println(" ")
	// 求数组中的最大值
	var maxarr [5] int = [...] int {10,50,12,100,3}
	maxval := maxarr[0]
	maxindex := 0
	for i:=0;i<len(maxarr);i++{
		if maxval<maxarr[i]{
			maxval = maxarr[i]
			maxindex = i
		}
	}
	fmt.Printf("最大值=%v,最大索引=%v\n",maxval,maxindex)

	// 切片 slice 需要一个数组但是个数不确定 --> 可以使用切片 ：切片是数组的一个引用，变量切片和访问切片和求切片长度都与数组类似
	// 切片长度可以变化，因此切片是一个可动态变化的数组  基本语法 var 切片名 [] 类型  var a [] int 
	// c 从底层来说 数据是一个引用类型 其实就是一个数据结构（struct 结构体）
	// type slice struct{
	// 	ptr *[2]int
	// 	len int
	// 	cap int
	// }

	// 切片的使用 方式1：定义一个切片  然后让切片去引用一个创建好的数组，
	slice01 := maxarr[1:3]
	fmt.Println("arr=",maxarr)
	fmt.Println("slice=",slice01)
	fmt.Println("slice len =",len(slice01))
	fmt.Println("slice cap =",cap(slice01))
	// 方式2：通过make来穿件切片 var 切片名[]type=make([]type,len,|cap|) type:数据类型，len：大小，cap：指定切片容量【可选，如果分配了 cap必须≥len】
	slice02 := make([] float64,5)
	fmt.Println("slice=",slice02)
	fmt.Println("slice len =",len(slice02))
	fmt.Println("slice cap =",cap(slice02))
	// 方式3：定义一个切片 直接就指定具体数组
	slice03  := []string{"tim","jsak"}
	fmt.Println("slice=",slice03)
	fmt.Println("slice len =",len(slice03))
	fmt.Println("slice cap =",cap(slice03))


	// 切片的遍历 1：for 遍历， 2：for range 遍历
	for i :=0 ;i<len(slice02);i++{
		fmt.Println("%v,%v",i,slice02[i])
	}

	for i,v :=range slice01 {
		fmt.Printf("i=%v,v=%v\n",i,v)
	}
	//  切片的使用注意事项和细节讨论
	//1 切片初始化 var slice = arr[startindex:endindex] 不包含arr[endindex]
	//2 切片初始化时 仍不能越界 范围在[0~len(arr)]之间 var slice = arr[0:endindex] 可简写 var slice = arr[:end]
	//3 cap是一个内置函数，用于统计切片容量
	//4 切片定义完，还不能使用，因为本身是一个空，需要引用到一个数组
	//5 切片可以继续切片
	//6 用append内置函数可以对切片进行动态追加
	//7 切片的copy 内置函数copy
	slice05 := make([] float64 ,4)
	copy(slice05,slice02)
	fmt.Println("slice05=",slice05)
}