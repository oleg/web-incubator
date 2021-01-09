package lcode;

public class MergeTwoSortedLists {
    //todo also solve recursively
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode curr = new ListNode(666), first = curr;
        while (l1 != null && l2 != null) {
            if (l1.val <= l2.val) {
                curr = curr.next = l1;
                l1 = l1.next;
            } else {
                curr = curr.next = l2;
                l2 = l2.next;
            }
        }
        if (l1 != null) {
            curr.next = l1;
        }
        if (l2 != null) {
            curr.next = l2;
        }
        return first.next;
    }

}




