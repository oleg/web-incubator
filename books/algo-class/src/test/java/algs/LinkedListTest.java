package algs;

import algs.LinkedList.Node;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

class LinkedListTest {

    @Test
    void test_toString() {
        var list = new LinkedList<Integer>();
        list.insert(1);
        list.insert(10);
        list.insert(100);

        assertEquals("[100, 10, 1]", list.toString());
    }

    @Test
    void create() {
        var list = new LinkedList<String>();
        list.insert("a");
        list.insert("b");

        assertEquals("[b, a]", list.toString());
    }

    @Test
    void search_not_found() {
        var list = new LinkedList<String>();
        list.insert("a");
        list.insert("b");

        Node<String> c = list.search("c");
        assertNull(c);
    }

    @Test
    void search_found() {
        var list = new LinkedList<String>();
        list.insert("a");
        list.insert("b");

        assertEquals("a", list.search("a").getValue());
        assertEquals("b", list.search("b").getValue());
    }

    @Test
    void delete_in_the_middle() {
        var list = new LinkedList<Integer>();
        list.insert(5);
        list.insert(4);
        list.insert(3);
        list.insert(2);
        list.insert(1);

        Node<Integer> e3 = list.search(3);
        list.delete(e3);

        assertEquals("[1, 2, 4, 5]", list.toString());
    }

    @Test
    void delete_in_the_start() {
        var list = new LinkedList<Integer>();
        list.insert(5);
        list.insert(4);
        list.insert(3);
        list.insert(2);
        list.insert(1);

        Node<Integer> e1 = list.search(1);
        list.delete(e1);

        assertEquals("[2, 3, 4, 5]", list.toString());
    }

    @Test
    void delete_in_the_end() {
        var list = new LinkedList<Integer>();
        list.insert(5);
        list.insert(4);
        list.insert(3);
        list.insert(2);
        list.insert(1);

        Node<Integer> e5 = list.search(5);
        list.delete(e5);

        assertEquals("[1, 2, 3, 4]", list.toString());
    }

}