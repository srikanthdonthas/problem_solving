package populatingnextrightpointers

/*
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children.
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
Initially, all next pointers are set to NULL.

Example 1:
Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
*/


type Node struct { Val int; Left *Node; Right *Node; Next *Node }
func connect(root *Node) *Node {
	return nil
}
