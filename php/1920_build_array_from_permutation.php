<?php
// Runtime: 40 ms
// Memory Usage: 15.8 MB

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function buildArray($nums) {
        $result = [];
        foreach($nums as $key => $value) {
            $result[$key] = $nums[$nums[$key]];
        }
        
        return $result;
    }
}