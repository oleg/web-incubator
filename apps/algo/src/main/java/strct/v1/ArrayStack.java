package strct.v1;

import java.lang.reflect.Array;

public class ArrayStack<T> {

    private T[] array;
    private int index;

    public ArrayStack(Class<T> clazz, int size) {
        @SuppressWarnings("unchecked")
        T[] ts = (T[]) Array.newInstance(clazz, size);
        this.array = ts;
    }

    public void push(T value) {
        if (index == array.length) {
            throw new IllegalStateException();
        }
        array[index++] = value;
    }

    public T pop() {
        if (index == 0) {
            throw new IllegalStateException();
        }
        return array[--index];
    }

}
