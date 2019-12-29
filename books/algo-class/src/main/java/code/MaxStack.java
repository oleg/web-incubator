package code;

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
            ListNode tmp = max;
            ListNode prev = null;
            while (tmp != null && tmp.value > node.value) {
                prev = tmp;
                tmp = tmp.next;
            }
            if (prev == null) {
                ListNode ttt = max;
                max = node;
                node.prev = null;
                node.next = ttt;
                ttt.prev = node;
            } else {
                ListNode ttt = prev.next;
                prev.next = node;
                node.prev = prev;
                node.next = ttt;
                if (ttt != null) {
                    ttt.prev = node;
                }
            }
        }
    }


    public int pop() {
        int result = head.value;
        ListNode sibling = head.sibling;
        head = head.next;
        if (head != null) {
            head.prev = null;
        }
        //remove sibling
        if (sibling.prev == null) {
            //sibling is a max
            max = sibling.next;
            if (sibling.next != null) {
                sibling.next.prev = null;
            }
        } else {
            sibling.prev.next = sibling.next;
            if (sibling.next != null) {
                sibling.next.prev = sibling.prev;
            }
        }
        return result;
    }

    public int popMax() {
        int result = max.value;
        ListNode sibling = max.sibling;
        max = max.next;
        if (max != null) {
            max.prev = null;
        }
        //remove sibling
        if (sibling.prev == null) {
            //sibling is a max
            head = sibling.next;
            if (sibling.next != null) {
                sibling.next.prev = null;
            }
        } else {
            sibling.prev.next = sibling.next;
            if (sibling.next != null) {
                sibling.next.prev = sibling.prev;
            }
        }

        return result;
    }

    public int top() {
        return head.value;
    }

    public int peekMax() {
        return max.value;
    }

    @Override
    public String toString() {
        return followNext("max", max) + "|" + followNext("head", head);
    }

    private StringBuilder followNext(String label, ListNode start) {
        StringBuilder str = new StringBuilder();
        str.append("(").append(label).append(": start->");
        ListNode tmp = start;
        while (tmp != null) {
            str.append(tmp.value).append("->");
            tmp = tmp.next;
        }
        str.append("end)");
        return str;
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
