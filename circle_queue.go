// 环形队列
package main

import (
	"errors"
	"fmt"
	"os"
)

type CQueue struct {
	MaxSize int    `json:"maxsize"`
	Array   [5]int `json:"array"`
	Head    int    `json:"head"`
	Tail    int    `json:"tail"`
}

func main() {

	fmt.Println("################cirlce queue###########################")

	queue := CQueue{
		Head:    0,
		Tail:    0,
		MaxSize: 5,
	}

	var key string
	var val int
	for {
		fmt.Println("############### menu ######################")
		fmt.Println("s   : show queue")
		fmt.Println("add : add num to queue")
		fmt.Println("get : get num from queue ")
		fmt.Println("exit: exit proccess")
		fmt.Println("############################################")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("input your num")
			fmt.Scanln(&val)
			err := queue.AddQ(val)
			if err != nil {
				fmt.Println("add  num fail ..........")
			} else {
				fmt.Println("add num success.........")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println("get val fail ", err)
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

func (q *CQueue) IsFull() bool {
	return (q.Tail+1)%q.MaxSize == q.Head
}

func (q *CQueue) IsEmpty() bool {
	return q.Head == q.Tail
}

func (q *CQueue) AddQ(val int) error {
	if q.IsFull() {
		return errors.New("q is full")
	}

	q.Array[q.Tail] = val
	q.Tail = (q.Tail + 1) % q.MaxSize
	return nil
}

func (q *CQueue) Pop() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("q is empty")
	}

	aa := q.Array[q.Head]
	q.Array[q.Head] = 0

	q.Head = (q.Head + 1) % q.MaxSize

	fmt.Println(q.Head, q.Tail)
	return aa, nil
}

func (q *CQueue) show() {
	if q.IsEmpty() {
		fmt.Println("q is empty")
		return
	}
	tmp := q.Head
	for i := 0; i < q.size(); i++ {
		fmt.Printf("array[%d] === %d \n", tmp, q.Array[tmp])
		tmp = (tmp + 1) % q.MaxSize
	}

	fmt.Println(q)

}

func (q *CQueue) size() int {
	return (q.Tail + q.MaxSize - q.Head) % q.MaxSize
}
