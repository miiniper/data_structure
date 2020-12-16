package main

import "fmt"

type linkReverse struct {
	Val  int
	Next *linkReverse
}

func main() {
	head := &linkReverse{Val: 3, Next: &linkReverse{Val: 2, Next: &linkReverse{Val: 1}}}
	fmt.Println(reversePrint1(head))
	fmt.Println(reversePrint2(head))
}

//链表反向输出：从尾到头输出
//解法1：递归实现
func reversePrint1(head *linkReverse) []int {
	res := []int{}
	if head == nil {
		return res
	}

	return append(reversePrint1(head.Next), head.Val)
}

//解法2：栈思想
type stack struct {
	Val []int
}

func (s *stack) pop() int {
	v := s.Val[len(s.Val)-1]
	s.Val = s.Val[:len(s.Val)-1]
	return v
}
func (s *stack) push(v int) {
	s.Val = append(s.Val, v)
}
func reversePrint2(head *linkReverse) []int {
	s := stack{}
	res := []int{}
	if head == nil {
		return res
	}
	tmp := head
	for {
		s.push(tmp.Val)
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	fmt.Println(s.Val)
	for {
		if len(s.Val) == 0 {
			break
		}
		res = append(res, s.pop())
	}
	return res
}
