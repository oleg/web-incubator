package algoclass.node;

class ListNodeMerger {

    ListNode merge(ListNode a, ListNode b) {
        checkNotNull(a, b);
        checkNoPrev(a, b);

        ListNode head = a.value() <= b.value() ? a : b;
        ListNode tail = null;

        while (a != null && b != null) {
            tail = a;
            if (a.value() <= b.value()) {
                a = a.next();
            } else {
                b = a.prependOne(b);
            }
        }
        if (b != null) {
            tail.replaceTail(b);
        }
        return head;
    }

    private void checkNotNull(ListNode a, ListNode b) {
        if (a == null) {
            throw new IllegalArgumentException("first argument must not be null");
        }
        if (b == null) {
            throw new IllegalArgumentException("second argument must not be null");
        }
    }

    //todo write tests for this
    private void checkNoPrev(ListNode a, ListNode b) {
        if (a.prev() != null) {
            throw new IllegalArgumentException("first argument must not have prev nodes");
        }
        if (b.prev() != null) {
            throw new IllegalArgumentException("second argument must not have prev nodes");
        }
    }

}
