<?php
// Runtime: 4 ms
// Memory Usage: 15.5 MB

class Solution {

    /**
     * @param String $address
     * @return String
     */
    function defangIPaddr($address) {
        return str_replace('.', '[.]', $address);
    }
}