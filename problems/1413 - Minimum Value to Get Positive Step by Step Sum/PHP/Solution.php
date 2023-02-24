<?php
// Runtime: 4 ms
// Memory Usage: 15.7 MB

class Solution {
    /**
     * @param int[] $nums
     * @return int
     */
    function minStartValue($nums) {
        $curr = 1;
        $res = 1;
        foreach ($nums as $num) {
            $curr += $num;
            if ($curr < 1) {
                $res += 1 - $curr;
                $curr = 1;
            }
        }
        
        return $res;
    }
}