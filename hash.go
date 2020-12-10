package main

import (
	"fmt"
	"strconv"
)

type Stu struct {
	Id   int
	Name string
	Next *Stu
}

//链表添加
func AddLink2(head, node *Stu) {
	tmp := head
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	tmp.Next = node
}

//选择头部
func SelectHead(Size, no int) int {
	return Size % no
}

//遍历所有
func showHLink(heads []Stu) {
	for _, head := range heads {
		tmp := head
		if tmp.Next == nil {
			fmt.Println("link is empty")
			return
		}
		for {
			fmt.Printf("[%d,%v] -> ", tmp.Id, tmp.Name)

			if tmp.Next == nil {
				break
			}
			tmp = *tmp.Next
		}
		fmt.Println()
	}

}

//生成模拟数据
func defaultStus(num int) []*Stu {
	ss := []*Stu{}
	for i := 0; i < num; i++ {
		s := &Stu{Name: "student" + strconv.Itoa(i), Id: i}
		ss = append(ss, s)
	}
	return ss
}

//查看头部数组
func showH(h []Stu) {
	for _, j := range h {
		fmt.Printf("[%d,%s] -> ", j.Id, j.Name)
	}
	fmt.Println()

}

func main() {
	//todo
	headLink := make([]Stu, 7)
	for i := 0; i < 7; i++ {
		p := Stu{Name: "head" + strconv.Itoa(i)}
		headLink[i] = p
	}

	showH(headLink)

	s1 := defaultStus(50)
	for _, j := range s1 {
		headNo := SelectHead(j.Id, 7)
		AddLink2(&headLink[headNo], j)
	}

	showHLink(headLink)
}
