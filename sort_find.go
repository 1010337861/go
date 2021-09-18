package main

import (
	"fmt"
)
// 冒泡排序的实现
func BubbleSort(arr *[5] int){
	fmt.Println("排序前arr=",(*arr))
	temp := 0 
	for i:=0 ;i<len(*arr)-1 ; i++{
		for j :=0 ;j<len(*arr)-1-i ; j ++{
			if (*arr)[j] > (*arr)[j+1]{
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

// 查找 
// 顺序查找
// 二分查找
func BinaryFind(arr*[6] int ,leftIndex int,rightIndex int,findVal int){
	//判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex{
		fmt.Println("找不到")
		return
	}
	middleIndex := (leftIndex+rightIndex)/2

	if (*arr)[middleIndex]>findVal {
		//说明我们要查找的数，应该在 leftIndex -> middleIndex-1 之间
		BinaryFind(arr,middleIndex+1,rightIndex,findVal)
	}else{
		//找到了
		fmt.Printf("找到了，下标为 %v\n",middleIndex)
	}
}


func main() {
	// 定义数组
	arr := [5] int {13,10,-10,111,2}
	BubbleSort(&arr)
	fmt.Println("main arr=",arr) // 用指针 所以不用返回值？ 不用return 新的arr


	// 顺序查找
	names := [4] string {"白眉鹰王","金毛狮王","紫衫龙王","青翼蝠王"}
	heroname := ""
	fmt.Println("请输入要查中的名字:")
	fmt.Scanln(&heroname)

	// 顺序查找第一种方式
	for i:=0 ;i<len(names);i++{
		if names[i] == heroname{
			fmt.Printf("找到%v，下标%v \n",heroname,i)
			break
		}else if i == (len(names)-1) {
			fmt.Printf("没找到%v\n",heroname)
		}
		}
	
	// 顺序查找第二种方式(推荐)
	index  := -1 
	for i := 0 ;i<len(names);i++{
		if heroname==names[i]{
			index = i 
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v，下标%v \n",heroname,index)
	}else{
		fmt.Printf("没找到%v\n",heroname)
	}
	


	// 二分查找（递归）--针对有序数组
	arr_1 := [6] int{1,2,5,89,1000,1254}
	BinaryFind(&arr_1,0,len(arr)-1,89)
	// 二维数组
	var  arr_2  [4][6]int 

	arr_2[1][2]=1
	arr_2[2][1]=2
	arr_2[2][3]=3
	// 遍历二维数组，按照要求输出图形
	for i:= 0; i<4;i++{
		for j :=0 ;j<6; j++{
			fmt.Print(arr_2[i][j]," ")
		}
		fmt.Println()
	}

}