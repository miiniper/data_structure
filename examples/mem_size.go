package main

import (
	"fmt"
	"unsafe"
)

// Event 参考地址 https://golang.org/ref/spec#Size_and_alignment_guarantees
// 视频解析 ：https://www.bilibili.com/video/BV1Ja4y1i7AF?from=search&seid=11845169762769004814

type Demo1 struct {
	NS      bool  `json:"ns"`      //1
	ID      int   `json:"id"`      //8
	Message int16 `json:"message"` //2
}

type Demo2 struct {
	NS      bool  `json:"ns"`      //1
	Message int16 `json:"message"` //2
	ID      int   `json:"id"`      //8
}

type Demo3 struct {
	NS      bool  `json:"ns"`      //1
	Message int16 `json:"message"` //2
	ID      int   `json:"id"`      //8
	Host    Demo2 `json:"host"`    //16
}

type Demo4 struct {
	ID      int   `json:"id"`      //8
	NS      bool  `json:"ns"`      //1
	Host    Demo2 `json:"host"`    //16
	Message int16 `json:"message"` //2
}

func main() {
	fmt.Println("service starting ... ")
	fmt.Println("Demo1 size :", unsafe.Sizeof(Demo1{})) //Demo1 size : 24
	fmt.Println("Demo2 size :", unsafe.Sizeof(Demo2{})) //Demo2 size : 16
	fmt.Println("Demo3 size :", unsafe.Sizeof(Demo3{})) //Demo3 size : 32
	fmt.Println("Demo4 size :", unsafe.Sizeof(Demo4{})) //Demo4 size : 40

}
