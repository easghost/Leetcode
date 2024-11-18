<?php

class MyHashMap {
    private array $list;

    /**
     */
    function __construct() {
        $this->list = [];
    }
  
    /**
     * @param Integer $key
     * @param Integer $value
     * @return NULL
     */
    function put($key, $value) {
        $this->list[$key] = $value;
    }
  
    /**
     * @param Integer $key
     * @return Integer
     */
    function get($key) {
        return $this->list[$key] ?? -1;
    }
  
    /**
     * @param Integer $key
     * @return NULL
     */
    function remove($key) {
        unset($this->list[$key]);
    }
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * $obj = MyHashMap();
 * $obj->put($key, $value);
 * $ret_2 = $obj->get($key);
 * $obj->remove($key);
 */