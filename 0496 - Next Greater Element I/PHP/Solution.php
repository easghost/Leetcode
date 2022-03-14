<?php

class Solution {

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Integer[]
     */
    function nextGreaterElement($nums1, $nums2) {
        $res = $check = [];
        for ($i = 0; $i < sizeof($nums1); $i++) {
            $check[$nums1[$i]] = $i;
            $res[$i] = -1;
        }

        $stack = []; 
        for ($i = 0; $i < sizeof($nums2); $i++) {
            while (!empty($stack) && $nums2[$i] > $stack[sizeof($stack)-1]) {
                $num = array_pop($stack);
                $res[$check[$num]] = $nums2[$i];
            }; 

            if ($check[$nums2[$i]] || $check[$nums2[$i]] === 0) {
                array_push($stack, $nums2[$i]);
            }
        };
        
        return $res;
    }
}