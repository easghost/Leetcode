/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	res := []*ListNode{}
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	return res[len(res)/2]
}