<?php

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function deleteDuplicates($head) {
        if ($head === null) {
            return null;
        }
        $next = $head->next;
        while ($next != null) {
            if ($next->val == $head->val) {
                $next = $next->next;
            } else {
                break;
            }
        }
        
        $head->next = $this->deleteDuplicates($next);
        
        return $head;
    }
}