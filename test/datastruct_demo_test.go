package test

import (
	"go-examples/basic"
	"testing"
)

func TestSliceDemo(t *testing.T) {
	basic.SliceDemo()
}

func TestMapDemo(t *testing.T) {
	basic.MapDemo()
}

func TestSetDemo(t *testing.T) {
	set := basic.NewSet()
	set.Add("A")
	set.Add("A")
	set.Add("B")
	set.Add("C")
	set.Remove("A")
	println(set.Contains("A"))
	for _, v := range set.GetAll() {
		println(v)
	}
	set.ToString()
}

func TestJsonDemo(t *testing.T) {
	basic.JsonDemo()
}

func TestConcurrent(t *testing.T) {
	basic.GoroutineDemo()
}
