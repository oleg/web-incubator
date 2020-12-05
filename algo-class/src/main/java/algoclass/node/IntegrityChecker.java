package algoclass.node;

class IntegrityChecker implements ListCallback {

    //todo test me
    @Override
    public void onChange(ListNode first, ListNode last) {
//        checkSizeIsPositive(size);
//        checkEmptyListHasFirstNull(first, size);
//        checkNonEmptyListHasNonNullFirst(first, size);
//        checkNonEmptyListHasNonNullLast(last, size);
        checkNothingBeforeFirst(first);
        checkNothingAfterLast(last);
        checkLastIsNullWhenFirstIsNull(first, last);
        checkLastIsNotNullWhenFirstIsNotNull(first, last);
//        traverseFirst(first, last, size);
//        traverseLast(first, last, size);
    }

    private void checkNonEmptyListHasNonNullLast(ListNode last, long size) {
        if (size > 0 && last == null) {
            throw new IllegalStateException("last can't be null if size is positive");
        }
    }

    private void checkNonEmptyListHasNonNullFirst(ListNode first, long size) {
        if (size > 0 && first == null) {
            throw new IllegalStateException("first can't be null if size is positive");
        }
    }

    private void checkEmptyListHasFirstNull(ListNode first, long size) {
        if (size == 0 && first != null) {
            throw new IllegalStateException("node must be null if size is 0");
        }
    }

    private void checkSizeIsPositive(long size) {
        if (size < 0) {
            throw new IllegalStateException("size can't be negative");
        }
    }

    private void checkNothingBeforeFirst(ListNode first) {
        if (first != null && first.prev() != null) {
            throw new IllegalStateException("first prev is not null");
        }
    }

    private void checkNothingAfterLast(ListNode last) {
        if (last != null && last.next() != null) {
            throw new IllegalStateException("last next is not null");
        }
    }

    private void checkLastIsNullWhenFirstIsNull(ListNode first, ListNode last) {
        if (first == null && last != null) {
            throw new IllegalStateException("first is null but last is not null");
        }
    }

    private void checkLastIsNotNullWhenFirstIsNotNull(ListNode first, ListNode last) {
        if (first != null && last == null) {
            throw new IllegalStateException("last is null but first is not null");
        }
    }

    private void traverseFirst(ListNode first, ListNode last, long size) {
        ListNode tmp = last;
        ListNode found = tmp;
        long count = 0;
        while (tmp != null) {
            found = tmp;
            tmp = tmp.prev();
            count++;
        }
        if (count != size) {
            throw new IllegalStateException("reached end and count is equal to list size");
        }
        if (found != first) {
            throw new IllegalStateException(String.format("first that is found: %s is not equal to expected first: %s", found, first));
        }
    }

    private void traverseLast(ListNode first, ListNode last, long size) {
        ListNode tmp = first;
        ListNode found = tmp;
        long count = 0L;
        while (tmp != null) {
            found = tmp;
            tmp = tmp.next();
            count++;
        }
        if (count != size) {
            throw new IllegalStateException(String.format("reached end and count: %d is not equal to list size: %d", count, size));
        }
        if (found != last) {
            throw new IllegalStateException("reached end is not equal to last");
        }
    }
}