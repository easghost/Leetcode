<?php

class Solution {

    /**
     * @param Integer[] $arr
     * @return Integer
     */
    function sumOddLengthSubarrays($arr) {
        $res = 0;
        $l = count($arr);
        
        for ($n = 1; $n <= $l; $n += 2) {
            for ($i = 0; $i < $l-($n-1); $i++) {
                for ($j = 0; $j < $n; $j++) {
                    $res += $arr[$i+$j];
                }
            }
        }
        
        return $res;
    }
}