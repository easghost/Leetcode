<?php

/**
* @param Integer[] $nums
* @param Integer $val
* @return Integer
*/
function removeElement(&$nums, $val) {
    foreach ($nums as $key => $value) {
        if ($value == $val) {
            unset($nums[$key]);
        }
    }
    
    return count($nums);
}