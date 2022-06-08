<?php

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function removePalindromeSub($s) {
        if (strlen($s) === 0) {
            return 0;
        }
        $l = 0;
        $r = strlen($s) - 1;
        while ($l < $r) {
            if ($s[$l] === $s[$r]) {
                $l++;
                $r--;
            } else {
                break;
            }
        }
        if ($l === $r || $l > $r) {
            return 1;
        }
        return 2;
    }
}