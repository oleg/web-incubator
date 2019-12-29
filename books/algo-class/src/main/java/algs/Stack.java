package algs;


class Stack<T> {

    private Object[] data;
    private int top = 0;

    public Stack(int size) {
        data = new Object[size];
    }

    boolean isEmpty() {
        return top == 0;
    }

    boolean isFull() {
        return top == data.length;
    }

    public void push(T elem) {
        if (isFull()) {
            throw new IllegalStateException("overflow");
        }
        data[top] = elem;
        top += 1;
    }

    public T pop() {
        if (isEmpty()) {
            throw new IllegalStateException("underflow");
        }
        top -= 1;
        @SuppressWarnings("unchecked")
        T elem = (T) data[top];
        return elem;
    }
}




