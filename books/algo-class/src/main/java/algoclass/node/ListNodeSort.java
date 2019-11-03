package algoclass.node;

class ListNodeSort {

    private ListNodeMerger listNodeMerger = new ListNodeMerger();

    ListNode sort(ListNode a) {
        if (a == null || a.next() == null) {
            return a;
        }
        ListNode b = split(a);
        ListNode sortedA = sort(a);
        ListNode sortedB = sort(b);
        return listNodeMerger.merge(sortedA, sortedB);
    }

    private ListNode split(ListNode a) { //is it possible to effectively find middle?
        ListNode result = a;
        long count = 0;
        while (a != null) {
            a = a.next();
            count++;
        }
        long half = count / 2L;
        for (int i = 0; i < half; i++) {
            result = result.next();
        }
        result.prev().next(null);
        result.prev(null);
        return result;
    }

}
