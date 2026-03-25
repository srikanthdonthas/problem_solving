package serializeanddeserializebinarytree

/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.
Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work.

Example 1:
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
*/


// Codec struct
type Codec struct {}
func Constructor() Codec { return Codec{} }
func (this *Codec) serialize(root *TreeNode) string { return "" }
func (this *Codec) deserialize(data string) *TreeNode { return nil } {
	
}
