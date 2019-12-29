package algs;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class BiStackTest {

    @Test
    void create() {
        var stack1 = BiStack.<Integer>create(10);
        BiStack<Integer> stack2 = stack1.getOther();

        assertNotEquals(stack1, stack2);
        assertEquals(stack1, stack2.getOther());
        assertEquals(stack1.getOther(), stack2);
    }

    @Test
    void empty() {
        var stack1 = BiStack.<Integer>create(10);
        BiStack<Integer> stack2 = stack1.getOther();

        assertTrue(stack1.isEmpty());
        assertTrue(stack2.isEmpty());
    }

    @Test
    void push_first() {
        var stack1 = BiStack.<Integer>create(10);
        BiStack<Integer> stack2 = stack1.getOther();

        stack1.push(7);

        assertFalse(stack1.isEmpty());
        assertTrue(stack2.isEmpty());
    }

    @Test
    void push_second() {
        var stack1 = BiStack.<Integer>create(10);
        BiStack<Integer> stack2 = stack1.getOther();

        stack1.push(7);
        stack2.push(9);

        assertFalse(stack1.isEmpty());
        assertFalse(stack2.isEmpty());
    }

    @Test
    void pop_first() {
        var stack1 = BiStack.<Integer>create(10);
        BiStack<Integer> stack2 = stack1.getOther();

        stack1.push(7);
        stack2.push(9);

        assertEquals(7, stack1.pop());
        assertEquals(9, stack2.pop());
    }

    @Test
    void multiple_push_and_pops() {
        var stack1 = BiStack.<String>create(10);
        BiStack<String> stack2 = stack1.getOther();

        stack1.push("a1");
        stack1.push("a2");
        stack1.push("a3");
        stack1.push("a4");
        stack1.push("a5");

        stack2.push("b1");
        stack2.push("b2");
        stack2.push("b3");
        stack2.push("b4");
        stack2.push("b5");

        assertEquals("a5", stack1.pop());
        assertEquals("a4", stack1.pop());
        assertEquals("a3", stack1.pop());
        assertEquals("a2", stack1.pop());
        assertEquals("a1", stack1.pop());

        assertEquals("b5", stack2.pop());
        assertEquals("b4", stack2.pop());
        assertEquals("b3", stack2.pop());
        assertEquals("b2", stack2.pop());
        assertEquals("b1", stack2.pop());
    }

    @Test
    void stack_underflow() {
        var stack1 = BiStack.<String>create(4);

        assertThrows(IllegalStateException.class, () -> stack1.pop());
        assertThrows(IllegalStateException.class, () -> stack1.getOther().pop());
    }

    @Test
    void stack_overflow() {
        var stack1 = BiStack.<String>create(4);
        stack1.push("a1");
        stack1.push("a2");

        BiStack<String> stack2 = stack1.getOther();
        stack2.push("b1");
        stack2.push("b2");

        System.out.println(stack1);

        assertThrows(IllegalStateException.class, () -> stack1.push("a3"));
        assertThrows(IllegalStateException.class, () -> stack2.push("b3"));
    }
}