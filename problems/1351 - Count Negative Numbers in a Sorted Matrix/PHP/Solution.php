<?php

class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function countNegatives($grid) {
        $m = count($grid);
        $n = count($grid[0]);
        $res = 0;

        for ($i = $m - 1; $i >= 0; $i--) {
            sort($grid[$i]);

            for ($j = 0; $j <= $n; $j++) {
                if ($grid[$i][$n-1] < 0) {
                    $res = $res + $n;
                    break;
                }

                if ($grid[$i][$j] < 0) {
                    $res++;
                } else {
                    break;
                }
            }
        }

        return $res;
    }
}