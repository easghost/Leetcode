<?php

class Solution {

    /**
     * @param Int[] $nums
     * @param Int $target
     * @return Int[]
     */
    function twoSum($nums, $target) {
       for($i = 0; $i < count($nums); $i++)  {
          $num = $nums[$i];
          $diff = $target - $num;
              for($j = $i+1; $j < count($nums); $j++) {
                  if($nums[$j] == $diff) {
                      $res = [$i,$j];
                      break;
                  }
             }
		}
        
        return $res ?? [];
    }
}