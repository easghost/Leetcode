<?php

class MyHashSet {
    private array $list;
    
    function __construct() {
        $this->list = [];   
    }
  
    /**
     * @param Integer $key
     * @return NULL
     */
    function add($key) {
        $this->list[$key] = true;
    }
  
    /**
     * @param Integer $key
     * @return NULL
     */
    function remove($key) {
        $this->list[$key] = false;
    }
  
    /**
     * @param Integer $key
     * @return Boolean
     */
    function contains($key) {
        return $this->list[$key] ?? false;
    }
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * $obj = MyHashSet();
 * $obj->add($key);
 * $obj->remove($key);
 * $ret_3 = $obj->contains($key);
 */