package algoclass.node;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class ListSortTest {

    @Test
    void sort_empty_doesnt_produce_errors() {
        List list = new List();

        list.sort();
    }

    @Test
    void sort_one_elements_works() {
        List list = new List();
        list.push(1);

        list.sort();

        assertEquals(1, list.pop());
    }

    @Test
    void sort_two_elements_works() {
        List list = new List();
        list.push(1);
        list.push(3);

        list.sort();

        assertEquals(1, list.pop());
        assertEquals(3, list.pop());
    }

    @Test
    void sort_three_elements_works() {
        List list = new List();
        list.push(2);
        list.push(1);
        list.push(3);

        list.sort();

        assertEquals(1, list.pop());
        assertEquals(2, list.pop());
        assertEquals(3, list.pop());
    }

    @Test
    void sort_four_elements_works() {
        List list = new List();
        list.push(2);
        list.push(1);
        list.push(4);
        list.push(3);

        list.sort();

        assertEquals(1, list.pop());
        assertEquals(2, list.pop());
        assertEquals(3, list.pop());
        assertEquals(4, list.pop());
    }

}