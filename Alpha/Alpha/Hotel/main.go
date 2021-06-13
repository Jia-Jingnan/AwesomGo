package main

import "fmt"

func main() {

	/**
	指针：指针主要是用来管理内存，存储的是另一个变量的内存地址
	*/

	var alpha int = 1
	fmt.Printf("alpha的值=%d, 内存地址为%v\n", alpha, &alpha)

}
