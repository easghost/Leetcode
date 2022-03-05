/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	m := make(map[*ListNode]bool)
	for cur != nil {
		if _, ok := m[cur]; ok {
			return cur
		}

		m[cur] = true
		cur = cur.Next
	}

	return nil
}