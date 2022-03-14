<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function largestPerimeter($nums) {
        sort($nums);
        for ($i = count($nums)-1; $i >= 2; $i--) {
            if ($this->isTriangle($nums[$i], $nums[$i-1], $nums[$i-2]) === true) {
                return $nums[$i] + $nums[$i-1] + $nums[$i-2];
            }
        }

        return 0;
    }

    function isTriangle(int $a, int $b, int $c): bool
    {
        return ($a + $b) > $c && ($a + $c) > $b && ($b + $c) > $a;
    }
}