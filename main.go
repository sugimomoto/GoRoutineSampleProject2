package main

import (
	"fmt"
	"strconv"
)

func main() {
	BufferSample()
}

func BufferSample() {

	// Bufferを指定していないと受信が存在しないので、デッドロックが発生する
	// Bufferありチャネルだと、とりあえず受信が存在しなくても、そのバッファーまでは送信が可能
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("End")

}

func NonBufferSample() {
	ch := make(chan string)

	go async(ch, 10)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	/* 送信が不足しているので、受信がいつまでも待機してしまい、Deadlockが発生する
	fmt.Println(<-ch)
	*/

}

func async(ch chan string, count int) {
	for i := 0; i < count; i++ {
		// チャネルが受信されるまで、送信は待機される
		ch <- strconv.Itoa(i)
	}
}
