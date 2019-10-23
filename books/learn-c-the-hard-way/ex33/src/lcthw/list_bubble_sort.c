#include <lcthw/list_bubble_sort.h>
#include <lcthw/dbg.h>

inline void ListNode_swap(ListNode * a, ListNode * b)
{
  void *temp = a->value;
  a->value = b->value;
  b->value = temp;
}

int List_bubble_sort(List *list, List_compare cmp)
{
  LIST_FOREACH(list, first, next, a) {
    LIST_FOREACH(list, first, next, b) {
      if (cmp(a->value, b->value) < 0) {
        ListNode_swap(a, b);
      }
    }
  }
  return 0;
}

