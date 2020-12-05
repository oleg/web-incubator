package algoclass.node;

class List {
    private final ListCallback listCallback = new IntegrityChecker();
    private final ListNodeSort sorter = new ListNodeSort();

    private ListNode first;
    private ListNode last;

    void push(int value) {
        ListNode wasFirst = first;
        first = new ListNode(value, null, wasFirst);
        onFirstChange(wasFirst, first);

        listCallback.onChange(first, last);
    }

    void pushLast(int value) {
        ListNode wasLast = last;
        last = new ListNode(value, wasLast, null);
        onLastChange(wasLast, last);

        listCallback.onChange(first, last);
    }

    int pop() {
        checkNotEmpty();

        int value = first.value();
        first = first.next();
        onFirstChange(first, null);

        listCallback.onChange(first, last);
        return value;
    }

    int popLast() {
        checkNotEmpty();

        int value = last.value();
        last = last.prev();
        onLastChange(last, null);

        listCallback.onChange(first, last);
        return value;
    }

    private void checkNotEmpty() {
        if (first == null) {
            throw new IllegalStateException("no elements in the list");
        }
    }

    private void onFirstChange(ListNode node, ListNode prevNode) {
        if (node != null) {
            node.prev(prevNode);
        } else {
            last = first;
        }
    }

    private void onLastChange(ListNode node, ListNode nextNode) {
        if (node != null) {
            node.next(nextNode);
        } else {
            first = last;
        }
    }

    void sort() {
        first = sorter.sort(first);
        last = first != null ? first.findLast() : null;

        listCallback.onChange(first, last);
    }

}
