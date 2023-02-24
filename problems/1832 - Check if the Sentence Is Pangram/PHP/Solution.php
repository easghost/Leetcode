<?php

class Solution {

    /**
     * @param String $sentence
     * @return Boolean
     */
    function checkIfPangram($sentence) {
        $chars = str_split($sentence);
        $alpha = range('a', 'z');
        $res = array_diff($alpha, $chars);
        
        return !$res;
    }
}