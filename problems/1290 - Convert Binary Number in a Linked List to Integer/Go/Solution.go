// Runtime: 0 ms
// Memory Usage: 2.2 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res<<1 + head.Val
		head = head.Next
	}
	return res
}