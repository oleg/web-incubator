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

void List_unshift(List *list, void *value)
{
  ListNode *node = calloc(1, sizeof(ListNode));
  check_mem(node);

  node->value = value;

  if (list->first == NULL) {
    list->first = node;
    list->last = node;
  } else {
    list->first->prev = node;
    node->next = list->first;
    list->first = node;
  }

  list->count++;

 error:
  return;
}

void *List_shift(List *list)
{
  int size = List_count(list);
  check(size > 0, "Can't shift empty list");

  void *result = list->first->value;

  if (size == 1) {
    list->first = NULL;
    list->last = NULL;
  } else {
    ListNode *new_first = list->first->next;
    new_first->prev = NULL;
    list->first = new_first;
  }
  
  list->count--;
  
  return result;
 error:
  return NULL;
}

void *List_remove(List *list, ListNode *node)
{
  
  return NULL;
 error:
  return NULL;
}

