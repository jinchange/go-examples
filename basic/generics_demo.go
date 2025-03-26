package basic

import (
	"errors"
	"fmt"
)

// Number 范型类型
type Number interface {
	int | float64
}

func SumAll[T Number](nums []T) T {
	var sum T
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

// Sum 范型方法
func Sum[T int | float64](a, b T) T {
	return a + b
}

// MySlice 范型切片
type MySlice[T int | string] []T

// MyMap 范型Map
type MyMap[K comparable, V int | string] map[K]V

func TestGenerics() {
	e := errors.New("this is a error")
	fmt.Println(e)

	// 指定类型
	sum := Sum[int](1, 2)
	fmt.Println(sum)
	// 类型推断
	sum2 := Sum(3.14, 4.68)
	fmt.Println(sum2)

	mySlice := MySlice[int]{1, 2, 3}
	fmt.Println(mySlice)
	mySlice2 := MySlice[string]{"A", "b", "c"}
	fmt.Println(mySlice2)

	myMap := MyMap[string, int]{
		"a": 1,
	}
	fmt.Println(myMap)

	myMap2 := make(MyMap[int, string])
	myMap2[1] = "a"
	fmt.Println(myMap2)

	nums := []float64{1.2, 2.6, 3, 4, 5, 6}
	fmt.Println(SumAll(nums))
}
