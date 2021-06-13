package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//浮点数定义及使用
	/**
	float32 4个字节 相当于float
	float64 8个字节 相当于double
	*/
	var alpha = 3.1415926

	fmt.Printf("alpha=%T", alpha)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(alpha))

	var floatV5 float32 = 3.14
	var floatV6 float64 = 3.14

	floatV6 = float64(floatV5)
	floatV6 = floatV6

	// 复数
	// 实数 + 虚数i
	// complex64 32为实数和32为虚数组成
	// complex128, 64位实数和64为虚数组成
	var complexVar complex64
	complexVar = 3.14 + 12i

	complexVar2 := complex(3.14, 2)

	fmt.Printf("complexVar的类型=%T，值=%v\n", complexVar, complexVar)
	fmt.Printf("complexVar2的类型=%T,值=%v\n", complexVar2, complexVar2)

	// 打印出实数和虚数部分
	fmt.Println(real(complexVar), imag(complexVar))

}
