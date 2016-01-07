// Go 之 字符串格式化
//
// Copyright (c) 2015 - Batu <1235355@qq.com>
//
package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main(){
	// 格式化整型，使用`%d`是一种
	// 标准的以十进制来输出整型的方式
	// 有符号十进制整数(int)（%ld、%Ld：长整型数据(long),%hd：输出短整形。）
	fmt.Println("=====%d,输出十进制====")
	fmt.Printf("%d\n", 110)

	// 输出整型的二进制表示方式
	fmt.Println("=====%b,输出二进制====")
	fmt.Printf("%b\n", 110)

	// 输出整型数值所对应的字符(char):一个字节，占8位
	// 可参考 ASCII
	fmt.Println("=====%c,输出一个值的字符(char)====")
	fmt.Printf("%c\n",97)

	// 输出一个值的十六进制,每个字符串的字节用两个字符输出
	fmt.Println("=====%x,输出一个值的十六进制,每个字符串的字节用两个字符输出====")
	fmt.Printf("0x%x\n", 10)
	fmt.Printf("%x\n", "abc")

	// 输出浮点型数值
	fmt.Println("=====%f,输出浮点型数值====")
	fmt.Printf("%f\n", 27.89)

	// 输出基本的字符串
	fmt.Println("=====%s,输出基本字符串====")
	fmt.Printf("%s-%s-%s\n","I","am","batu")

	// 输出带双引号的字符串
	fmt.Println("=====%q,输出带双引号的字符串====")
	fmt.Printf("%q\n","string")



	// Go提供了几种打印格式，用来格式化一般的Go值
	p := point{1, 2}

	fmt.Println("=====%p,输出一个指针的值====")
	fmt.Printf("%p\n", &p)
	fmt.Println("=====%v,输出结构体的对象值====")
	fmt.Printf("%v\n", p)
	// 如果所格式化的值是一个结构体对象，那么`%+v`的格式化输出
	fmt.Println("=====%+v,输出结构体的成员名称和值====")
	fmt.Printf("%+v\n", p)
	fmt.Println("=====%#v,输出一个值的Go语法表示方式====")
	fmt.Printf("%#v\n",p)
	fmt.Println("=====%T,输出一个值的数据类型====")
	fmt.Printf("%T\n",p)

	// 当输出数字的时候，经常需要去控制输出的宽度和精度。
	// 可以使用一个位于%后面的数字来控制输出的宽度，默认情况下输出是右对齐的，左边加上空格
	fmt.Println("=====控制输出的宽度和精度====")
	fmt.Printf("|%5d|%5d|\n", 12, 345)
	fmt.Println("=====输出宽度,同时指定浮点数====")
	fmt.Printf("|%5.2f|%5.2f|\n", 1.2, 3.45)
	fmt.Println("=====左对齐====")
	fmt.Printf("|%-5.2f|%-5.2f|\n", 1.2, 3.45)

}