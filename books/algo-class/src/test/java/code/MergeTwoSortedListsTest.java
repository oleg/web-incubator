package code;

import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

class MergeTwoSortedListsTest {

    @Test
    void factory_produce_correct_list() {
        assertEquals("N(1, N(2, N(3, N(4, null))))", factory(1, 2, 3, 4).toString());
        assertEquals("N(8, N(7, N(6, N(5, null))))", factory(8, 7, 6, 5).toString());
    }

    @Test
    void merge_works() {
        ListNode l1 = factory(1, 2, 3);
        ListNode l2 = factory(4, 5, 6);

        ListNode m = new MergeTwoSortedLists().mergeTwoLists(l1, l2);
        assertEquals("N(1, N(2, N(3, N(4, N(5, N(6, null))))))", m.toString());
    }

    @Test
    void merge_works_single_first() {
        ListNode l1 = factory(1, 3, 4);
        ListNode l2 = factory(2);

        ListNode m = new MergeTwoSortedLists().mergeTwoLists(l1, l2);
        assertEquals("N(1, N(2, N(3, N(4, null))))", m.toString());
    }

    @Test
    void merge_works_single_second() {
        ListNode l1 = factory(1);
        ListNode l2 = factory(2, 3, 4);

        ListNode m = new MergeTwoSortedLists().mergeTwoLists(l1, l2);
        assertEquals("N(1, N(2, N(3, N(4, null))))", m.toString());
    }

    @Test
    void returns_first_if_second_is_null() {
        ListNode l1 = factory(1, 6);
        ListNode m = new MergeTwoSortedLists().mergeTwoLists(l1, null);
        assertEquals("N(1, N(6, null))", m.toString());
    }

    @Test
    void returns_second_if_first_is_null() {
        ListNode l2 = factory(77, 99);
        ListNode m = new MergeTwoSortedLists().mergeTwoLists(null, l2);
        assertEquals("N(77, N(99, null))", m.toString());
    }
    
    @Test
    void nulls() {
        ListNode m = new MergeTwoSortedLists().mergeTwoLists(null, null);
        assertNull(m);
    }

    private static ListNode factory(int x, int... xs) {
        ListNode f = new ListNode(x);
        ListNode e = Arrays.stream(xs)
                .mapToObj(ListNode::new)
                .reduce(f, (a, b) -> {
                    a.next = b;
                    return b;
                });
        return f;
    }
}