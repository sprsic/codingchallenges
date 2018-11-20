package main

import (
	"fmt"
)

/*
Given the root to a binary tree, implement serialize(root), which serializes the tree into a array, and deserialize(s), which deserializes the array back into the tree.

For example, given the following Node class

class Node:
    def func (val, left=None, right=None):
	left = left
	right = right
	val = val
*/
//unexported
type Node struct {
	left  *Node
	right *Node
	val   int
}

func main() {

	tree := Node{&Node{&Node{nil, nil, 5}, nil, 15}, &Node{nil, nil, 10}, 46}

	arrayOfNodes := serializeDFS(&tree, []int{})

	fmt.Printf("Serialized DFS tree:  %v \n", serializeDFS(&tree, []int{}))
	fmt.Printf("Serialized BFS tree:  %v \n", serializeBFS(&tree, []int{}))

	deserialisedTree := deserializeNode(arrayOfNodes)
	fmt.Printf("Deserialised tree:    %v \n", serializeDFS(deserialisedTree, []int{}))
}

// DFS
func serializeDFS(root *Node, arr []int) []int {

	if root != nil {
		arr = append(arr, root.val)
		arr = serializeDFS(root.right, arr)
		arr = serializeDFS(root.left, arr)
	}
	return arr
}

func serializeBFS(root *Node, arrayOfNodes []int) []int {

	queue := []*Node{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		n := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if n.left != nil {
			queue = append(queue, n.left)
		}

		arrayOfNodes = append(arrayOfNodes, n.val)
		if n.right != nil {
			queue = append(queue, n.right)
		}

	}
	return arrayOfNodes
}

func deserializeNode(array []int) *Node {
	var tree *Node
	for _, node := range array {
		tree = insert(tree, node)
	}
	return tree
}

func insert(tree *Node, val int) *Node {

	if tree == nil {
		return &Node{nil, nil, val}
	}

	if tree.val >= val {
		tree.left = insert(tree.left, val)
	}

	if tree.val < val {
		tree.right = insert(tree.right, val)
	}
	return tree
}
