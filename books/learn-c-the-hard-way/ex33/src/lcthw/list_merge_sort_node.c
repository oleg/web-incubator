#include <lcthw/list_merge_sort_node.h>
#include <lcthw/dbg.h>

inline void ListNode_swap(ListNode * a, ListNode * b)
{
  void *temp = a->value;
  a->value = b->value;
  b->value = temp;
}

ListNode* ListNode_insert_after(ListNode *a2, ListNode *b2) {
  printf("insert after:: %s -> %s \n", a2->value, b2->value);
  //  ListNode *a1 = a2->prev;
  ListNode *a3 = a2->next;
  
  //  ListNode *b1 = b2->prev;
  ListNode *b3 = b2->next;

  a2->next = b2;
  b2->prev = a2;
  
  b2->next = a3;
  if (a3 != NULL) {
    a3->prev = b2;
  }
  
  /* puts("exit insert after"); */
  return b3;
}

//todo:oleg do I need to copy this node pointer?
ListNode *ListNode_find_nth(ListNode *node, int n) {
  //  puts("find nth");
  while (n > 0) {
    node = node->next;
    n--;
  }
  return node;
}

ListNode *ListNode_merge(ListNode *head, int head_size,
                         ListNode *middle, int middle_size,
                         List_compare cmp)
{
  printf("merge:: head %s(%d), middle %s(%d) \n", head->value, head_size, middle->value, middle_size);


  ListNode *last;
  if (cmp(head->value, middle->value) <= 0) {
    last = head;
    head = head->next;
  } else {
    last = middle;
    middle = middle->next;
  }
  ListNode *first = last;

  //  while (h < head_size && m < middle_size) {  
  while (head != NULL && middle != NULL) {
    if (cmp(head->value, middle->value) <= 0) {
      last = head;
      head = head->next;
      /* h++; */
      
    } else {
      ListNode *temp = ListNode_insert_after(last, middle);
      last = middle;
      middle = temp;
      //middle = middle->next;
      /* m++; */
    }
  }
  
  /* while (head != NULL) {
   *   head = head->next;
   * } */

  while (head != NULL) {
    ListNode *tmp = ListNode_insert_after(last, head);
    last = head;
    head = tmp;
    //      last = middle;
    //m++;
  }

  while (middle != NULL) {
    ListNode *tmp = ListNode_insert_after(last, middle);
    last = middle;
    middle = tmp;
    //      last = middle;
    //m++;
  }

  puts("after merges:");
  ListNode *temp = first;
  for(int i = 0; i < 10 && temp != NULL; i++, temp = temp->next) {
    printf("       --> %s \n", temp->value);
  }
  
  //head->next = NULL;
  printf("exit merge: %s \n", first->value);
  first->prev = NULL;
  last->next = NULL;
  return first;
}

ListNode *ListNode_sort(ListNode *head, int size, List_compare cmp)
{
  if (head == NULL) {
    printf("sort head: NULL (%d))\n", size);        
  } else {
    printf("sort head: %s(%d)\n", head->value, size);        
  }
  if (size <= 1) {
    return head;
  }
  
  int head_size = size / 2;
  int middle_size = size - head_size;
  
  ListNode *middle = ListNode_find_nth(head, head_size);//todo rename to split
  middle->prev->next = NULL;
  middle->prev = NULL;

  head =ListNode_sort(head, head_size, cmp);
  middle = ListNode_sort(middle, middle_size, cmp);
  
  return ListNode_merge(head, head_size, middle, middle_size, cmp);
}

void List_merge_sort_node(List *list, List_compare cmp)
{
  puts("before list sort");
  ListNode* head = ListNode_sort(list->first, list->count, cmp);
  puts("after list sort");
  

  list->first = head;

  while (head != NULL) {
    /* printf("next value is %s \n", head->value); */
    head = head->next;
    //    puts("find end");
  }
  puts("just before end");
  list->last = head;
  puts("returning reult");

  LIST_FOREACH(list, first, next, cur) {
    printf("%s \n", cur->value);
  }
}
