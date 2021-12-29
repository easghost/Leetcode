// Runtime: 4 ms
// Memory Usage: 6.5 MB

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	root.Left = connect(root.Left)
	root.Right = connect(root.Right)

	return root
}