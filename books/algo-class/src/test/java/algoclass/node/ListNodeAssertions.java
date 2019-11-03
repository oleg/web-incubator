package algoclass.node;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNull;

public class ListNodeAssertions {

    public static void assertList(ListNode first, int... expectedValues) {
        assertNull(first.prev(), "prev of the first should be null");

        ListNode last = assertExpectedAndGetLast(first, expectedValues);
        assertNull(last.next(), "next of the last should be null");

        ListNode againFirst = assertExpectedAndGetFirst(last, expectedValues);
        assertEquals(first, againFirst, "First that found should be the same");
    }

    private static ListNode assertExpectedAndGetLast(ListNode node, int[] expectedValues) {
        for (int i = 0; i < expectedValues.length; i++) {
            int value = node.value();
            assertEquals(expectedValues[i], value, String.format("wrong element at position [%d] (first to last)", i));

            if (i < expectedValues.length - 1) {
                node = node.next();
            }
        }
        return node;
    }

    private static ListNode assertExpectedAndGetFirst(ListNode node, int[] expectedValues) {
        for (int i = expectedValues.length - 1; i >= 0; i--) {
            int value = node.value();
            assertEquals(expectedValues[i], value, String.format("wrong element at position [%d] (last to first)", i));

            if (i > 0) {
                node = node.prev();
            }
        }
        return node;
    }

    public static ListNode makeNodes(int firstValue, int... restValues) { //todo rename class of move
        ListNode prev = new ListNode(firstValue, null, null);
        ListNode first = prev;
        for (int value : restValues) {
            ListNode next = new ListNode(value, prev, null);
            prev.next(next);
            prev = next;
        }
        return first;
    }

}
