package main

import "fmt"

type PeoNode struct {
	Num  int      `json:"num"`
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Next *PeoNode `json:"next"`
}

func main() {
	fmt.Println("####################single link###################################")
	head := &PeoNode{Name: "head"}
	p1 := &PeoNode{Name: "p1", Age: 11, Num: 1}
	p2 := &PeoNode{Name: "p2", Age: 12, Num: 2}
	p3 := &PeoNode{Name: "p3", Age: 13, Num: 3}
	p4 := &PeoNode{Name: "p4", Age: 14, Num: 4}

	insertLink2(head, p2)
	insertLink2(head, p4)
	insertLink2(head, p1)
	insertLink2(head, p3)

	showLink(head)

	fmt.Println()
	delLink(head, 2)
	showLink(head)

}

//列表最后插入
func insertLink(head, newnode *PeoNode) {
	tmp := head
	for {
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	tmp.Next = newnode

}

//遍历链表
func showLink(head *PeoNode) {
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

// 排序插入
func insertLink2(head, newnode *PeoNode) {
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
	} else {
		newnode.Next = tmp.Next
		tmp.Next = newnode

	}

}

func delLink(head *PeoNode, i int) {
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

	} else {
		fmt.Println("not found ", i)
	}

}
