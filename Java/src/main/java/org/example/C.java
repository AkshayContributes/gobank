package org.example;

public class C{

    public static void main(String[] args) {
        A obj = new B();
        System.out.println(obj.d1);
        System.out.println(obj.d);
        obj.fun1();
        obj.fun();
        obj.sfun();
    }
}
