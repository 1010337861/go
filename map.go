package main

import (
	"fmt"
)

//map的基本介绍：map是key-value的数据结构，有称为字段或关联数组，类似于python的字典
// map 的声明  var map 变量名 map[keytype]valuetype
// key的类型：bool int float string ptr channel struct interface  array 都可以 ；通常为int 和string 注意 key 不能是slice map 和function 因为不能用 == 来比较判断
//valuetype类型： 和keytype 基本一致
func  main() {
	// map 声明举例
	var a map [string]string
	// var a map [string] int 
	// var a map [int] string 
	// var a map [string]map[string][string]
	// 注意：声明是不会分配内存的，初始化需要make，分配内存之后才能赋值和使用
	
	// 使用方法一 声明 ->  make-> 赋值
	a = make(map[string]string,10)
	a["no_1"] = "宋江"
	a["no_2"] = "吴用"
	a["no_1"] = "武松"
	a["no_3"] = "吴用"
	fmt.Println(a)
	// 对上述代码的说明：
	// map 在使用前一定要make 分配内存
	//map 的key有唯一性不能重复，如果重复了则以最后的key-value为准
	//map 是无序的 散列


	// map的使用
	//方法二 声明+make-> 赋值
	cities := make(map[string]string,3)
	cities["no_1"] = "北京"
	cities["no_2"] = "哈尔滨"
	cities["no_3"] = "杭州"
	fmt.Println(cities)
	
	//方法三 声明+初始化赋值
	hero := map[string]string{
		"hero1":"ironman",
		"hero2":"spiderman",
		"hero3":"greenman",
	}
	hero["hero4"] = "freshman"
	fmt.Println(hero)

	// map 的增删改查
	// 增加与修改 map["key"] = value
	// 删除  delect(map."key") 存在就删除 ，不存在不操作
	//查 findval = map[findkey]


	// map 遍历 for--range 的结构遍历
	for k,v := range cities{
		fmt.Printf("%v=%v\n",k,v)
	} 
	// 使用for-range 遍历一个复杂的map
	studentmap :=make(map[string]map[string]string)

	studentmap["stu01"] =make(map[string]string,3)
	studentmap["stu01"]["name"]="tom"
	studentmap["stu01"]["sex"]="男"
	studentmap["stu01"]["address"]="天安门"

	studentmap["stu02"] = make(map[string]string,3)//这句话不能少
	studentmap["stu02"]["name"]="marry"
	studentmap["stu02"]["sex"]="女"
	studentmap["stu02"]["address"]="地安门"
	for k1,v1 :=range studentmap{
		fmt.Printf("%v=%v\n",k1,v1)
		for k2,v2 := range v1{
			fmt.Printf("%v=%v\n",k2,v2)
		}
	}
	fmt.Println(studentmap)
	fmt.Printf("map的长度=%d\n",len(studentmap))  


	//map的长度 len(map)
	//map的切片:
	//1声明一个map切片 
	var monsters []map[string]string
	monsters = make([]map[string]string,2) // 准备放入两个键

	// 2 增加第一个妖怪信息
	if monsters[0] == nil {
		monsters[0]=make(map[string]string,3)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "360"
	}

	if monsters[1] ==nil {
		monsters[1]=make(map[string]string,3)
		monsters[1]["name"] = "玉兔"
		monsters[1]["age"] = "500"
	}
	// 这里我们需要使用切片的append函数，可以动态的增加monster
	//1 先定义一个monster信息
	newmonsters :=map[string]string{
		"name":"新的妖怪~xxx",
		"age":"150",
	}
	monsters = append(monsters,newmonsters)
	fmt.Println(monsters)

}