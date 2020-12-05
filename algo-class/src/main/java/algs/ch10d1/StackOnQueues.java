package algs.ch10d1;

/**
 * 10.1-7
 */
public class StackOnQueues<T> {

    private Queue<T> main;
    private Queue<T> temp;

    StackOnQueues(int size) {
        main = new Queue<>(size);
        temp = new Queue<>(size);
    }

    boolean isEmpty() {
        return main.isEmpty();
    }

    boolean isFull() {
        return main.isFull();
    }

    public void push(T elem) {
        main.enqueue(elem);
    }

    public T pop() {
        T elem;
        for (elem = main.dequeue(); !main.isEmpty(); elem = main.dequeue()) {
            temp.enqueue(elem);
        }
        swapMainAndTemp();
        return elem;
    }

    private void swapMainAndTemp() {
        Queue<T> tmp = main;
        main = temp;
        temp = tmp;
    }
}
