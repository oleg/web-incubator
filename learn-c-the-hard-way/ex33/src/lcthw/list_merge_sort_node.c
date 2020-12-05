#include <lcthw/list_merge_sort_node.h>
#include <lcthw/dbg.h>

ListNode* ListNode_insert_before(ListNode *a, ListNode *b)
{

  if (b->prev != NULL) {
    //todo fail
  }
  ListNode *b_next = b->next;
  ListNode *a_prev = a->prev;
  if (a_prev != NULL) {
    a_prev->next = b;
    b->prev = a_prev;
  }

  b->next = a;
  a->prev = b;

  if (b_next != NULL) {
    b_next->prev = NULL;
  }
  
  return b_next;
}

void ListNode_replace_tail(ListNode *last, ListNode *b)
{
  last->next = b;
  b->prev = last;
}

ListNode *ListNode_find_nth(ListNode *node, int n)
{
  while (n > 0) {
    node = node->next;
    n--;
  }
  return node;
}

ListNode *ListNode_find_last(ListNode *node)
{
  ListNode *last = NULL;
  while (node != NULL) {
    last = node;
    node = node->next;
  }
  return last;
}

ListNode *ListNode_merge(ListNode *a, 
                         ListNode *b,
                         List_compare cmp)
{
  ListNode *first = cmp(a->value, b->value) <= 0 ? a : b;
  ListNode *last = NULL;

  while (a != NULL && b != NULL) {
    last = a;
    if (cmp(a->value, b->value) <= 0) {
      a = a->next;
    } else {
      b = ListNode_insert_before(a, b);
    }
  }

  if (b != NULL) {
    ListNode_replace_tail(last, b);
  }

  return first;
}

ListNode *ListNode_sort(ListNode *head, int size, List_compare cmp)
{
  if (size <= 1) {
    return head;
  }
  
  int head_size = size / 2;
  int middle_size = size - head_size;
  
  ListNode *middle = ListNode_find_nth(head, head_size);//todo rename to split
  middle->prev->next = NULL;
  middle->prev = NULL;

  head = ListNode_sort(head, head_size, cmp);
  middle = ListNode_sort(middle, middle_size, cmp);
  
  return ListNode_merge(head, middle, cmp);
}

void List_merge_sort_node(List *list, List_compare cmp)
{
  ListNode* head = ListNode_sort(list->first, list->count, cmp);

  list->first = head;
  list->last = ListNode_find_last(head);
}
