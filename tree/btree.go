//二叉树 binary tree
package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

//前序遍历：根左右
//中序遍历：左根右
//后序遍历：左右根
//二叉树查找
//节点修改
//插入
//建树
//深度（层数）

//new b tree
func NewTree() *Node {
	var tree *Node
	va := []int{100, 114, 62, 123, 78, 45, 43, 70, 89, 90}
	for _, i := range va {
		tree = TreeAdd(tree, i)
	}
	return tree
}

//add node to tree
func TreeAdd(t *Node, v int) *Node {
	if t == nil {
		return &Node{Data: v}
	}
	if v <= t.Data {
		t.Left = TreeAdd(t.Left, v)
	} else {
		t.Right = TreeAdd(t.Right, v)
	}

	return t
}

func PreTree(t *Node) {
	if t == nil {
		return
	}
	//PreTree(t.Left)
	fmt.Println(t.Data)
	PreTree(t.Left)
	PreTree(t.Right)
	//fmt.Println(t.Data)
}

func main() {
	t1 := NewTree()
	PreTree(t1)
	//fmt.Println(t1.Data)
}
