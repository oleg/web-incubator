package algs.ch10d1;

/**
 * 10.1-2
 */
class BiStack<T> {

    private final Object[] data;
    private int top;
    private BiStack<T> other;
    private boolean reversed;

    private BiStack(Object[] data, boolean reversed) {
        this.data = data;
        this.reversed = reversed;
        this.top = 0;
    }

    static <T> BiStack<T> create(int size) {
        Object[] data = new Object[size];

        BiStack<T> stack1 = new BiStack<>(data, false);
        BiStack<T> stack2 = new BiStack<>(data, true);
        stack1.other = stack2;
        stack2.other = stack1;

        return stack1;
    }

    BiStack<T> getOther() {
        return other;
    }

    boolean isEmpty() {
        return top == 0;
    }

    void push(T elem) {
        if (isFull()) {
            throw new IllegalStateException("overflow");
        }
        data[index()] = elem;
        top += 1;
    }

    private boolean isFull() {
        return top + other.top == data.length;
    }

    T pop() {
        if (isEmpty()) {
            throw new IllegalStateException("underflow");
        }
        top -= 1;
        @SuppressWarnings("unchecked")
        T elem = (T) data[index()];
        return elem;
    }

    private int index() {
        return reversed ? (data.length - 1) - top : top;
    }

}
