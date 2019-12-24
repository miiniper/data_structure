package main

import "fmt"

type PeoNode2 struct {
	Num  int       `json:"num"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
	Next *PeoNode2 `json:"next"`
	Pre  *PeoNode2 `json:"pre"`
}

func main() {
	fmt.Println("#######################double link#################################")
	head := &PeoNode2{Name: "head"}
	p1 := &PeoNode2{Name: "p1", Age: 11, Num: 1}
	p2 := &PeoNode2{Name: "p2", Age: 12, Num: 2}
	p3 := &PeoNode2{Name: "p3", Age: 13, Num: 3}
	p4 := &PeoNode2{Name: "p4", Age: 14, Num: 4}
	showDLink(head)
	insertDLink2(head, p3)
	insertDLink2(head, p1)
	insertDLink2(head, p2)
	insertDLink2(head, p4)
	showDLink(head)
	fmt.Println("逆向输出")
	showDLink2(head)
	delDLink(head, 4)
	fmt.Println()
	showDLink(head)
}

//在链表末尾插入
func insertDLink(head, newnode *PeoNode2) {
	tmp := head
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	tmp.Next = newnode
	newnode.Pre = tmp

}

//正向输出，next实现
func showDLink(head *PeoNode2) {
	tmp := head
	if tmp.Next == nil {
		fmt.Println("link is empty")
		return
	}

	for {
		fmt.Printf("[%d,%s,%d] -> ", tmp.Num, tmp.Name, tmp.Age)

		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}

}

//逆向输出，根据pre实现
func showDLink2(head *PeoNode2) {
	tmp := head

	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}

	if tmp.Pre == nil {
		fmt.Println("link is empty")
		return
	}

	for {
		fmt.Printf("[%d,%s,%d] -> ", tmp.Num, tmp.Name, tmp.Age)

		if tmp.Pre == nil {
			break
		}
		tmp = tmp.Pre
	}

}

//有序的插入双链表
func insertDLink2(head, newnode *PeoNode2) {
	tmp := head
	fff := true
	for {
		if tmp.Next == nil {
			break
		} else if tmp.Next.Num > newnode.Num {
			break
		} else if tmp.Next.Num == newnode.Num {
			fff = false
			break
		}
		tmp = tmp.Next
	}
	if !fff {
		fmt.Printf("%d已存在存在\n", newnode.Num)
		return
	}
	newnode.Pre = tmp
	newnode.Next = tmp.Next
	tmp.Next = newnode
	if newnode.Next != nil { //判断是否是最后一个节点
		newnode.Next.Pre = newnode
	}

}

func delDLink(head *PeoNode2, i int) {
	tmp := head
	fff := false
	for {
		if tmp.Next == nil {
			break
		} else if tmp.Next.Num == i {
			fff = true
			break
		}
		tmp = tmp.Next
	}
	if fff {
		tmp.Next = tmp.Next.Next
		if tmp.Next != nil { //删除的是否是最后一个节点
			tmp.Next.Pre = tmp
		}
	} else {
		fmt.Println("not found ", i)
	}

}
