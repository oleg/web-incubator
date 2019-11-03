package algoclass.node;

class ListNode {
    private ListNode prev;
    private ListNode next;
    private int value;

    ListNode(int value, ListNode prev, ListNode next) {
        this.prev = prev;
        this.next = next;
        this.value = value;
    }

    int value() {
        return value;
    }

    ListNode next() {
        return next;
    }

    void next(ListNode next) {
        this.next = next;
    }

    ListNode prev() {
        return prev;
    }

    void prev(ListNode prev) {
        this.prev = prev;
    }

    /**
     * returns other's tail
     */
    ListNode appendOne(ListNode other) {
        if (other.prev() != null) {
            throw new IllegalArgumentException("Must not have prev");
        }
        ListNode oldNext = this.join(other);
        ListNode otherNext = other.next();
        other.join(oldNext);
        if (otherNext != null) {
            otherNext.prev(null);
        }
        return otherNext;
    }

    /**
     * returns this's tail
     */
    ListNode replaceTail(ListNode other) {
        if (other.prev() != null) {
            throw new IllegalArgumentException("Must not have prev");
        }
        ListNode oldNext = this.join(other);
        if (oldNext != null) {
            oldNext.prev(null);
        }
        return oldNext;
    }

    /**
     * returns other's tail
     */
    ListNode prependOne(ListNode other) {
        if (other.prev() != null) {
            throw new IllegalArgumentException("Must not have prev");
        }
        if (prev != null) {
            prev.join(other);
        }
        ListNode otherNext = other.join(this);
        if (otherNext != null) {
            otherNext.prev(null);
        }
        return otherNext;
    }

    private ListNode join(ListNode other) {
        ListNode oldNext = this.next;
        this.next(other);
        if (other != null) {
            other.prev(this);
        }
        return oldNext;
    }

    //todo test me
    ListNode findLast() {
        ListNode next = this;
        ListNode last = next;
        while (next != null) {
            last = next;
            next = next.next();
        }
        return last;
    }

    @Override
    public String toString() {
        return "ListNode(" + value +
                (prev != null ? " prev" : " null") +
                (next != null ? " next" : " null") +
                ')';
    }

    //
    void printDown() {
        ListNode tmp = this;
        while (tmp != null) {
            System.out.print(tmp.value());
            System.out.print("->");
            tmp = tmp.next();
        }
        System.out.println();
    }

}
