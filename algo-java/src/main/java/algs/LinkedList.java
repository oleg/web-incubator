package algs;

import java.util.Objects;

public class LinkedList<T> {
    public static class Node<T> {
        private Node<T> prev;
        private Node<T> next;

        private final T value;

        public Node(T value) {
            this.value = value;
        }

        public T getValue() {
            return value;
        }

        @Override
        public String toString() {
            String p = prev == null ? "" : "->";
            String n = next == null ? "" : "->";
            return "(" + p + value + n + ")";
        }
    }

    private Node<T> head;

    public void insert(T elem) {
        Node<T> node = new Node<>(elem);
        node.next = head;
        if (head != null) {
            head.prev = node;
        }
        head = node;
    }

    public void delete(Node<T> node) {
        //todo check head == null
        //todo check node == null
        Node<T> prev = node.prev;
        Node<T> next = node.next;
        if (prev != null) {
            prev.next = next;
        } else {
            head = next;
        }
        if (next != null) {
            next.prev = prev;
        }
    }

    public Node<T> search(T elem) {
        Node<T> tmp = head;
        while (tmp != null && !Objects.equals(tmp.getValue(), elem)) {
            tmp = tmp.next;
        }
        return tmp;
    }

    @Override
    public String toString() {
        Node<T> tmp = this.head;
        StringBuilder st = new StringBuilder("[");
        while (tmp != null) {
            st.append(tmp.getValue());
            st.append(", ");
            tmp = tmp.next;
        }
        st.delete(st.length() - 2, st.length());
        st.append("]");
        return st.toString();
    }
}
