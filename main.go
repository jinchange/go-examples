package main

import (
	"fmt"
)

func main() {
	mapDemo()
}

func sliceDemo() {
	nums := [...]int{1, 2, 3}
	fmt.Println(nums)
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(nums2)
	ints := append(nums2, 5)
	fmt.Println(ints)
	fmt.Println(ints[:3])
	fmt.Println(cap(ints))
	fmt.Println(len(ints))

	//a := make([]int, 10)
	//fmt.Println(a)

	strs := make([]string, 10)
	fmt.Println(len(strs), cap(strs))
	strs[0] = "abc"
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
	//for i := range len(strs) {
	//	fmt.Println(strs[i])
	//}
	//for i, v := range strs {
	//	fmt.Println(i, v)
	//}
	// slice 和 ArrayList 的不同，先分配内存后赋值使用，先创建引用后自动扩容
}

func mapDemo() {
	// 定义
	//var m map[string]int
	// m1 := make(map[string]int, 10)
	// TODO map函数发生了什么？
	m := make(map[string]int, 10)
	fmt.Println(m)
	// 初始化
	m = map[string]int{
		"Tom": 18,
		"Bob": 25,
	}
	//赋值
	m["Jenny"] = 30
	// 取值
	age, exist := m["Tom"]
	fmt.Println(age, exist)
	fmt.Println("length:", len(m))
	// 遍历, TODO 为什么可以在遍历的时候赋值？
	for k, v := range m {
		fmt.Println(k, v)
		if k == "Jenny" {
			m["Team"] = 45
		}
	}
	fmt.Println("length:", len(m))
}
