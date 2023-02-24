<?php

class Solution {

    /**
     * @param String $path
     * @return String
     */
    function simplifyPath(String $path) {
        $s = explode('/',$path);
        array_shift($s);
        $res = [];
        foreach ($s as $val) {
            if ($val === "" || $val === ".") {
                continue;
            }
            if ($val ==="..") {
                if (count($res)) array_pop($res);
            } else {
                array_push($res,$val);
            }
        }
        
        return "/" . implode('/',$res);
    }
}