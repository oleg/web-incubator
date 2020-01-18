package code;


public class ReverseList {

    public ListNode reverseList(ListNode head) {
        ListNode prev = null;
        ListNode curr = head;
        while (curr != null) {
            ListNode tmp = curr.next;
            curr.next = prev;
            prev = curr;
            curr = tmp;
        }
        return prev;
    }

    public ListNode reverseList_rec(ListNode head) {
        return reverseListRecursive(null, head);
    }

    public ListNode reverseListRecursive(ListNode first, ListNode second) {
        if (second == null) {
            return first;
        }
        ListNode secondNext = second.next;
        second.next = first;
        return reverseListRecursive(second, secondNext);
    }

}


