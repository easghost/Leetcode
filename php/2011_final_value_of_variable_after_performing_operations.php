<?php
// Runtime: 12 ms
// Memory Usage: 15.6 MB

class Solution {
    /**
     * @param String[] $operations
     * @return Integer
     */
    function finalValueAfterOperations($operations) {
        $res = 0;
        foreach ($operations as $operation) {
            if (strpos($operation, '--') !== false) {
                $res--;
            } else {
                $res++;
            }
        }
        
        return $res;
    }
}