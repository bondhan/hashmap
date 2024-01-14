package com.bondhan.test;

import com.bondhan.java.*;

import static org.junit.jupiter.api.Assertions.assertEquals;

class HashMapTest {

    HashMap hm;

    @org.junit.jupiter.api.BeforeEach
    void setup() {
        hm = new HashMap();
    }

    @org.junit.jupiter.api.Test
    void testHash() {
        assertEquals(0x61, hm.hash("a"));
        assertEquals(0x41, hm.hash("A"));
    }

    @org.junit.jupiter.api.Test
    void testGetIndex() {
        assertEquals(0, hm.getIndex(""));
        assertEquals(0x61 & (hm.SIZE - 1), hm.getIndex("a"));
        assertEquals(0x41 & (hm.SIZE - 1), hm.getIndex("A"));

    }

    @org.junit.jupiter.api.Test
    void testPutGet() {
        assertEquals(hm.Put("a", "a"), hm.Get("a"));
        assertEquals(hm.Put("a", "b"), hm.Get("a"));
        assertEquals(hm.Put("aB", "aB"), hm.Get("aB"));
        assertEquals(hm.Put("aB", "aB"), hm.Get("aB"));

        hm.Put("abba", "1");
        hm.Put("abab", "2");
        hm.Put("aabb", "3");

        assertEquals("3", hm.Get("aabb"));
        assertEquals("2", hm.Get("abab"));
        assertEquals("1", hm.Get("abba"));
    }

    @org.junit.jupiter.api.Test
    void testRemove(){
        hm.Put("a", "1");
        String ret = hm.Remove("a");
        assertEquals("1", ret);

        assertEquals("", hm.Get("a"));

        hm.Put("abba", "1");
        hm.Put("abab", "2");
        hm.Put("aabb", "3");

        hm.Remove("aabb");
        hm.Remove("abab");
        hm.Remove("abba");

        assertEquals("", hm.Get("abba"));
        assertEquals("", hm.Get("abab"));
        assertEquals("", hm.Get("aabb"));
    }
}