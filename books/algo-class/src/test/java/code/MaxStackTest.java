package code;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MaxStackTest {

    @Test
    void push_and_pop() {
        MaxStack stack = new MaxStack();
        stack.push(100);
        stack.push(200);
        stack.push(300);

        assertEquals(300, stack.pop());
        assertEquals(200, stack.pop());
        assertEquals(100, stack.pop());
    }

    @Test
    void push_and_pop_multiple() {
        MaxStack stack = new MaxStack();

        stack.push(100);
        stack.push(200);
        assertEquals(200, stack.pop());
        assertEquals(100, stack.pop());

        stack.push(300);
        stack.push(400);
        assertEquals(400, stack.pop());
        assertEquals(300, stack.pop());

        stack.push(500);
        assertEquals(500, stack.pop());
    }

    @Test
    void top() {
        MaxStack stack = new MaxStack();
        stack.push(100);
        stack.push(200);
        stack.push(300);

        assertEquals(300, stack.top());
        stack.pop();
        assertEquals(200, stack.top());
        stack.pop();
        assertEquals(100, stack.pop());
    }

    @Test
    void peekMax() {
        MaxStack stack = new MaxStack();
        stack.push(3);
        stack.push(1);
        stack.push(2);
        stack.push(5);
        stack.push(1);

        assertEquals(5, stack.peekMax());
        stack.pop();
        stack.pop();
        assertEquals(3, stack.peekMax());
    }

    @Test
    void complex_test() {
        MaxStack stack = new MaxStack();
        stack.push(5);
        stack.push(1);
        stack.push(5);

        assertEquals(5, stack.top());
        assertEquals(5, stack.popMax());
        assertEquals(1, stack.top());
        assertEquals(5, stack.peekMax());
        assertEquals(1, stack.pop());
        assertEquals(5, stack.top());
    }

    @Test
    void push_bigger() {
        MaxStack stack = new MaxStack();
        stack.push(2);
        stack.push(8);

    }

    @Test
    void push_smaller() {
        MaxStack stack = new MaxStack();
        stack.push(8);
        stack.push(2);

    }
}