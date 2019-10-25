#include "minunit.h"
#include <lcthw/list_merge_sort_node.h>
//#include <lcthw/list.h>
#include <assert.h>
#include <string.h>

List_compare cmp = (List_compare) strcmp;

int is_sorted(List *words)
{
  LIST_FOREACH(words, first, next, cur) {
    if (cur->next && strcmp(cur->value, cur->next->value) > 0) {
      debug("%s %s", (char *) cur->value, (char *) cur->next->value);
      return 0;
    }
  }
  return 1;
}

char *test_sort_empty()
{
  List *list = List_create();
  List_merge_sort_node(list, cmp);

  List *expected = List_create();
  mu_assert(List_equal(list, expected, cmp) == 1, "Sorted empty list not equal to empty list.");
  return NULL;
}

char *test_sort_two_sorted()
{
  List *list = List_create();
  List_push(list, "a");
  List_push(list, "b");
  List_merge_sort_node(list, cmp);
  
  List *expected = List_create();
  List_push(expected, "a");
  List_push(expected, "b");
  mu_assert(List_equal(list, expected, cmp) == 1, "List is not sorted.");

  return NULL;
}

char *test_sort_three_sorted()
{
  List *list = List_create();
  List_push(list, "a");
  List_push(list, "b");
  List_push(list, "c");  
  List_merge_sort_node(list, cmp);
  
  List *expected = List_create();
  List_push(expected, "a");
  List_push(expected, "b");
  List_push(expected, "c");  
  mu_assert(List_equal(list, expected, cmp) == 1, "List is not sorted.");

  return NULL;
}
  
char *test_sort_five_sorted()
{
  List *list = List_create();
  List_push(list, "a");
  List_push(list, "b");
  List_push(list, "c");
  List_push(list, "d");
  List_push(list, "e");  
  List_merge_sort_node(list, cmp);
  
  List *expected = List_create();
  List_push(expected, "a");
  List_push(expected, "b");
  List_push(expected, "c");
  List_push(expected, "d");
  List_push(expected, "e");  
  mu_assert(List_equal(list, expected, cmp) == 1, "List is not sorted.");
  return NULL;
}
  
char *test_sort_two_reversed()
{
  List *list = List_create();
  List_push(list, "b");
  List_push(list, "a");
  List_merge_sort_node(list, cmp);
  
  List *expected = List_create();
  List_push(expected, "a");
  List_push(expected, "b");
  mu_assert(List_equal(list, expected, cmp) == 1, "List is not sorted.");

  return NULL;
}
  
char *test_sort_three_reversed()
{
  List *list = List_create();
  List_push(list, "c");
  List_push(list, "b");
  List_push(list, "a");
  List_merge_sort_node(list, cmp);
  
  List *expected = List_create();
  List_push(expected, "a");
  List_push(expected, "b");
  List_push(expected, "c");  
  mu_assert(List_equal(list, expected, cmp) == 1, "List is not sorted.");

  return NULL;
}
  
char *test_sort_five_reversed()
{
  List *list = List_create();
  List_push(list, "e");
  List_push(list, "d");  
  List_push(list, "c");
  List_push(list, "b");
  List_push(list, "a");
  List_merge_sort_node(list, cmp);
  
  List *expected = List_create();
  List_push(expected, "a");
  List_push(expected, "b");
  List_push(expected, "c");
  List_push(expected, "d");
  List_push(expected, "e");    
  mu_assert(List_equal(list, expected, cmp) == 1, "List is not sorted.");

  return NULL;
}
  
char *test_sort_unsorted()
{
  List *list = List_create();
  List_push(list, "e");
  List_push(list, "c");  
  List_push(list, "a");
  List_push(list, "b");
  List_push(list, "d");
  List_merge_sort_node(list, cmp);
  
  List *expected = List_create();
  List_push(expected, "a");
  List_push(expected, "b");
  List_push(expected, "c");
  List_push(expected, "d");
  List_push(expected, "e");    
  mu_assert(List_equal(list, expected, cmp) == 1, "List is not sorted.");

  return NULL;

  return NULL;
}

char *test_merge_sort_node()
{
  List *words = List_create();
  List_push(words, "XXXX1");
  List_push(words, "1234");
  List_push(words, "abcd");
  List_push(words, "xjvef");
  List_push(words, "NDSS");
  
  List_merge_sort_node(words, (List_compare) strcmp);
  mu_assert(is_sorted(words), "Words are not sorted after merge sort.");

  List_merge_sort_node(words, (List_compare) strcmp);
  mu_assert(is_sorted(words), "Should still be sorted after merge sort.");

  return NULL;
}

char *all_tests()
{
  mu_suite_start();

  mu_run_test(test_sort_empty);
  mu_run_test(test_sort_two_sorted);
  mu_run_test(test_sort_three_sorted);
  mu_run_test(test_sort_five_sorted);
  mu_run_test(test_sort_two_reversed);
  mu_run_test(test_sort_three_reversed);
  mu_run_test(test_sort_five_reversed);
  mu_run_test(test_sort_unsorted);  
  //mu_run_test(test_merge_sort_node);
  
  return NULL;
}


RUN_TESTS(all_tests);
