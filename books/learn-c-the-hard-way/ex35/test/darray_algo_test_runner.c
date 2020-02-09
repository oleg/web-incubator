#include "unity.h"
#include "unity_fixture.h"

TEST_GROUP_RUNNER(DArrayAlgo)
{
  RUN_TEST_CASE(DArrayAlgo, quick_sort);
  RUN_TEST_CASE(DArrayAlgo, heap_sort);
  RUN_TEST_CASE(DArrayAlgo, merge_sort);
}
