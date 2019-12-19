// single queue
// front和rear 指向不包含首的元素
package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	Maxsize int    `json:"maxsize"`
	Front   int    `json:"front"`
	Rear    int    `json:"rear"`
	Array   [5]int `json:"array"`
}

func main() {
	fmt.Println("################queue array##################")
	queue := Queue{Front: -1,
		Rear:    -1,
		Maxsize: 5,
	}

	var key string
	var val int
	for {

		fmt.Println("s   : show queue")
		fmt.Println("add : add num to queue")
		fmt.Println("get : get num from queue ")
		fmt.Println("exit: exit proccess")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("input your num")
			fmt.Scanln(&val)
			err := queue.Add(val)
			if err != nil {
				fmt.Println("add  num fail ..........")
			} else {
				fmt.Println("add num success.........")
			}
		case "get":
			//fmt.Println("out key ")
			val, err := queue.Out()
			if err != nil {
				fmt.Println("get val fail ")
			} else {
				fmt.Println("get num ====", val)
			}
		case "s":
			queue.show()

		case "exit":
			os.Exit(0)

		}
	}

}

func (q *Queue) Add(val int) error {
	if q.Rear == q.Maxsize-1 {
		fmt.Println("queue is full ")
		return errors.New("queue is full")
	}
	q.Rear++
	q.Array[q.Rear] = val
	return nil
}

func (q *Queue) show() {

	if q.Front == q.Rear {
		fmt.Println("nothing s")
		return
	}

	for i := q.Front + 1; i <= q.Rear; i++ {
		fmt.Printf("array is : num==%d , val == %d\n", i, q.Array[i])
	}

}

func (q *Queue) Out() (int, error) {
	if q.Rear == q.Front {
		err := errors.New("queue is empty")
		return -1, err
	}
	q.Front++
	return q.Array[q.Front], nil
}
