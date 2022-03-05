<?php
// Runtime: 148 ms
// Memory Usage: 15.7 MB

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function arrangeCoins($n) {
        $res = 1;
        while ($n >= $res) {
            $n -= $res;
            $res++;
        }
        
        return $res - 1;
    }
}