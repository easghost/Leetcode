<?php
// Runtime: 35 ms
// Memory Usage: 19.4 MB

/**
 * 
 * 
 * Definition for a Node.
 * class Node {
 *     function __construct($val = 0) {
 *         $this->val = $val;
 *         $this->left = null;
 *         $this->right = null;
 *         $this->next = null;
 *     }
 * }
 */

class Solution {
    /**
     * @param Node $root
     * @return Node
     */
    public function connect($root) {
        if ($root === null || $root->left === null) {
            return $root;
        }
        
        $root->left->next = $root->right;
        if ($root->next !== null) {
            $root->right->next = $root->next->left;
        }
        
        $root->left = $this->connect($root->left);
        $root->right = $this->connect($root->right);
        
        return $root;
    }
}