<?php

class Solution {

    /**
     * @param Integer $n
     * @return Boolean
     */
    function isHappy($n) {
        if ($n === 1 || $n === 7) {
            return true;
        }
        while ($n >= 9) {
            $x = $n;
            $sum = 0;
            while($x > 0) {
                $sum += pow($x%10, 2);
                $x = floor($x/10);
            }

            if ($sum === 1 || $sum === 7) {
                return true;
            }
            $n = $sum;
        }
        
        return false;
    }
}