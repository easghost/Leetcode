<?php

class Solution {

    /**
     * @param Integer[][] $coordinates
     * @return Boolean
     */
    function checkStraightLine($coordinates) {
        $x0 = $coordinates[0][0];
        $x1 = $coordinates[1][0];
        $y0 = $coordinates[0][1];
        $y1 = $coordinates[1][1];
        
        for ($i = 1; $i < count($coordinates); $i++) {
            $x2 = $coordinates[$i][0];
            $y2 = $coordinates[$i][1];
            
            if (($x1 - $x0) * ($y2 - $y1) !== ($y1 - $y0) * ($x2 - $x1)) {
                return false;
            }
        }
        
        return true;
    }
}