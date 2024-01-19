package com.bondhan.java;

import java.io.*;

public class HashMap {

    public final int SIZE = 64;

    private Node[] Nodes;

    public HashMap() {
        Nodes = new Node[SIZE];
    }

    public HashMap(int initial) {
        Nodes = new Node[initial];
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

    public String Put(String key, String value) {
        int index = getIndex(key);

        Node n = new Node(key, value);

        if (Nodes[index] == null) {
            Nodes[index] = n;
            return value;
        }

        Node curr = Nodes[index];

        while (curr != null) {
            if (Objects.equals(curr.Key, key)) {
                curr.Value = value;
                return value;
            }

            if (curr.Next == null) {
                curr.Next = n;
                return value;
            }

            curr = curr.Next;
        }

        return value;
    }

    public String Get(String key) {
        int index = getIndex(key);

        Node curr = Nodes[index];

        while (curr != null) {

            if (Objects.equals(curr.Key, key)) {
                return curr.Value;
            }

            if (curr.Next == null) {
                return "";
            }

            curr = curr.Next;
        }

        return "";
    }

    public String Remove(String key) {
        int index = getIndex(key);

        Node curr = Nodes[index];
        Node prev = curr;

        while (curr != null) {

            if (Objects.equals(curr.Key, key)){
                if (prev == curr) {
                    String val = Nodes[index].Value;
                    Nodes[index] = null;
                    return val;
                }

                if (prev != curr) {
                    if (curr.Next != null) {
                        prev.Next = curr.Next;
                        return curr.Value;
                    }

                    prev.Next = null;
                    return curr.Value;
                }
            }

            if (curr.Next != null) {
                prev = curr;
                curr = curr.Next;
            }
        }

        return "";
    }
}