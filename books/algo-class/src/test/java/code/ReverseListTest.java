package code;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

class ReverseListTest {
    @Test
    void empty() {
        ListNode lr = new ReverseList().reverseList(null);
        assertNull(lr);
    }

    @Test
    void one() {
        ListNode lr = new ReverseList().reverseList(new ListNode(1));
        assertEquals(1, lr.val);
        assertNull(lr.next);
    }

    @Test
    void two() {
        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);

        ListNode lr = new ReverseList().reverseList(l1);

        assertEquals(2, lr.val);
        assertEquals(1, lr.next.val);
        assertNull(lr.next.next);
    }

    @Test
    void five() {
        ListNode l1 = new ListNode(1);
        ListNode l2 = new ListNode(2);
        ListNode l3 = new ListNode(3);
        ListNode l4 = new ListNode(4);
        l1.next = l2;
        l2.next = l3;
        l3.next = l4;

        ListNode lr = new ReverseList().reverseList(l1);
        assertEquals(4, lr.val);
        assertEquals(3, lr.next.val);
        assertEquals(2, lr.next.next.val);
        assertEquals(1, lr.next.next.next.val);
        assertNull(lr.next.next.next.next);
    }


}