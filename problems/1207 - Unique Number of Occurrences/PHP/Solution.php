<?php

class Solution {

    /**
     * @param Integer[] $arr
     * @return Boolean
     */
    function uniqueOccurrences($arr) {
        $m = [];
        $m1 = [];
        foreach ($arr as $item) {
            $m[$item]++;
            
        }

        foreach ($m as $v) {
            $m1[$v] = true;
        }

        return count($m) === count($m1);
    }
}