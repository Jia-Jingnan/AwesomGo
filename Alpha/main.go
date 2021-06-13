package main

import "fmt"

// 全局变量
var (
	golf  int
	hotel []int
	india interface{}
)

func main() {
	fmt.Println("Hello, Go")

	// 变量定义方式1
	var aplpha int
	// 变量定义方式2
	bravo := "我是Bravo"
	// 变量定义方式3
	var charlie = "我是Charlie"
	// 变量定义方式4
	var delta, echo, foxtrot = "Delat", "Echo", 4
	fmt.Println(aplpha, bravo, charlie, delta, echo, foxtrot)

	// 交换两个变量的值
	var juliet = "我是Juliet"
	var kilo = "我是Kilo"
	juliet, kilo = kilo, juliet
	fmt.Println("julet:"+juliet, "kilo:"+kilo)
}
