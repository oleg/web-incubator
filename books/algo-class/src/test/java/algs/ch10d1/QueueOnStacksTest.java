package algs.ch10d1;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class QueueOnStacksOnStacksTest {

    @Test
    void is_empty_after_creation() {
        var q = new QueueOnStacks<String>(3);
        assertTrue(q.isEmpty());
    }

    @Test
    void cant_dequeue_an_empty_queue() {
        var q = new QueueOnStacks<String>(3);
        assertThrows(IllegalStateException.class, q::dequeue);
    }

    @Test
    void not_empty_after_enqueue() {
        var q = new QueueOnStacks<String>(5);
        q.enqueue("a");
        assertFalse(q.isEmpty());
    }

    @Test
    void dequeue_returns_first_element_that_was_queued() {
        var q = new QueueOnStacks<String>(3);
        q.enqueue("a");
        q.enqueue("b");
        q.enqueue("c");

        assertEquals("a", q.dequeue());
        assertEquals("b", q.dequeue());
        assertEquals("c", q.dequeue());
    }

    @Test
    void is_full_and_not_empty() {
        var q = new QueueOnStacks<String>(2);
        q.enqueue("a");
        q.enqueue("b");

        assertTrue(q.isFull());
        assertFalse(q.isEmpty());
    }

    @Test
    void is_full_and_not_empty_size_5() {
        var q = new QueueOnStacks<String>(5);
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
        var q = new QueueOnStacks<String>(2);
        q.enqueue("a");
        q.enqueue("b");
        q.dequeue();
        q.enqueue("c");

        assertTrue(q.isFull());
        assertFalse(q.isEmpty());
    }

    @Test
    void is_empty_and_not_full() {
        var q = new QueueOnStacks<String>(2);
        assertTrue(q.isEmpty());
        assertFalse(q.isFull());
    }

    @Test
    void is_empty_and_not_full_after_dequeue() {
        var q = new QueueOnStacks<String>(2);
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
        var q = new QueueOnStacks<String>(3);
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

    @Test
    void many_elements() {
        var q = new QueueOnStacks<String>(7);
        q.enqueue("1");
        q.enqueue("2");
        q.enqueue("3");
        q.enqueue("4");
        q.enqueue("5");
        q.enqueue("6");
        q.enqueue("7");

        assertEquals("1", q.dequeue());
        assertEquals("2", q.dequeue());
        assertEquals("3", q.dequeue());
        assertEquals("4", q.dequeue());
        assertEquals("5", q.dequeue());
        assertEquals("6", q.dequeue());
        assertEquals("7", q.dequeue());
    }

    @Test
    void complex() {
        var q = new QueueOnStacks<String>(3);
        q.enqueue("1");
        assertEquals("1", q.dequeue());

        q.enqueue("2");
        q.enqueue("3");
        assertEquals("2", q.dequeue());
        assertEquals("3", q.dequeue());

        q.enqueue("4");
        assertEquals("4", q.dequeue());

        q.enqueue("5");
        q.enqueue("6");
        q.enqueue("7");
        assertEquals("5", q.dequeue());
        assertEquals("6", q.dequeue());
        q.enqueue("8");
        assertEquals("7", q.dequeue());
        assertEquals("8", q.dequeue());

    }

}