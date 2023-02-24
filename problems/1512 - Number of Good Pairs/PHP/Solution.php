<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function numIdenticalPairs($nums) {
        $hMap = (object)[];
        $res = 0;
		
        foreach ($nums as $num) {
            $res += $hMap->$num;
            $hMap->$num++;
        }
        
		return $res;
    }
}