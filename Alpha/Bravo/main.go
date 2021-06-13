package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//go中的数据类型
	/*
		int8  占用1个字节		-2^7 ~ 2^7
		int64 占用2个字节
		int32 占用4个字节
		int64 占用8个字节
	*/
	// int类型定义与使用
	var alpha = 100 // int
	bravo := 200
	var charlie int32 //int32
	delta := 300      // int

	// int转换为int32
	charlie = int32(delta) //int 和 int32占用的空间大小不同
	var foxtrot int64 = 123456789
	// 输出类型
	fmt.Printf("alpha=%T,bravo=%T,charlie=%T", alpha, bravo, charlie)

	// 判断占用空间大小
	fmt.Println()
	fmt.Println(unsafe.Sizeof(foxtrot))
}
