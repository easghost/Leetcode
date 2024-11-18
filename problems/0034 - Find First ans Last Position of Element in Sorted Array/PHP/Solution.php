<?php

class Solution {

/**
 * @param Integer[] $nums
 * @param Integer $target
 * @return Integer[]
 */
function searchRange($nums, $target) {
    $first = -1;
    $last = -1;

    foreach ($nums as $key => $value) {
        if ($value === $target) {
            if ($first === -1) {
                $first = $key;
            }

            $last = $key;
        }
    }

    return [$first, $last];
}
}