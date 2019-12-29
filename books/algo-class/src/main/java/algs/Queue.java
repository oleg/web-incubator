package algs;

public class Queue<T> {

    private Object[] data;
    private int head;
    private int tail;

    public Queue(int size) {
        data = new Object[size + 1];
        head = 0;
        tail = 0;
    }

    public boolean isEmpty() {
        return head == tail;
    }

    public boolean isFull() {
        return next(tail) == head;
    }


    public T dequeue() {
        if (isEmpty()) {
            throw new IllegalStateException("underflow");
        }
        @SuppressWarnings("unchecked")
        T elem = (T) data[head];
        head = next(head);
        return elem;
    }

    public void enqueue(T elem) {
        if (isFull()) {
            throw new IllegalStateException("overflow");
        }
        data[tail] = elem;
        tail = next(tail);
    }

    private int next(int index) {
        if ((++index) != data.length) {
            return index;
        }
        return 0;
    }
}
