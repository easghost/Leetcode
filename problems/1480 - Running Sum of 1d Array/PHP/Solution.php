<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function runningSum($nums) {
        if (count($nums) < 1) {
            return $nums;
        }
        
        $res = [];
        $sum = 0;
        foreach ($nums as $num) {
            $sum += $num;
            $res[] = $sum;
        }
        
        return $res;
    }
}