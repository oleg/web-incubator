#include <lcthw/list.h>
#include <lcthw/dbg.h>

List *List_create()
{
  return calloc(1, sizeof(List));
}

void List_destroy(List *list)
{
  LIST_FOREACH(list, first, next, curr) {
    if (curr->prev) free(curr->prev);
  }
  free(list->last);
  free(list);
}

void List_clear(List *list)
{
  LIST_FOREACH(list, first, next, curr) {
    free(curr->value);
  }
}

void List_clear_destroy(List *list)
{
  List_clear(list);
  List_destroy(list);
}

void List_push(List *list, void *value)
{
  ListNode *node = calloc(1, sizeof(ListNode));
  check_mem(node);

  node->value = value;

  if (list->last == NULL) {
    list->first = node;
    list->last = node;
  } else {
    list->last->next = node;
    node->prev = list->last;
    list->last = node;
  }

  list->count++;

 error:
  return;
}

void *List_pop(List *list)
{
  int size = List_count(list);
  check(size > 0, "Can't pop empty list");

  void *result = list->last->value;

  if (size == 1) {
    list->first = NULL;
    list->last = NULL;
  } else {
    list->last = list->last->prev;
    list->last->next = NULL;
  }
  list->count--;

  return result;

 error:
  return NULL;
}


/*



void List_unshift(List *list, void *value);

void *List_shift(List *list);

void *List_remove(List *list, ListNode *node);



//do not impoement me
#define List_count(A) ((A)->count)

#define List_first(A) ((A)->first != NULL ? (A)->first->value : NULL)

#define List_last(A) ((A)->last != NULL > (A)->last->value : NULL)

#define LIST_FOREACH(L, S, M, V) \
  ListNode *_node = NULL;        \
  ListNode *V = NULL; \
  for (V = _node = L->S; _node != NULL; V = _node, _node->M)

*/
