// Runtime: 4 ms
// Memory Usage: 15.6 MB

<?php
class Solution {

    /**
     * @param Integer $num
     * @return Integer
     */
    function numberOfSteps($num) {
        if ($num === 0 || $num < 0) {
            return 0;
        }
        
        $counter = 0;
        while ($num !== 0) {
            if ($num % 2 !== 0) {
                $num--;
                $counter++;
            }
            if ($num > 0) {
                $num = $num / 2;
                $counter++;
            }
        }
        
        return $counter;
    }
}