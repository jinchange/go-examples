package basic

import (
	"encoding/json"
	"fmt"
)

func MapDemo() {
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
			delete(m, "Tom")
		}
	}
	fmt.Println("length:", len(m))
}

// Set set --------------
type Set struct {
	m map[string]struct{}
}

func NewSet() *Set {
	set := new(Set)
	set.m = map[string]struct{}{}
	return set
}

func (set *Set) Add(val string) {
	set.m[val] = struct{}{}
}

func (set *Set) Contains(val string) bool {
	_, e := set.m[val]
	return e
}

func (set *Set) Remove(val string) {
	delete(set.m, val)
}

func (set *Set) GetAll() []string {
	slice := make([]string, len(set.m))
	index := 0
	for s := range set.m {
		slice[index] = s
		index++
	}
	return slice
}

func (set *Set) ToString() {
	fmt.Println(set.GetAll())
}

// -------------

func SetDemo() {
	// map key唯一充当set使用
	set := make(map[string]struct{})
	set["A"] = struct{}{}
	set["A"] = struct{}{}
	set["B"] = struct{}{}

	for s := range set {
		fmt.Println(s)
	}

	if _, exist := set["C"]; exist {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
}

func JsonDemo() {
	type Person struct {
		Name     string   `json:"name"`
		Age      int      `json:"age"`
		Language []string `json:"language"`
	}

	person := Person{
		Name:     "ZJ",
		Age:      18,
		Language: []string{"EN", "ZH"},
	}
	// 序列化
	bytes, err := json.Marshal(person)
	if err == nil {
		fmt.Println(string(bytes))
	} else {
		fmt.Println(err)
	}
	if indent, err := json.MarshalIndent(person, "", "\t"); err == nil {
		fmt.Println(string(indent))
	}

	// 反序列化
	p := Person{}
	if err := json.Unmarshal(bytes, &p); err == nil {
		fmt.Println(p.Age)
	}
}

func SliceDemo() {
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
