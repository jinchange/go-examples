package basic

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// SellerAdaptor /* 1. 包 2. 函数 3. 结构体 4. 方法 5. 接口 6.type */
type SellerAdaptor interface {
	HandleRequest(requestStr string) string
	HandleResponse(responseStr string) string
}

type OpenSellerAdaptor struct {
}

func (sellerAdaptor *OpenSellerAdaptor) HandleRequest(requestStr string) string {
	return requestStr + " conv"
}

func (sellerAdaptor *OpenSellerAdaptor) HandleResponse(responseStr string) string {
	return responseStr + " conv"
}

// Student 定义结构体
type Student struct {
	name    string
	age     int
	address Address // 指定字段名为address，也可不写默认为对象名
}

type Address struct {
	city    string
	country string
}

func structDemo() {
	// 初始化结构体
	student := Student{
		name: "ZJ",
		age:  18,
		address: Address{
			city:    "SZ",
			country: "CH",
		},
	}
	fmt.Println(student)

	// 空结构体没有字段，不占用内存空间，可以通过unsafe.SizeOf函数来计算占用的字节大小
	type Empty struct {
	}
	fmt.Println(unsafe.Sizeof(Empty{}))
}

type Man struct {
	Name string
	Age  int
}

func (man *Man) sayHi() {
	fmt.Println("Hi,My name is " + man.Name + " my age is " + strconv.Itoa(man.Age))
}

func (man *Man) updateMan(name string, age int) {
	man.Age = age
	man.Name = name
}

func methodDemo() {
	// new 返回指针
	man := new(Man)
	man.updateMan("WW", 30)
	man.sayHi()
}

type OS int8

const (
	Android OS = 1
	IOS     OS = 1
)

// 可用于定义枚举
func typeDemo() {
	fmt.Println(Android)
	fmt.Println(IOS)
}

func TestInterface() {
	var sellerAdaptor SellerAdaptor = new(OpenSellerAdaptor)
	fmt.Println(reflect.TypeOf(sellerAdaptor))
	request := sellerAdaptor.HandleRequest("requestString")
	fmt.Println(request)
	response := sellerAdaptor.HandleResponse("requestString")
	fmt.Println(response)
}

// TODO Any ?
