package main

import (
	"fmt"
	"time"
)

func execute(ch chan string) {
	fmt.Println("start ...")
	time.Sleep(time.Second * 3)
	ch <- "done"
	fmt.Println("end ...")
}

func main() {

	ch := make(chan string)
	go execute(ch)
	select {
	case res := <-ch:
		fmt.Println(res)

	case <-time.After(time.Second * 2):
		fmt.Println("result timeout!")
	}

}
