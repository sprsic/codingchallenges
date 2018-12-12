package main

import (
	"fmt"
)

/*
	A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.
	Given the root to a binary tree, count the number of unival subtrees.
*/
type Node struct {
	val int 
	left *Node 
	right *Node 
}


func main(){

/*
   0
  / \
 1   0
    / \
   1   0
  / \
 1   1
*/

	tree :=	 &Node{0, &Node{1, nil, nil}, &Node{0, &Node{1, &Node{1, nil, nil}, &Node{1, nil, nil}}, &Node{0, nil, nil}}}


	fmt.Printf("%d", countSubtree(tree))
}

func countSubtree(tree *Node) int {
	if tree == nil {
		return 0
	}
	// recursion unwrap tree
	left := countSubtree(tree.left)
	right := countSubtree(tree.right)
	if isUnival(tree, tree.val) {
		//it is unival so count that subtree +1
		return  1 + left + right
	} else {
		return left + right
	}
}

func isUnival(tree *Node, val int) bool {
	// basic case nil tree is unival
	if tree == nil {
		return true
	}
	// check the value with the subtrees
	if tree.val == val {
		return isUnival(tree.left, val) && isUnival(tree.right, val)
	}
	return false
}