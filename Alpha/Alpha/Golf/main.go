package main

import "fmt"

func main() {
	/**
	bool:布尔类型数据
	true false
	*/
	var boolVar1 bool
	boolVar1 = true
	boolVar2 := (true == false)

	fmt.Println(boolVar1, boolVar2)

	if 3 == 4 {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

	fmt.Println(!true, !!false, !false)

	// 不支持 ===，判断相等为 ==
}
