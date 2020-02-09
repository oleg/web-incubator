#include <string.h>
#include <darray_algo.h>
#include <darray.h>
#include "unity.h"
#include "unity_fixture.h"

TEST_GROUP(DArrayAlgo);

int testcmp(char **a, char **b)
{
  return strcmp(*a, *b);
}

int is_sorted(DArray * array)
{
    int i = 0;
    for (i = 0; i < DArray_count(array) - 1; i++) {
        if (strcmp(DArray_get(array, i), DArray_get(array, i + 1)) > 0) {
            return 0;
        }
    }
    return 1;
}

static DArray *words = NULL;

TEST_SETUP(DArrayAlgo)
{
  words = DArray_create(0, 5);
  char *data[] = { "asdfasfd", "werwar", "13234", "asdfasfd", "oioj" };
  int i = 0;
  for (i = 0; i < 5; i++) {
      DArray_push(words, data[i]);
  }
}

TEST_TEAR_DOWN(DArrayAlgo)
{
    DArray_destroy(words);
}

//       DArray_heapsort
//
TEST(DArrayAlgo, quick_sort)
{
    int rc = DArray_qsort(words, (DArray_compare) testcmp);
    TEST_ASSERT_EQUAL_INT_MESSAGE(0, rc, "sort failed");
    TEST_ASSERT_EQUAL_INT_MESSAGE(1, is_sorted(words), "darray is not sorted");
}

TEST(DArrayAlgo, heap_sort)
{
    int rc = DArray_heapsort(words, (DArray_compare) testcmp);
    TEST_ASSERT_EQUAL_INT_MESSAGE(0, rc, "sort failed");
    TEST_ASSERT_EQUAL_INT_MESSAGE(1, is_sorted(words), "darray is not sorted");
}

TEST(DArrayAlgo, merge_sort)
{
    int rc = DArray_mergesort(words, (DArray_compare) testcmp);
    TEST_ASSERT_EQUAL_INT_MESSAGE(0, rc, "sort failed");
    TEST_ASSERT_EQUAL_INT_MESSAGE(1, is_sorted(words), "darray is not sorted");
}
