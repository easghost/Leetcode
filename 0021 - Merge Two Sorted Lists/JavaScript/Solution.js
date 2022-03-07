/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
 var mergeTwoLists = function(list1, list2) {
    let list = new ListNode()
    let head = list
    
    while (list1 !== null && list2 !== null) {
        if (list1.val < list2.val) {
            list.next = new ListNode(list1.val)
            list1 = list1.next
        } else {
            list.next = new ListNode(list2.val)
            list2 = list2.next
        }
        
        list = list.next
    }
    
    if (list1 !== null) {
        list.next = list1
    } else if (list2 !== null) {
        list.next = list2    
    }
    
    return head.next
};