package main

import "fmt"

func main() {
	/**
	字符串:string
	*/

	var stringVar1 string
	stringVar1 = "Hello,贾敬楠\n"

	fmt.Println(stringVar1)

	fmt.Printf("类型%T\n", stringVar1)

	var stringVar2 = "hello, imooc,贾敬楠"
	stringVar2Len := len(stringVar2)

	for index := 0; index < stringVar2Len; index++ {
		fmt.Printf("%s--字节=%d，值=%c\n", stringVar2, stringVar2[index], stringVar2[index])
	}

	// 另外一种方式的for循环
	for _, value := range stringVar2 {
		fmt.Printf("%s--字节=%d，值=%c, 类型=%T\n", stringVar2, value, value, value)
	}

}
