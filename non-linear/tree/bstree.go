package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (n *Node) InsertDataToNode(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = NewNode(data)
		} else {
			n.left.InsertDataToNode(data)
		}
	} else {
		if n.right == nil {
			n.right = NewNode(data)
		} else {
			n.right.InsertDataToNode(data)
		}
	}
}

type BSTree struct {
	root *Node
}

func (bst *BSTree) InsertNodeToTree(data int) *BSTree {
	if bst.root == nil {
		bst.root = NewNode(data)
	} else {
		bst.root.InsertDataToNode(data)
	}

	return bst
}

func (bst *BSTree) DFS() {
	if bst.root == nil {
		return
	}

	bst.InOrderTraversal(bst.root, 0)
}

func (bst *BSTree) InOrderTraversal(node *Node, level int) {
	if node == nil {
		return
	}

	bst.InOrderTraversal(node.left, level+1)
	fmt.Printf("%s%d\n", strings.Repeat(" * ", level), node.data)

	bst.InOrderTraversal(node.right, level+1)
}

func (bst *BSTree) InterativeInorderTraversal() {
	curr := bst.root
	fmt.Println(curr.data)
	var stack = []int{curr.data}

	for len(stack) > 0 {
		if curr != nil {
			stack = append(stack, curr.data)
			curr = curr.left
			fmt.Println(curr.data)
		} else {
			stack = append(stack[:len(stack)-1], stack[len(stack)+1:]...)
		}
	}

	fmt.Println(stack)
}

func (bst *BSTree) BFS() {
	if bst.root == nil {
		return
	}

	queue := []*Node{bst.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d -> ", node.data)

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

/* Construct the following tree
          10
        /   \
       /     \
      20       30
     /      /   \
    /      /     \
   40      50      60
         / \
        /   \
       70     80
*/

func main() {
	bst := &BSTree{}
	bst.InsertNodeToTree(10).
		InsertNodeToTree(20).
		InsertNodeToTree(30).
		InsertNodeToTree(40).
		InsertNodeToTree(50).
		InsertNodeToTree(60).
		InsertNodeToTree(70).
		InsertNodeToTree(80)

	bst.InterativeInorderTraversal()
}
