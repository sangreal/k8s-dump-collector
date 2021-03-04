package com.smartnews.dump;

public class Main {
    public static void main(String args[]) throws Exception {
        try {
            Thread.sleep(10000);
        } catch (InterruptedException e) {
            System.out.println(e.getMessage());
        }
        Integer[] array = new Integer[10000 * 10000];
    }
}
