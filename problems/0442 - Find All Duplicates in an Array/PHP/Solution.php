<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function findDuplicates($nums) {
        $res = [];
        if (count($nums) === 1) {
            return $res;
        }

        foreach ($nums as $num) {
            $num = (int)$num;
            if ($nums[$num-1] > 0) {
                $nums[$num-1] = $nums[$num-1]*-1;
            } else {
                $res[] = $num;
            }
        }

        return $res;
    }
}