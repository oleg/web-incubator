#include <lcthw/darray.h>
#include <stdio.h>
#include <string.h>

DArray *DArray_create(size_t element_size, size_t initial_max)
{
  DArray *array = malloc(sizeof(DArray));
  if (!array) {
    fprintf(stderr, "[ERROR] %s\n", "Out of memory");
    return NULL;
  }

  array->max = initial_max;
  if (array->max <= 0) {
    fprintf(stderr, "[ERROR] %s\n", "You must set an initial_max > 0.");
    return NULL;
  }

  array->contents = calloc(initial_max, sizeof(void *));
  if (!array->contents) {
    fprintf(stderr, "[ERROR] %s\n", "Out of memory");
  }

  array->end = 0;
  array->element_size = element_size;
  array->expand_rate = DEFAULT_EXPAND_RATE;
  return array;
}

void DArray_clear(DArray * array)
{
  int i = 0;
  if (array->element_size > 0) {
    for (i = 0; i < array->max; i++) {
      if (array->contents[i] != NULL) {
        free(array->contents[i]);
      }
    }
  }
}

static inline int DArray_resize(DArray * array, size_t newsize)
{
  array->max = newsize;
  if (array->max <= 0) {
    fprintf(stderr, "[ERROR] %s\n", "The newsize must be > 0.");
    return -1;
  }

  void *contents = realloc(array->contents, array->max * sizeof(void *));
  if (!array) {
    fprintf(stderr, "[ERROR] %s\n", "Out of memory");
    return -1;
  }

  array->contents = contents;
  return 0;
}

int DArray_expand(DArray * array)
{
  size_t old_max = array->max;
  int res = DArray_resize(array, array->max + array->expand_rate);
  if (res != 0) {
    fprintf(stderr, "[ERROR] %s\n", "Failed to expand array to new size");
    return -1;
  }

  memset(array->contents + old_max, 0, array->expand_rate + 1);
  return 0;
}

int DArray_contract(DArray * array)
{
  int new_size = array->end < (int)array->expand_rate ?
    (int)array->expand_rate : array->end;
  return DArray_resize(array, new_size + 1);
}

void DArray_destroy(DArray * array)
{
  if (array) {
    if (array->contents)
      free(array->contents);
    free(array);
  }
}

void DArray_clear_destroy(DArray * array)
{
  DArray_clear(array);
  DArray_destroy(array);
}

int DArray_push(DArray * array, void *el)
{
  array->contents[array->end] = el;
  array->end++;

  if (DArray_end(array) >= DArray_max(array)) {
    return DArray_expand(array);
  } else {
    return 0;
  }
}

void *DArray_pop(DArray * array)
{
  if (array->end - 1 < 0) {
    fprintf(stderr, "[ERROR] %s\n", "Attempt to pop from empty array");
    return NULL;
  }
  void *el = DArray_remove(array, array->end - 1);
  array->end--;

  if (DArray_end(array) > (int)array->expand_rate
      && DArray_end(array) % array->expand_rate) {
    DArray_contract(array);
  }
  return el;
}
