<?php

class Solution {

    /**
     * @param String $allowed
     * @param String[] $words
     * @return Integer
     */
    function countConsistentStrings($allowed, $words) {
        $count = count($words);
        $wordSet = array_fill(0, 26, 0);

        for ($i = 0; $i < strlen($allowed); $i++) {
            $wordSet[ord($allowed[$i]) - ord('a')]++;
        }

        foreach ($words as $word) {
            for ($i = 0; $i < strlen($word); $i++) {
                if ($wordSet[ord($word[$i]) - ord('a')] == 0) {
                    $count--;
                    break;
                }
            }
        }

        return $count;
    }
}