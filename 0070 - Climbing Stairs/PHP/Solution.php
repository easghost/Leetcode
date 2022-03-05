<?php

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function climbStairs($n) {
        $x=1;
        $y=1;
        for($i=1; $i < $n ; $i++){
            $t=$y;
            $y=$x+$y;
            $x=$t;
        }
        return $y;
    }
}