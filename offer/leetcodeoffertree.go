package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := []int{3, 9, 20, 15, 7}
	b := []int{9, 3, 15, 20, 7}
	t := buildTree(a, b)
	fmt.Println(t)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

}
