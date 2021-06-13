package main

import "fmt"

func main() {
	/*
		字符
		byte(uint8)
		rune unicod
		字符 -> ascii码值 -> 二进制
	*/

	var charVar1 byte = '0'
	charVar2 := '贾'

	fmt.Printf("charVar1 = %d, charVar2=%d\n", charVar1, charVar2)

	// 字符参与运算, %d输出字符对应的ascii码值，
	var1 := 'a'
	fmt.Printf("和=%d， a的ascii码值=%d\n", 100+var1, var1)

	// %c,可以输出正常值
	fmt.Printf("原值为%c\n, 类型为%T", var1, var1)
}
