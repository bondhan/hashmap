package com.bondhan.java;

import java.io.*;

public class HashMap {

    public final int SIZE = 64;

    private Node[] Nodes;

    public HashMap() {
        Nodes = new Node[SIZE];
    }

    public int hash(String key) {
        int total = 0;
        for (int i = 0; i < key.length(); i++) {
            total += key.codePointAt(i);
        }

        return total;
    }

    public int getIndex(String key) {
        int h = hash(key);

        return h & (SIZE - 1);
    }

    public void Put(String key) {

    }
}