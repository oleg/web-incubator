package lcode;

public class MaxStack {

    ListNode head;
    ListNode max;

    public void push(int value) {
        ListNode node1 = new ListNode(value);
        ListNode node2 = new ListNode(value);
        node1.sibling = node2;
        node2.sibling = node1;
        pushToHead(node1);
        pushToMax(node2);
    }

    private void pushToHead(ListNode node) {
        if (head != null) {
            head.prev = node;
        }
        node.prev = null;
        node.next = head;
        head = node;
    }

    public void pushToMax(ListNode node) {
        if (max == null) {
            max = node;
        } else {
            ListNode prev = findGreater(node);

            ListNode tmp;
            if (prev == null) {
                tmp = max;
                max = node;
            } else {
                tmp = prev.next;
                prev.next = node;
            }
            node.next = tmp;

            node.prev = prev;
            if (tmp != null) {
                tmp.prev = node;
            }
        }
    }

    private ListNode findGreater(ListNode node) {
        ListNode prev = null;
        ListNode tmp = max;
        while (tmp != null && tmp.value > node.value) {
            prev = tmp;
            tmp = tmp.next;
        }
        return prev;
    }

    public int pop() {
        int value = head.value;
        ListNode node = head.sibling;
        head = head.next;
        if (head != null) {
            head.prev = null;
        }
        max = insertNode(node, max);
        return value;
    }

    public int popMax() {
        int value = max.value;
        ListNode node = max.sibling;
        max = max.next;
        if (max != null) {
            max.prev = null;
        }
        head = insertNode(node, head);
        return value;
    }

    public int top() {
        return head.value;
    }

    public int peekMax() {
        return max.value;
    }

    private ListNode insertNode(ListNode node, ListNode oldTop) {
        if (node.prev == null) {
            ListNode newTop = node.next;
            if (newTop != null) {
                newTop.prev = null;
            }
            return newTop;
        } else {
            node.prev.next = node.next;
            if (node.next != null) {
                node.next.prev = node.prev;
            }
            return oldTop;
        }
    }

    private static class ListNode {
        int value;
        ListNode prev;
        ListNode next;
        ListNode sibling;

        public ListNode(int value) {
            this.value = value;
        }

        @Override
        public String toString() {
            return "(" + value + ")";
        }
    }
}
