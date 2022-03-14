<?php

class Solution {

    /**
     * @param Integer[] $arr
     * @return Boolean
     */
    function canMakeArithmeticProgression($arr) {
        sort($arr);
        $diff = $arr[1]-$arr[0];
        
        for ($i = 0; $i < count($arr)-1; $i++) {
            if (($arr[$i+1]-$arr[$i]) !== $diff) {
                return false;
            }
        }
        
        return true;
    }
}