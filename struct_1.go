package main

import (
	"fmt"
	"encoding/json"
)
// 面向对象编程
//说明：1、golang支持面向对象编程(oop),但是和传统的面向对象编程有区别，
//		2、golang中没有类(class),对应的为结构体(struct)，
//		3、golang中没有继承，方法重载，构造函数，析构函数，隐藏的this指针等方法，但存在相应特性


// 结构的声明方式 type structname struct{field1 type field2 type ....   }

//基本介绍
//从概念或叫法上看：结构体字段 = 属性 =field 
// 字段是结构体的一个组成部分，一般是基本数据类型、数组也可能是引用类型

//注意事项
//字段声明语法同变量 示例  字段名 字段类型
//字段的类型可以为：基本类型，数组或引用类型
//在创建一个结构体变量后，如果没有给字段赋值，都对应一个默认值，规则同前面
// 没分配内存空间是 slice map的默认值都是nil  
// 结构体默认值拷贝



// 演示
type Person struct {
	Name string 
	Age int
	Scores [5]float64
	ptr *int //指针
	slice []int //切片
	map1 map[string]string 
}

// 注意事项和细节 结构体的所有字段在内存中是连续的
type Point struct {
	x int 
	y int 
}

type Rect struct {
	leftup,
	rightDown Point
}

type Rect2 struct{
	leftup,
	rightDown *Point
}
// struct 的每个字段上可以写上一个tag 改tag 可以通过反射直接获取，常见的使用场景就是序列化与饭序列化
type Monster  struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}
// 方法的声明及调用 golang中的方法是作用在指定数据类型上的（即：和指定的数据类型绑定） 因此自定义类型都可以有方法
//
// type A struct {
// 	Num int
// }
// func (a A)test(){
// 	fmt.Println(a.Num)
// }
// 对上述语法的说明 func(a A)test() {}表示A结构体有一方法 方法名为test 
// (a A ) 体现test 方法是A类型绑定的

// 例： 
func (m Monster) test(){
	fmt.Println("test()name=",m.Name)
}

func main()  {
	//创建结构体变量和访问结构体字段
	//方式一 --直接声明
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil{
		fmt.Println("p1.prt is nil")
	}
	if p1.slice == nil{
		fmt.Println("p1.slice is nil")
	}
	if p1.map1 == nil{
		fmt.Println("p1.map1 is nil")
	}
	//使用slice 一定要make
	p1.slice = make([]int,10)
	p1.slice[2]=110

	//使用map 一定要make
	p1.map1 = make(map[string]string)
	p1.map1["key1"]="tom~"
	fmt.Println(p1)

	//方式二{}
	p2 := Person{Name:"mary",
				Age:20,
			}
	fmt.Println(p2)

	//方式三-&
	var p3 *Person = new(Person)
	p3.Name = "sitm"
	p3.Age =30
	fmt.Println(p3)

	//方式4 - {}
	var p4 *Person =&Person{}
	p4.Name = "xxxx"
	p4.Age =120
	fmt.Println(p4)
	
	r1 := Rect{Point{1,2},Point{3,4}}

	fmt.Printf("r1.leftup.x地址等于%p r1.leftup.y地址=%p\nr1.rightDown.x地址=%p r1.rightDown.y地址=%p\n",&r1.leftup.x,&r1.leftup.y,&r1.rightDown.x,&r1.rightDown.y)
	
	r2 := Rect2{&Point{10,20},&Point{30,40}}
	fmt.Printf("r2.leftup.x地址等于%p r2.leftup.y地址=%p\nr2.rightDown.x地址=%p r2.rightDown.y地址=%p\n",&r2.leftup.x,&r2.leftup.y,&r2.rightDown.x,&r2.rightDown.y)

	monster :=Monster{
		Name :"niu",
		Age : 100,
		Skill:"bang",
	}
	jsonStr,err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误",err)
	}
	fmt.Println("jsonstr",string(jsonStr))
	//调用方法
	monster.test()
}