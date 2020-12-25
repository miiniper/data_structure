//二叉树 binary tree
package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

//bfs
//前序遍历：根左右
//中序遍历：左根右
//后序遍历：左右根
//查找1
//修改
//删除1
//插入1
//建树1
//深度（层数）1
//宽度1
//最值1

//生成一个有序二叉树即：左节点小于当前节点，右节点大于当前节点
//默认 va := []int{100, 114, 62, 123, 78, 45, 43, 70, 89, 90}
//default tree
func NewNode(data int) *Node {
	return &Node{Data: data}
}

//readme的二叉树
func DTree() *Node {
	var tree *Node
	va := []int{100, 114, 62, 123, 78, 45, 43, 70, 89, 90}
	for _, i := range va {
		tree = TreeAdd(tree, i)
	}
	return tree
}

func NewT() *Node {
	return &Node{}
}

//数组变二叉树
func (t *Node) NewTreeArr(array []int) *Node {
	if len(array) == 0 {
		return t
	}
	for _, i := range array {
		//		fmt.Printf("k==%d,i==%d\n", k, i)
		t.TreeAddNode(NewNode(i))
	}
	return t
}

//add node to tree
func (t *Node) TreeAddNode(n *Node) {
	if t.Data == 0 {
		t.Data = n.Data
	}
	if n.Data == t.Data {
		return
	}
	if n.Data < t.Data {
		if t.Left == nil {
			t.Left = n
		} else {
			t.Left.TreeAddNode(n)
		}
	} else {
		if t.Right == nil {
			t.Right = n
		} else {
			t.Right.TreeAddNode(n)
		}
	}
}

//add data  to tree
func TreeAdd(t *Node, v int) *Node {
	if t == nil {
		return &Node{Data: v}
	}
	if t.Data == v {
		return t
	}
	if v < t.Data {
		t.Left = TreeAdd(t.Left, v)
	} else {
		t.Right = TreeAdd(t.Right, v)
	}
	return t
}

//遍历，默认前序遍历
func PTree(t *Node) {
	if t == nil {
		return
	}
	fmt.Println(t.Data)
	PTree(t.Left)
	//fmt.Println(t.Data)  //中序遍历
	PTree(t.Right)
	//fmt.Println(t.Data) //后序遍历
}

//广度优先遍历
func Bfs(t *Node) {
	//res := make([]int, 0)
	if t == nil {
		return
	}

	q := []*Node{t}
	for len(q) > 0 {
		length := len(q)
		for length > 0 {
			length--
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			fmt.Printf("%d,", q[0].Data)
			q = q[1:]
		}
	}
}
func (t *Node) Width() int {
	//res := make([]int, 0)
	if t == nil {
		return 0
	}
	width := 1
	q := []*Node{t}
	for len(q) > 0 {
		length := len(q)
		if length > width {
			width = length
		}
		for length > 0 {
			length--
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			q = q[1:]
		}
	}
	return width
}

//搜索
func (t *Node) TSearch(data int) (error, *Node) {
	if t == nil {
		return errors.New("data not found"), nil
	}
	if data == t.Data {
		return nil, t
	}
	if data < t.Data {
		return t.Left.TSearch(data)
	}
	return t.Right.TSearch(data)
}

//修改：先删除在插入
//删除
//第一种情况，删除的是根节点，且根节点没有儿子，直接删除即可。
//第二种情况，删除的节点有父亲节点，但没有子树，也就是删除的是叶子节点，直接删除即可。
//第三种情况，删除的节点只有一个子树，那么该子树直接替换被删除的节点即可。
//第四种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点，
//这时二叉查找树的性质又满足了。右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到。
func DeleteNode(t, data *Node) *Node {
	if t == nil {
		return nil
	}
	if data.Data < t.Data {
		//left todo
		t.Left = DeleteNode(t.Left, data)
		return t
	}
	if data.Data > t.Data {
		//right todo
		t.Right = DeleteNode(t.Right, data)
		return t
	}
	//只有根
	if t.Left == nil && t.Right == nil {
		return nil
	}
	//存在右子树
	if t.Left == nil && t.Right != nil {
		t = t.Right
		return t
	}
	//存在左子树
	if t.Right == nil && t.Left != nil {
		t = t.Left
		return t
	}
	//左右子树都存在，找到右子树最左节点替换
	index := t.Right.MinNode()
	fmt.Println("====", index.Data)
	t.Data = index.Data
	t.Right = DeleteNode(t.Right, t)
	return t

}

//树的高度
func (t *Node) Depth() int {
	if t == nil {
		return 0
	}
	return Max(t.Right.Depth(), t.Left.Depth()) + 1
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//最大值:右树叶子节点
func (t *Node) MaxNode() *Node {
	if t.Right == nil {
		return t
	}
	return t.Right.MaxNode()
}

//最小值:左树叶子节点
func (t *Node) MinNode() *Node {
	if t.Left == nil {
		return t
	}
	return t.Left.MinNode()
}

//test search demo
func demoSearch(t *Node) {
	err, t2 := t.TSearch(90)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("search ", t2.Data)

}

func demoDelete(t *Node) {
	d := NewNode(78)
	t2 := DeleteNode(t, d)
	Bfs(t2)
}

//翻转二叉树
func invertTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

//测试数组
var array = []int{100, 114, 62, 123, 78, 45, 43, 70, 89, 90}

func main() {
	t1 := NewT()
	t1.NewTreeArr(array)
	//PTree(t1)
	Bfs(t1)
	//demoSearch(t1)
	//fmt.Println("minV == ", t1.MinNode().Data)
	//fmt.Println("maxV == ", t1.MaxNode().Data)
	//fmt.Println("depth == ", t1.Depth())
	//fmt.Println("width == ", t1.Width())
	fmt.Println()
	demoDelete(t1)

}
