package lcode;

class ListNode {
    int val;
    ListNode next;

    ListNode(int x) {
        val = x;
    }

    @Override
    public String toString() {
        return "N(" + val + ", " + next + ")";
    }
}
