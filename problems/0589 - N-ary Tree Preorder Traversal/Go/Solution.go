/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func preorder(root *Node) []int {
    res := []int{}
    if root == nil {
        return res
    }
    stack := []*Node{ root }
    for len(stack) > 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        res = append(res, pop.Val)
        for i := len(pop.Children) - 1; i >= 0; i-- {
            stack = append(stack, pop.Children[i])
        }
    }
    return res
}