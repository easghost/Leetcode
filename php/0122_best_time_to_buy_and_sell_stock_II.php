<?php
// Runtime: 16 ms
// Memory Usage: 17.7 MB

class Solution {
    /**
     * @param int[] $prices
     * @return int
     */
    function maxProfit($prices) {
        $res = 0;
        for ($i = 0; $i < count($prices); $i++) {
            if ($prices[$i] < $prices[$i+1]) {
                $res += $prices[$i+1] - $prices[$i];
            }
        }
        
        return $res;
    }
}