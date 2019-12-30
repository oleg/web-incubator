package algs;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class QueueTest {

    @Test
    void is_empty_after_creation() {
        var q = new Queue<String>(3);
        assertTrue(q.isEmpty());
    }

    @Test
    void cant_dequeue_an_empty_queue() {
        var q = new Queue<String>(3);
        assertThrows(IllegalStateException.class, q::dequeue);
    }

    @Test
    void not_empty_after_enqueue() {
        var q = new Queue<String>(5);
        q.enqueue("a");
        assertFalse(q.isEmpty());
    }

    @Test
    void dequeue_returns_first_element_that_was_queued() {
        var q = new Queue<String>(3);
        q.enqueue("a");
        q.enqueue("b");
        q.enqueue("c");

        assertEquals("a", q.dequeue());
        assertEquals("b", q.dequeue());
        assertEquals("c", q.dequeue());
    }

    @Test
    void is_full_and_not_empty() {
        var q = new Queue<String>(2);
        q.enqueue("a");
        q.enqueue("b");

        assertTrue(q.isFull());
        assertFalse(q.isEmpty());
    }

    @Test
    void is_full_and_not_empty_size_5() {
        var q = new Queue<String>(5);
        q.enqueue("a");
        q.enqueue("b");
        q.enqueue("c");
        q.enqueue("d");
        q.enqueue("e");

        assertTrue(q.isFull());
        assertFalse(q.isEmpty());
    }

    @Test
    void is_full_and_not_empty_after_dequeue() {
        var q = new Queue<String>(2);
        q.enqueue("a");
        q.enqueue("b");
        q.dequeue();
        q.enqueue("c");

        assertTrue(q.isFull());
        assertFalse(q.isEmpty());
    }
    @Test
    void is_empty_and_not_full() {
        var q = new Queue<String>(2);
        assertTrue(q.isEmpty());
        assertFalse(q.isFull());
    }

    @Test
    void is_empty_and_not_full_after_dequeue() {
        var q = new Queue<String>(2);
        q.enqueue("a");
        q.enqueue("b");
        q.dequeue();
        q.enqueue("c");
        q.dequeue();
        q.dequeue();

        assertTrue(q.isEmpty());
        assertFalse(q.isFull());
    }

    @Test
    void queue_acts_like_a_circle() {
        var q = new Queue<String>(3);
        q.enqueue("a");
        q.enqueue("b");
        q.enqueue("c");

        q.dequeue();//a
        q.dequeue();//b

        q.enqueue("d");
        q.enqueue("e");

        assertEquals("c", q.dequeue());
        assertEquals("d", q.dequeue());
        assertEquals("e", q.dequeue());
    }

}