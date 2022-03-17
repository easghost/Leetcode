<?php

class Solution {

/**
 * @param String $word1
 * @param String $word2
 * @return String
 */
function mergeAlternately($word1, $word2) {
    $res = "";
    $max = max(strlen($word1), strlen($word2));
    
    for ($i = 0; $i < $max; $i++) {
        if ($i < strlen($word1)) {
            $res .= $word1[$i];
        }
        
        if ($i < strlen($word2)) {
            $res .= $word2[$i];
        }
    }
    
    return $res;
}
}