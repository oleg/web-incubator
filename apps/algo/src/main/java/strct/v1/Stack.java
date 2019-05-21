package strct.v1;

public class Stack<T> {
    private Node<T> head;

    public void push(T value) {
        head = new Node<>(value, head);
    }

    public T pop() {
        if (head == null) {
            throw new IllegalStateException("Stack is empty");
        }
        T value = head.value;
        head = head.next;
        return value;
    }

    private class Node<T> {
        private T value;
        private Node<T> next;

        public Node(T value, Node<T> next) {
            this.value = value;
            this.next = next;
        }
    }
}
