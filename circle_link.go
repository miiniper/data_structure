//单向环形列表
package main

import (
	"fmt"
)

type CNode struct {
	On   int    `json:"on"`
	Name string `json:"name"`
	Next *CNode
}

func main() {
	fmt.Println("####################circle link#############################")
	head := &CNode{}

	ShowCLink(head)

	c1 := &CNode{On: 1, Name: "c1"}
	c2 := &CNode{On: 2, Name: "c2"}
	c3 := &CNode{On: 3, Name: "c3"}
	c4 := &CNode{On: 4, Name: "c4"}

	AddNode(head, c1)
	AddNode(head, c2)
	AddNode(head, c3)
	AddNode(head, c4)
	ShowCLink(head)
	fmt.Println("")
	fmt.Println("delete ")
	DelCNode(head, 1)
	ShowCLink(head)

}

func AddNode(head, newNode *CNode) {
	if head.Next == nil {
		head.Name = newNode.Name
		head.On = newNode.On
		head.Next = head
		return

	}
	tmp := head
	for {
		if tmp.Next == head {
			tmp.Next = newNode
			newNode.Next = head
			break
		}
		tmp = tmp.Next
	}

}

func ShowCLink(head *CNode) {
	if head.Next == nil {
		fmt.Println("circle link is empty")
		return
	}
	tmp := head
	for {
		fmt.Printf("number = %d,name = %s -> ", tmp.On, tmp.Name)
		if tmp.Next == head {
			break
		}
		tmp = tmp.Next
	}
}
func DelCNode(head *CNode, i int) *CNode {
	tmp := head
	del := false
	// 只有一个node，并且删除，返回空link
	if tmp.Next == head && tmp.On == i {
		head.Next = nil
		return head
	}

	//删除head ，需要重新找一个head返回
	if head.On == i {
		head.On = head.Next.On
		head.Name = head.Next.Name
		head.Next = head.Next.Next
		return head
	}

	for {
		if tmp.Next.On == i {
			del = true
			break
		}
		tmp = tmp.Next
		if tmp.Next == head {
			break
		}
	}

	//正常删除

	if del {
		tmp.Next = tmp.Next.Next
	} else {
		fmt.Println("not found ", i)
	}

	return head

}
