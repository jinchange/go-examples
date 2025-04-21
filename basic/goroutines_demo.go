package basic

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func createGoroutine() {
	go fmt.Println("run")
	go func() {

	}()
	go run()
}

func run() {
	fmt.Println("run")
}

func channelDemo() {
	// create channel
	// channel 就是阻塞队列
	// 有缓冲，无缓冲
	intch := make(chan int, 10)
	defer close(intch)
	// write
	intch <- 100
	// read
	val, ok := <-intch
	fmt.Println(val, ok)
}

func selectDemo() {
	// 监视管道
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	defer func() {
		close(chA)
		close(chB)
		close(chC)
	}()
	go Send(chA)
	go Send(chB)
	go Send(chC)
	for {
		select {
		case n, ok := <-chA:
			fmt.Println("A", n, ok)
			break
		case n, ok := <-chB:
			fmt.Println("B", n, ok)
			break
		case n, ok := <-chC:
			fmt.Println("B", n, ok)
			break
		default:
			fmt.Println("no")
		}
	}
}

func Send(ch chan<- int) {
	i := 0
	for {
		time.Sleep(time.Millisecond)
		ch <- i
		i++
	}
}

func waitGroupDemo() {
	var buyerCount sync.WaitGroup
	buyerResponseList := make(chan int, 10)
	buyerCount.Add(10)
	for i := range 10 {
		go func() {
			buyerResponseList <- i
			defer buyerCount.Done()
		}()
	}
	buyerCount.Wait()
	// 关闭后就不会阻塞了
	close(buyerResponseList)
	for i := range buyerResponseList {
		fmt.Println(i)
	}
}

func contextDemo() {
	//f := func(ctx context.Context, k string) {
	//	if v := ctx.Value(k); v != nil {
	//		fmt.Println("found value:", v)
	//		return
	//	}
	//	fmt.Println("key not found:", k)
	//}
	//k := "language"
	//ctx := context.WithValue(context.Background(), k, "Go")
	//
	//f(ctx, k)
	//f(ctx, "color")

	ctx, cFunc := context.WithCancel(context.Background())

}

func GoroutineDemo() {
	contextDemo()
}
