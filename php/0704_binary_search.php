// Runtime: 80 ms
// Memory Usage: 17.1 MB

class Solution {

    /**
    * @param Integer[] $nums
    * @param Integer $target
    * @return Integer
    */
    function search($nums, $target) {
        if (count($nums) === 0) {
            return 0;
        }
        foreach($nums as $key => $value) {
            if ($value === $target) {
                return $key;
            }
        }
        return -1;
    }
}