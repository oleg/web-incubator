package algoclass.node;

import org.junit.jupiter.api.Test;

import static algoclass.node.ListNodeAssertions.assertList;
import static algoclass.node.ListNodeAssertions.makeNodes;
import static org.junit.jupiter.api.Assertions.*;

class ListNodeTest {

    @Test
    void prepend_to_first() {
        ListNode a = makeNodes(10);
        ListNode b = makeNodes(20);

        a.prependOne(b);

        assertList(b, 20, 10);
    }

    @Test
    void prepended_item_must_not_have_prev_elements() {
        ListNode a = makeNodes(10);
        ListNode b1 = makeNodes(20, 30);
        ListNode b2 = b1.next();

        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                () -> a.prependOne(b2));
        assertEquals("Must not have prev", exception.getMessage());
    }

    @Test
    void prepend_to_second() {
        ListNode a1 = makeNodes(10, 30);
        ListNode a2 = a1.next();
        ListNode b = makeNodes(20);

        a2.prependOne(b);

        assertList(a1, 10, 20, 30);
    }

    @Test
    void prepend_to_second_that_has_next() {
        ListNode a1 = makeNodes(10, 20, 30, 40);
        ListNode a2 = a1.next();
        ListNode a3 = a2.next();
        ListNode b = makeNodes(7);

        a3.prependOne(b);

        assertList(a1, 10, 20, 7, 30, 40);
    }

    @Test
    void prepend_element_that_has_next() {
        ListNode a1 = makeNodes(10, 20);
        ListNode a2 = a1.next();
        ListNode b = makeNodes(7, 8);

        a2.prependOne(b);

        assertList(a1, 10, 7, 20);
    }

    @Test
    void returns_rest_of_prepended_element() {
        ListNode a1 = makeNodes(10, 20);
        ListNode a2 = a1.next();
        ListNode b1 = makeNodes(7, 8, 9);

        ListNode b2 = a2.prependOne(b1);

        assertList(b2, 8, 9);
    }

    @Test
    void returns_null_if_there_is_no_other_elements() {
        ListNode a1 = makeNodes(10);
        ListNode b1 = makeNodes(7);

        ListNode b2 = a1.prependOne(b1);

        assertNull(b2);
    }

    @Test
    void should_not_append_element_with_prev() {
        ListNode a1 = makeNodes(10);
        ListNode b1 = makeNodes(7, 8);
        ListNode b2 = b1.next();

        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                () -> a1.appendOne(b2));
        assertEquals("Must not have prev", exception.getMessage());
    }

    @Test
    void append_adds_element_to_the_end() {
        ListNode a1 = makeNodes(10);
        ListNode b1 = makeNodes(7);

        a1.appendOne(b1);

        assertList(a1, 10, 7);
    }

    @Test
    void append_leaves_original_tail() {
        ListNode a1 = makeNodes(10, 20, 30);
        ListNode b1 = makeNodes(6, 7, 8);

        a1.appendOne(b1);

        assertList(a1, 10, 6, 20, 30);
    }

    @Test
    void append_returns_not_appended_tail() {
        ListNode a1 = makeNodes(10, 20, 30);
        ListNode b1 = makeNodes(6, 7, 8);

        ListNode b2 = a1.appendOne(b1);

        assertList(b2, 7, 8);
    }

    @Test
    void appendTail_should_add_complete_tail() {
        ListNode a1 = makeNodes(10, 20, 30);
        ListNode a3 = a1.next().next();
        ListNode b1 = makeNodes(6, 7, 8);

        a3.replaceTail(b1);

        assertList(a1, 10, 20, 30, 6, 7, 8);
    }


    @Test
    void appendTail_returns_old_tail() {
        ListNode a1 = makeNodes(10, 20, 30, 40);
        ListNode a2 = a1.next();
        ListNode b1 = makeNodes(6, 7, 8);

        ListNode a3 = a2.replaceTail(b1);
        assertList(a3, 30, 40);
    }


}