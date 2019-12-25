//约瑟夫问题
package main

import (
	"fmt"
)

type Child struct {
	No   int    `json:"no"`
	Next *Child `json:"next"`
}

func main() {
	fmt.Println("#################约瑟夫问题##########################")
	c1 := NewChild(41)
	ShowChild(c1)

	fmt.Println("")
	Josephus(c1, 1, 3)
}

/*参考韩顺平讲解
两个node，一个head，一个head的后一个。
删除head所在的node
*/
func Josephus(first *Child, start, count int) {

	if first == nil {
		fmt.Println("不能进行游戏")
		return
	}

	tail := first
	// 使tail指向head前一个，也就是最后
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	//使first从start出发，tail跟随first移动
	for {
		if first.No == start {
			break
		}
		first = first.Next
		tail = tail.Next
	}
	//上面的循环结束后first指向start的node

	for {

		//开始移动count -1次：first指向的节点就是删除的node
		for i := 1; i <= count-1; i++ {
			first = first.Next
			tail = tail.Next
		}

		//进行删除
		fmt.Printf("%d is out \n", first.No)

		first = first.Next // first 先前移动一个
		tail.Next = first  //tail 指向first

		if tail == first {
			break
		}

	}

	fmt.Println("留下来的child ：", first.No)

}

func NewChild(num int) *Child {
	if num < 1 {
		fmt.Println("child must > 1")
		return nil
	}

	head := &Child{}
	tmp := &Child{}

	//使用两个node，head不动指向头
	//tmp作为临时node，指向最尾，且tmp.next -> head,形成环
	//cc 为新的node添加到tmp后， cc变为最尾的tmp,且tmp.next -> head,形成环

	for i := 1; i <= num; i++ {
		cc := &Child{No: i}
		if i == 1 {
			head = cc
			head.Next = head
			tmp = head
		} else {
			tmp.Next = cc   //断开环，添加新node
			tmp = cc        //  tmp 成为最后的node
			tmp.Next = head //形成环

		}
	}

	return head
}

func ShowChild(first *Child) {

	if first == nil {
		fmt.Println("child is empty")
		return
	}
	tmp := first
	for {
		fmt.Printf("NO = %d -> ", tmp.No)
		if tmp.Next == first {
			break
		}
		tmp = tmp.Next
	}
}
