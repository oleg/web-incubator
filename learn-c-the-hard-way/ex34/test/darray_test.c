#include <lcthw/darray.h>
#include "unity.h"
#include "unity_fixture.h"

TEST_GROUP(DArrayCode);

static DArray *array = NULL;


TEST_SETUP(DArrayCode)
{
  array = DArray_create(sizeof(int), 100);
}

TEST_TEAR_DOWN(DArrayCode)
{
}


TEST(DArrayCode, create)
{
  TEST_ASSERT_NOT_NULL_MESSAGE(array, "DArray_create failed.");
  TEST_ASSERT_NOT_NULL_MESSAGE(array->contents, "contents are wrong in darray");
  TEST_ASSERT_EQUAL_INT_MESSAGE(0, array->end, "end isn't at the right spot");
  TEST_ASSERT_EQUAL_size_t_MESSAGE(sizeof(int), array->element_size, "element size is wrong.");
  TEST_ASSERT_EQUAL_INT_MESSAGE(100, array->max, "wrong max length on initial size");
}

TEST(DArrayCode, destroy)
{
  DArray_destroy(array);
}

TEST(DArrayCode, new)
{
  int *val1 = DArray_new(array);
  TEST_ASSERT_NOT_NULL_MESSAGE(val1, "failed to make a new element");

  int *val2 = DArray_new(array);
  TEST_ASSERT_NOT_NULL_MESSAGE(val2, "failed to make a new element");
}

TEST(DArrayCode, set)
{
  int *val1 = DArray_new(array);
  DArray_set(array, 0, val1);
  
  int *val2 = DArray_new(array);
  DArray_set(array, 1, val2);
}

TEST(DArrayCode, get)
{
  int *val1 = DArray_new(array);
  DArray_set(array, 0, val1);
  
  int *val2 = DArray_new(array);
  DArray_set(array, 1, val2);

  TEST_ASSERT_EQUAL_INT_MESSAGE(val1, DArray_get(array, 0), "Wrong first value.");
  TEST_ASSERT_EQUAL_INT_MESSAGE(val2, DArray_get(array, 1), "Wrong second value.");
}

TEST(DArrayCode, remove)
{
  int *val1 = DArray_new(array);
  DArray_set(array, 0, val1);
  int *val2 = DArray_new(array);
  DArray_set(array, 1, val2);
  
  int *val_check = DArray_remove(array, 0);
  TEST_ASSERT_NOT_NULL_MESSAGE(val_check, "Should not get NULL.");
  TEST_ASSERT_EQUAL_INT_MESSAGE(*val_check, *val1, "Should get the first value.");
  TEST_ASSERT_NULL_MESSAGE(DArray_get(array, 0), "Should be gone.");
  DArray_free(val_check);

  val_check = DArray_remove(array, 1);
  TEST_ASSERT_NOT_NULL_MESSAGE(val_check, "Should not get NULL.");
  TEST_ASSERT_EQUAL_INT_MESSAGE(*val_check, *val2, "Should get the first value.");
  TEST_ASSERT_NULL_MESSAGE(DArray_get(array, 1), "Should be gone.");
  DArray_free(val_check);
}

TEST(DArrayCode,  expand_contract)
{
  DArray_set(array, 0, DArray_new(array));
  DArray_set(array, 1, DArray_new(array));

  int old_max = array->max;
  DArray_expand(array);

  TEST_ASSERT_EQUAL_UINT_MESSAGE(array->max, old_max + array->expand_rate, "Wrong size after expand.");

  DArray_contract(array);
  TEST_ASSERT_EQUAL_UINT_MESSAGE(array->max, array->expand_rate + 1, "Should stay at the expand_rate at least.");

  DArray_contract(array);
  TEST_ASSERT_EQUAL_UINT_MESSAGE(array->max, array->expand_rate + 1, "Should stay at the expand_rate at least.");
}

TEST(DArrayCode, push_pop)
{
  DArray_expand(array);
  DArray_contract(array);
  
  for (int i = 0; i < 1000; i++) {
    int *val = DArray_new(array);
    *val = i * 333;
    DArray_push(array, val);
  }

  TEST_ASSERT_EQUAL_INT_MESSAGE(1201, array->max, "Wrong max size.");

  for (int i = 999; i >= 0; i--) {
    int *val = DArray_pop(array);
    TEST_ASSERT_NOT_NULL_MESSAGE(val, "Shouldn't get a NULL.");
    TEST_ASSERT_EQUAL_INT_MESSAGE(*val, i * 333, "Wrong value.");
    DArray_free(val);
  }
}
