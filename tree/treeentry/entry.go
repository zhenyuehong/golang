package main

import (
	"fmt"
	"golang/tree"
)

//树的遍历
//前序	中	左	右
//中序	左	中	右
//后序	左	右	中
type myTreeNode struct {
	node *tree.Node
}

//从左到右
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 5}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{2, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Print("In-order traversal: ")
	root.Traverse() //中序遍历树

	fmt.Print("My own post-order traversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder() //后序遍历

	//节点统计
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	//查找最大節點
	c := root.TraverseWithChannel()
	maxNodeValue := 0
	for node := range c {
		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}
	}
	fmt.Println("Max node value:", maxNodeValue)
}
