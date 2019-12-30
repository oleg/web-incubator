package algs;

import java.util.Arrays;

public class Deque<T> {

    private Object[] data;
    private int first;
    private int last;

    public Deque(int size) {
        data = new Object[size];
        first = 0;
        last = 0;
    }

    public void addFirst(T elem) {
        first = prev(first);
        data[first] = elem;
    }

    public T removeFirst() {
        T elem = (T) data[first];
        first = next(first);
        return elem;
    }

    public void addLast(T elem) {
        data[last] = elem;
        last = next(last);
    }

    public T removeLast() {
        last = prev(last);
        T elem = (T) data[last];
        return elem;
    }

    private int next(int index) {
        if ((++index) < data.length) {
            return index;
        }
        return 0;
    }

    private int prev(int index) {
        if ((--index) >= 0) {
            return index;
        }
        return data.length - 1;
    }

    @Override
    public String toString() {
        return "((" + first + ")" + Arrays.toString(data) + "(" + last + "))";
    }
}
