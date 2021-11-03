<?php
// Runtime: 8 ms
// Memory Usage: 15.6 MB

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer[] $index
     * @return Integer[]
     */
    function createTargetArray($nums, $index) {
        $res = [];
        
        foreach ($nums as $key => $val) {
            array_splice($res, $index[$key], 0, $val);
        }
        
        return $res;
    }
}