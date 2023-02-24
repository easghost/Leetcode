<?php

class Solution {

    /**
     * @param Integer $startValue
     * @param Integer $target
     * @return Integer
     */
    function brokenCalc($startValue, $target) {
        $res = 0;
        while ($startValue < $target) {
            if ($target % 2 === 0) {
                $target /= 2;
            } else {
                $target++;
            }
            $res++;
        }
        
        return $res + $startValue - $target;
    }
}