<?php

class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function countVowelStrings($n) {
        $res1 = array_fill(0,5,1);
        $res2 = [];
        
        for ($i = 1; $i <= $n; $i++) {
            $res2[0] = 1;
            for ($j = 1; $j < 5; $j++) {
                $res2[$j] = $res2[$j - 1] + $res1[$j];
            }
            
            $res1 = $res2;
            $res2 = $res1;
        }
        
        return $res1[4];
    }
}
