/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    Next := head.Next
    for Next != nil  {
        if Next.Val == head.Val {
            Next = Next.Next
        } else {
            break;
        }
    }
    head.Next = deleteDuplicates(Next)
    
    return head 
}