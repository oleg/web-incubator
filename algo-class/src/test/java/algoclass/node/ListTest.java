package algoclass.node;

import org.junit.jupiter.api.Test;

import java.util.stream.IntStream;

import static algoclass.node.ListNodeAssertions.assertList;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class ListTest {

    @Test
    void cant_pop_element_after_creation() {
        List list = new List();
        assertThrows(IllegalStateException.class, () -> {
            list.pop();
        });
    }

    @Test
    void can_pop_one_element_after_push() {
        List list = new List();
        list.push(10);
        assertEquals(10, list.pop());
    }

    @Test
    void can_popLast_one_element_after_push() {
        List list = new List();
        list.push(10);
        assertEquals(10, list.popLast());
    }

    @Test
    void cant_pop_twice_after_one_push() {
        List list = new List();
        list.push(10);
        list.pop();
        assertThrows(IllegalStateException.class, () -> {
            list.pop();
        });
    }

    @Test
    void cant_popLast_twice_after_one_push() {
        List list = new List();
        list.push(10);
        list.popLast();
        assertThrows(IllegalStateException.class, () -> {
            list.popLast();
        });
    }

    @Test
    void cant_pop_and_popLast_after_one_push() {
        List list = new List();
        list.push(10);
        list.pop();
        assertThrows(IllegalStateException.class, () -> {
            list.popLast();
        });
    }

    @Test
    void cant_popLast_and_pop_after_one_push() {
        List list = new List();
        list.push(10);
        list.popLast();
        assertThrows(IllegalStateException.class, () -> {
            list.pop();
        });
    }

    ///
    @Test
    void can_pop_two_elements_after_two_pushes() {
        List list = new List();
        list.push(10);
        list.push(20);
        assertEquals(20, list.pop());
        assertEquals(10, list.pop());
    }

    @Test
    void can_pop_10_elements_after_10_elements_are_pushed() {
        List list = new List();
        IntStream.range(0, 10).forEach(list::push);
        IntStream.range(0, 10).forEach(ignore -> list.pop());
    }

    @Test
    void pop_returns_elements_up_to_the_first_one_in_reverse_order() {
        List list = new List();
        list.push(7);
        list.push(6);
        list.push(5);
        list.push(4);
        list.push(3);
        assertEquals(3, list.pop());
        assertEquals(4, list.pop());
        assertEquals(5, list.pop());
        assertEquals(6, list.pop());
        assertEquals(7, list.pop());
    }

    @Test
    void can_pop_once_after_pushLast() {
        List list = new List();
        list.pushLast(8);

        assertEquals(8, list.pop());
    }

    @Test
    void can_popLast_once_after_pushLast() {
        List list = new List();
        list.pushLast(9);

        assertEquals(9, list.popLast());
    }

    @Test
    void can_pop_three_times_after_pushLast() {
        List list = new List();
        list.pushLast(8);
        list.pushLast(9);
        list.pushLast(10);

        assertEquals(8, list.pop());
        assertEquals(9, list.pop());
        assertEquals(10, list.pop());
    }

    @Test
    void popLast_returns_last_element() {
        List list = new List();
        list.push(3);
        list.push(2);
        list.push(1);

        assertEquals(3, list.popLast());
    }

    @Test
    void repeated_popLast_returns_last_element_up_to_the_first() {
        List list = new List();
        list.push(3);
        list.push(2);
        list.push(1);

        assertEquals(3, list.popLast());
        assertEquals(2, list.popLast());
        assertEquals(1, list.popLast());
    }

    @Test
    void pushLast_adds_the_elements_to_the_end() {
        List list = new List();
        list.pushLast(1);
        list.pushLast(2);
        list.pushLast(3);

        assertEquals(1, list.pop());
        assertEquals(2, list.pop());
        assertEquals(3, list.pop());
    }

    @Test
    void can_pushLast_after_pop() {
        List list = new List();
        list.push(1);
        list.pop();
        list.pushLast(2);

        assertEquals(2, list.pop());
    }

    @Test
    void can_popLast_after_push() {
        List list = new List();
        list.push(1);
        list.push(2);
        list.push(3);

        assertEquals(1, list.popLast());
        assertEquals(2, list.popLast());
        assertEquals(3, list.popLast());
    }

}