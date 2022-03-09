/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    hMap := make(map[int]int)
    for curr := head; curr != nil; curr = curr.Next {
        hMap[curr.Val]++
    }
    
    prev, curr := head, head.Next
    for curr != nil {
        if hMap[curr.Val] > 1 {
            prev.Next = curr.Next
            curr = curr.Next
        } else {
            prev = curr
            curr = curr.Next
        }
    }
    if hMap[head.Val] > 1 {
        return head.Next
    }
    
    return head
}