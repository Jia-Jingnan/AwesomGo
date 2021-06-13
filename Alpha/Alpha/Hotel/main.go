package main

import "fmt"

func main() {

	/**
	指针：指针主要是用来管理内存，存储的是另一个变量的内存地址
	*/

	var alpha int = 1
	fmt.Printf("alpha的值=%d, 内存地址为%v\n", alpha, &alpha)

	var bravo *int = &alpha
	fmt.Printf("bravo的值为%v, 地址%v\n", bravo, &bravo)

	fmt.Println(alpha)
	*bravo = 200
	fmt.Println(alpha)

	var charlie int = 123
	var delta *int = &charlie
	fmt.Println(*delta, charlie)

	/**
	值类型：整型，浮点型，bool,array,string,栈中分配
	引用类型:指针，slice，map，chan，interface，堆中分配
	*/

}
