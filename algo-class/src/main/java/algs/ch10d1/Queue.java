package algs.ch10d1;

/**
 * 10.1-4
 */
class Queue<T> {

    private Object[] data;
    private int head;
    private int tail;

    Queue(int size) {
        data = new Object[size + 1];
        head = 0;
        tail = 0;
    }

    boolean isEmpty() {
        return head == tail;
    }

    boolean isFull() {
        return next(tail) == head;
    }


    T dequeue() {
        if (isEmpty()) {
            throw new IllegalStateException("underflow");
        }
        @SuppressWarnings("unchecked")
        T elem = (T) data[head];
        head = next(head);
        return elem;
    }

    void enqueue(T elem) {
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
