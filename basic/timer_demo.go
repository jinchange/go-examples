package basic

import (
	"fmt"
	"time"
)

func TimerDemo() {
	afterFunc := time.AfterFunc(3*time.Second, func() {
		fmt.Println("afterFunc")
		fmt.Println(time.Now().Format(time.DateTime))
	})
	time.Sleep(4 * time.Second)
	afterFunc.Stop()

	//d := 3 * time.Second
	//timer := time.NewTimer(d)
	//defer func() {
	//	timer.Stop()
	//	fmt.Println("stop timer!")
	//}()
	//
	//for {
	//	timer.Reset(d)
	//	select {
	//	case t := <-timer.C:
	//		fmt.Println(t.Format(time.DateTime))
	//	}
	//}

	//fmt.Println(time.Now().Format(time.DateTime))
	//select {
	//case t := <-time.After(3 * time.Second):
	//	fmt.Println(t.Format(time.DateTime))
	//}

	//start := time.Now()
	//time.Sleep(1 * time.Second)
	//t := time.Now()
	//t.
	//elapsed := t.Sub(start)
	//fmt.Println(elapsed)
}
