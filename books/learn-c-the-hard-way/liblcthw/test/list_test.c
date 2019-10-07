#include "minunit.h"
#include <lcthw/list.h>
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

char *test_count()
{
  char *test1 = "test1 data";
  char *test2 = "test2 data";
  char *test3 = "test3 data";

  List *list = List_create();

  mu_assert(List_count(list) == 0, "Wrong count empty list");
  
  List_push(list, test1);
  mu_assert(List_count(list) == 1, "Wrong count 1 element");

  List_push(list, test2);
  mu_assert(List_count(list) == 2, "Wrong count 2 elements");
  
  List_push(list, test3);
  mu_assert(List_count(list) == 3, "Wrong count 3 elements");
  
  return NULL;
}

char *test_pop_empty()
{

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

  char *val = List_pop(list);
  mu_assert(val == test3, "Wrong value on pop 1");
  
  return NULL;
}

char *all_tests()
{
  mu_suite_start();
  mu_run_test(test_create);
  mu_run_test(test_destroy);
  mu_run_test(test_push);
  mu_run_test(test_count);
  mu_run_test(test_push_pop);
  return NULL;
}

RUN_TESTS(all_tests);
