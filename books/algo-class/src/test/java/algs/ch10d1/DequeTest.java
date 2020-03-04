package algs.ch10d1;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class DequeTest {

    @Test
    void add_and_remove_first() {
        var deq = new Deque<String>(10);
        deq.addFirst("c");
        deq.addFirst("b");
        deq.addFirst("a");

        assertEquals("a", deq.removeFirst());
        assertEquals("b", deq.removeFirst());
        assertEquals("c", deq.removeFirst());
    }

    @Test
    void add_and_remove_last() {
        var deq = new Deque<String>(10);
        deq.addLast("a");
        deq.addLast("b");
        deq.addLast("c");

        assertEquals("c", deq.removeLast());
        assertEquals("b", deq.removeLast());
        assertEquals("a", deq.removeLast());
    }

}