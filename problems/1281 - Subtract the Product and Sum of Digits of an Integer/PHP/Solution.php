<?php

class Solution {

/**
 * @param Integer $n
 * @return Integer
 */
function subtractProductAndSum($n) {
    $sum = 0;
    $prod = 1;
    
    while ($n >= 1) {
        $sum += $n % 10;
        $prod *= $n % 10;
        $n = $n / 10;
    }
    
    return $prod - $sum;
}
}