#include <lcthw/list_merge_sort.h>
#include <lcthw/dbg.h>

inline void ListNode_swap(ListNode * a, ListNode * b)
{
  void *temp = a->value;
  a->value = b->value;
  b->value = temp;
}

List *List_split(List *left) {
  
  int half = left->count / 2;
  ListNode *middle = left->first;
  for (int i = 0; i < half; i++) {
    middle = middle->next;
  }
  
  //what if 2 elements
  List *right = List_create();
  right->first = middle;
  right->last = left->last;
  right->count = left->count - half;
  
  left->last = middle->prev;
  left->last->next = NULL;
  middle->prev = NULL;
  left->count = half;
  
  return right;
}

List *List_merge(List *left, List *right, List_compare cmp) {
  List *result = List_create();
  
  ListNode *l_node = left->first;
  ListNode *r_node = right->first;

  while (l_node != NULL && r_node != NULL) {
    if (cmp(l_node->value, r_node->value) <= 0) {
      List_push(result, l_node->value);
      l_node = l_node->next;
    } else {
      List_push(result, r_node->value);
      r_node = r_node->next;
    }
  }

  while (l_node != NULL) {
    List_push(result, l_node->value);
    l_node = l_node->next;
  }
  
  while (r_node != NULL) {
    List_push(result, r_node->value);
    r_node = r_node->next;
  }
  
  return result;
}

List *List_merge_sort(List *left, List_compare cmp)
{
  if (left->count == 1) {
    return left;
  }
  List *right = List_split(left);//funny
  
  List *left_sorted = List_merge_sort(left, cmp);
  List *right_sorted = List_merge_sort(right, cmp);
  
  return List_merge(left_sorted, right_sorted, cmp);
}


