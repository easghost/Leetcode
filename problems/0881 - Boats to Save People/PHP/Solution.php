<?php

class Solution {

    /**
     * @param Integer[] $people
     * @param Integer $limit
     * @return Integer
     */
    function numRescueBoats($people, $limit) {
        sort($people);
        $res = 0;
        $start = 0;
        
        for ($end = count($people)-1; $start <= $end; $end--) {
            if (
                $end > 0 
                && $people[$end] + $people[$end-1] <= $limit
            ) {
                $end--;    
            } elseif (
                $start < $end
                && $people[$end] + $people[$start] <= $limit
            ) {
                $start++;
            }
            
            $res++;
        }
        
        return $res;
    }
}