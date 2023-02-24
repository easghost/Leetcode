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
        if ($head === null || $head->next === null) {
            return $head;
        }
        
        $hMap = [];
        for ($i = $head; $i !== null; $i = $i->next) {
            $hMap[$i->val]++;
        }
        
        $prev = $head;
        $curr = $head->next;
        
        while ($curr !== null) {
            if ($hMap[$curr->val] > 1) {
                $prev->next = $curr->next;
                $curr = $curr->next;
            } else {
                $prev = $curr;
                $curr = $curr->next;
            }
        }
        
        if ($hMap[$head->val] > 1) {
            return $head->next;
        }
        
        return $head;
    }
}