#include "unity.h"
#include "unity_fixture.h"

TEST_GROUP_RUNNER(DArrayCode)
{
  RUN_TEST_CASE(DArrayCode, create);
  RUN_TEST_CASE(DArrayCode, destroy);
  RUN_TEST_CASE(DArrayCode, new);
  RUN_TEST_CASE(DArrayCode, set);
  RUN_TEST_CASE(DArrayCode, get);
  RUN_TEST_CASE(DArrayCode, remove);
  RUN_TEST_CASE(DArrayCode, expand_contract);
  RUN_TEST_CASE(DArrayCode, push_pop);
}
