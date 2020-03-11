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
        T item = null;
        while (!main.isEmpty()) {
            item = main.dequeue();

            if (!main.isEmpty()) { //without it ????
                temp.enqueue(item);
            }
        }
        System.out.println(temp);

        Queue<T> tmp = main;
        main = temp;
        temp = tmp;
        return item;
    }
}
