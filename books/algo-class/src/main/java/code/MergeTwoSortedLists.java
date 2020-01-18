package code;

import java.util.Iterator;

public class MergeTwoSortedLists {

    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {


        //x + y + x + y


        ListNode newList = l1;

        while (l1 != null && l2 != null) {
            newList.next = l2;

            l1 = l1.next;
            l2 = l2.next;
        }

        return null;
    }

}




