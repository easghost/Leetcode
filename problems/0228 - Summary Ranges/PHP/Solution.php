<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return String[]
     */
    function summaryRanges($nums) {
        $res = [];
    
        for ($i = 0; $i < count($nums); $i++) {
            $j = $i;
            while ($j < count($nums) - 1 && $nums[$j] == $nums[$j + 1] - 1) {
                $j++;
            }
            
            if ($i == $j) {
                array_push($res, strval($nums[$i]));
            } else {
                array_push($res, sprintf("%d->%d", $nums[$i], $nums[$j]));
                $i = $j;
            }
        }

        return $res;
    }
}