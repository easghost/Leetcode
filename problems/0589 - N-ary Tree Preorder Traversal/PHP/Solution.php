<?php

/**
 * Definition for a Node.
 * class Node {
 *     public $val = null;
 *     public $children = null;
 *     function __construct($val = 0) {
 *         $this->val = $val;
 *         $this->children = array();
 *     }
 * }
 */

class Solution {
    /**
     * @param Node $root
     * @return integer[]
     */
    function preorder($root) {
        $result = [];
        $nodesQueue = [$root];
        
        while ($node = array_pop($nodesQueue)) {
            $result[] = $node->val;

            if (empty($node->children)) {
                continue;
            }

            for ($i = count($node->children) - 1; $i >= 0; $i--) {
                $nodesQueue[] = $node->children[$i];
            }
        }
        
        return $result;
    }
}