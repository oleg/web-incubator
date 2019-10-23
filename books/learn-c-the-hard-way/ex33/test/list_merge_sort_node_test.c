#include "minunit.h"
#include <lcthw/list_merge_sort_node.h>
#include <assert.h>
#include <string.h>

char *values[] = { "XXXX", "1234", "abcd", "xjvef", "NDSS" };

#define NUM_VALUES 5

List *create_words()
{
  List *words = List_create();

  for(int i = 0; i < NUM_VALUES; i++) {
    List_push(words, values[i]);
  }

  return words;
}

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

char *test_merge_sort_node()
{
  List *words = create_words();
  
  List *res = List_merge_sort_node(words, (List_compare) strcmp);
  mu_assert(is_sorted(res), "Words are not sorted after merge sort.");

  List *res2 = List_merge_sort_node(res, (List_compare) strcmp);
  mu_assert(is_sorted(res), "Should still be sorted after merge sort.");

  List_destroy(words);
  List_destroy(res);
  List_destroy(res2);
  
  return NULL;
}

char *all_tests()
{
  mu_suite_start();

  mu_run_test(test_merge_sort_node);
  
  return NULL;
}


RUN_TESTS(all_tests);
