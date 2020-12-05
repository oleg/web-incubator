package algs.ch10d1;

/**
 * 10.1-6
 */
public class QueueOnStacks<T> {

    private Stack<T> main;
    private Stack<T> temp;


    QueueOnStacks(int size) {
        main = new Stack<>(size);
        temp = new Stack<>(size);
    }

    boolean isEmpty() {
        return main.isEmpty();
    }

    boolean isFull() {
        return main.isFull();
    }

    T dequeue() {
        T elem;
        for (elem = main.pop(); !main.isEmpty(); elem = main.pop()) {
            temp.push(elem);
        }
        while (!temp.isEmpty()) {
            main.push(temp.pop());
        }

        return elem;
    }

    void enqueue(T elem) {
        main.push(elem);
    }

}
