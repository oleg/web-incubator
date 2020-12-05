#include "minunit.h"
#include <lcthw/list.h>
#include <string.h>
#include <assert.h>


char *test_create()
{
  List *list = List_create();

  mu_assert(list != NULL, "Failed to create list");
  return NULL;
}

char *test_destroy()
{
  List *list = List_create();

  List_clear_destroy(list);

  return NULL;
}

char *test_push()
{
  char *test1 = "test1 data";
  char *test2 = "test2 data";
  char *test3 = "test3 data";

  List *list = List_create();

  List_push(list, test1);
  mu_assert(List_last(list) == test1, "Wrong last value on push 1");

  List_push(list, test2);
  mu_assert(List_last(list) == test2, "Wrong last value on push 2");

  List_push(list, test3);
  mu_assert(List_last(list) == test3, "Wrong last value on push 3");

  return NULL;
}

char *test_count_push()
{
  List *list = List_create();

  mu_assert(List_count(list) == 0, "Wrong count empty list");

  List_push(list, "1");
  mu_assert(List_count(list) == 1, "Wrong count 1 element");

  List_push(list, "22");
  mu_assert(List_count(list) == 2, "Wrong count 2 elements");

  List_push(list, "333");
  mu_assert(List_count(list) == 3, "Wrong count 3 elements");

  return NULL;
}

char *test_pop_empty()
{
  List *list = List_create();

  char *val = List_pop(list);

  mu_assert(val == NULL, "should return null");

  return NULL;
}

char *test_pop()
{
  char *test1 = "test1 data";
  char *test2 = "test2 data";
  char *test3 = "test3 data";

  List *list = List_create();

  List_push(list, test1);
  List_push(list, test2);
  List_push(list, test3);

  char *val;
  val = List_pop(list);
  mu_assert(val == test3, "Wrong value on first pop");

  val = List_pop(list);
  mu_assert(val == test2, "Wrong value on second pop");

  val = List_pop(list);
  mu_assert(val == test1, "Wrong value on third pop");

  return NULL;
}

char *test_count_pop()
{
  List *list = List_create();

  mu_assert(List_count(list) == 0, "Wrong count empty list");

  List_push(list, "1");
  List_push(list, "22");
  List_push(list, "333");

  char *val;
  val = List_pop(list);
  mu_assert(List_count(list) == 2, "Wrong count on first pop");

  val = List_pop(list);
  mu_assert(List_count(list) == 1, "Wrong count on second pop");

  val = List_pop(list);
  mu_assert(List_count(list) == 0, "Wrong count on third pop");

  return NULL;
}

char *test_unshift()
{
  char *test1 = "test 1 data";
  char *test2 = "test 2 data";
  char *test3 = "test 3 data";

  List *list = List_create();

  List_unshift(list, test1);
  mu_assert(List_first(list) == test1, "Wrong first value on unshift 1");

  List_unshift(list, test2);
  mu_assert(List_first(list) == test2, "Wrong first value on unshift 2");

  List_unshift(list, test3);
  mu_assert(List_first(list) == test3, "Wrong first value on unshift 3");

  return NULL;
}

char *test_count_unshift()
{
  List *list = List_create();

  mu_assert(List_count(list) == 0, "Wrong count empty list");

  List_unshift(list, "1");
  mu_assert(List_count(list) == 1, "Wrong count on unshift 1 element");

  List_unshift(list, "22");
  mu_assert(List_count(list) == 2, "Wrong count on unshift 2 elements");

  List_unshift(list, "333");
  mu_assert(List_count(list) == 3, "Wrong count on unshift 3 elements");

  return NULL;
}

char *test_shift_empty()
{
  List *list = List_create();

  char *val = List_shift(list);

  mu_assert(val == NULL, "should return null");

  return NULL;
}

char *test_count_shift()
{
  List *list = List_create();

  mu_assert(List_count(list) == 0, "Wrong count empty list");

  List_push(list, "1");
  List_push(list, "22");
  List_push(list, "333");


  List_shift(list);
  mu_assert(List_count(list) == 2, "Wrong count on first shift");

  List_shift(list);
  mu_assert(List_count(list) == 1, "Wrong count on second shift");

  List_shift(list);
  mu_assert(List_count(list) == 0, "Wrong count on third shift");

  return NULL;
}

char *test_shift()
{
  char *test1 = "test1 data";
  char *test2 = "test2 data";
  char *test3 = "test3 data";

  List *list = List_create();

  List_push(list, test1);
  List_push(list, test2);
  List_push(list, test3);

  char *val;
  val = List_shift(list);
  mu_assert(val == test1, "Wrong value on first shift");

  val = List_shift(list);
  mu_assert(val == test2, "Wrong value on second shift");

  val = List_shift(list);
  mu_assert(val == test3, "Wrong value on third shift");

  return NULL;
}

char *test_count_remove()
{
  List *list = List_create();

  mu_assert(List_count(list) == 0, "Wrong count empty list");

  List_push(list, "1");
  List_push(list, "22");
  List_push(list, "333");

  List_remove(list, list->first->next);
  mu_assert(List_count(list) == 2, "Wrong count on remove of the middle");

  List_remove(list, list->last);
  mu_assert(List_count(list) == 1, "Wrong count on remove of the last");

  List_remove(list, list->first);
  mu_assert(List_count(list) == 0, "Wrong count on remove of the first");

  return NULL;
}

char *test_remove()
{
  char *test1 = "test1 data";
  char *test2 = "test2 data";
  char *test3 = "test3 data";

  List *list = List_create();

  List_push(list, test1);
  List_push(list, test2);
  List_push(list, test3);

  char *val = List_remove(list, list->first->next);
  mu_assert(val == test2, "Wrong removed element.");
  mu_assert(List_count(list) == 2, "Wrong count after remove.");
  mu_assert(List_first(list) == test1, "Wrong first after remove.");
  mu_assert(List_last(list) == test3, "Wrong last after remove.");

  return NULL;
}

char *test_equal()
{
  List *a = List_create();
  List_push(a, "1");
  List_push(a, "3");
  List_push(a, "5");
  List_push(a, "4");
  List_push(a, "2");

  List *b = List_create();
  List_push(b, "1");
  List_push(b, "3");
  List_push(b, "5");
  List_push(b, "4");

  mu_assert(List_equal(a, b, (List_compare) strcmp) == 0, "equals when size is different");

  List_push(b, "2");

  mu_assert(List_equal(a, b, (List_compare) strcmp) == 1, "not equals when size and values are the same");

  List_pop(b);
  List_push(b, "xxx");
  mu_assert(List_equal(a, b, (List_compare) strcmp) == 0, "not equals when size is the same but values are different");
  return NULL;
}

char *all_tests()
{
  mu_suite_start();
  mu_run_test(test_create);
  mu_run_test(test_destroy);
  mu_run_test(test_push);
  mu_run_test(test_count_push);
  mu_run_test(test_pop_empty);
  mu_run_test(test_pop);
  mu_run_test(test_count_pop);
  mu_run_test(test_unshift);
  mu_run_test(test_count_unshift);
  mu_run_test(test_shift_empty);
  mu_run_test(test_count_shift);
  mu_run_test(test_shift);
  mu_run_test(test_count_remove);
  mu_run_test(test_remove);
  mu_run_test(test_equal);
  return NULL;
}

RUN_TESTS(all_tests);
