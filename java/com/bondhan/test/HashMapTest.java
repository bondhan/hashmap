package com.bondhan.test;

import com.bondhan.java.*;

import static org.junit.jupiter.api.Assertions.assertEquals;

class HashMapTest {

    static HashMap hm;

    @org.junit.jupiter.api.BeforeAll
    static void setup() {
        hm = new HashMap();
    }

    @org.junit.jupiter.api.Test
    void testHash() {
        assertEquals(0x61, hm.hash("a"));
        assertEquals(0x41, hm.hash("A"));
    }

    @org.junit.jupiter.api.Test
    void testGetIndex(){
        assertEquals(0, hm.getIndex(""));
        assertEquals(0x61 & (hm.SIZE - 1), hm.getIndex("a"));
        assertEquals(0x41 & (hm.SIZE - 1), hm.getIndex("A"));

    }

    @org.junit.jupiter.api.Test
    void put() {
    }
}