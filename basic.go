package main

import (
    "fmt"
    "unsafe"
    "strconv"
) 

func main(){ 
    var n1 int = 1
    fmt.Println("n1=",n1,"\n")
    

    var n2 int64 = 257
    fmt.Printf("n2的数据类型%T n2占用字节数 %d\n",n2,unsafe.Sizeof(n2))


    // golang中的字符类型
    var c1  string = "a"
    var c2  string = "1"
    fmt.Println("c1 =",c1,"c2 =",c2," \n" )

    var n3 int = 22269
    fmt.Printf("n3=%c\n",n3)

    var b1 bool = true 
    fmt.Println("b1=",b1,"\n")

    // golang中的字符串类型
    var  c4  string = "哈哈"
    c4 ="呜呜"
    fmt.Println("c4=",c4,"\n")
    
    // 基本数据类型的转换
    //字符串与其他类型相互转换    
    var str string 
    //使用fmt.Sprintf 进行类型转换
    str = fmt.Sprintf("%d",n1)
    fmt.Println("str=",str)

    //字符串转bool

    str = "true"
    b1,_=strconv.ParseBool("str")
    fmt.Printf("b1 type %T,b1的值%v\n",b1,b1)

    // 字符串转整型
    str = "12345"
    n2 ,_=strconv.ParseInt(str,10,0)
    fmt.Printf("n2 type %T,n2的值%v\n",n2,n2)


}