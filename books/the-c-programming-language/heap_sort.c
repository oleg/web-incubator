#include <stdio.h>
#include <stdlib.h>

#include "minunit.c"
#include "rand.c"
#include "arrays.c"

void max_heapify(int curr, int *array, int heap_size) {
  int l = left(curr);
  int r = right(curr);
  int largest = curr;
  printf("curr: %d, l: %d, r: %d \n", curr, l, r);
  
  if (l <= heap_size && array[l-1] > array[curr-1]) {
    largest = l;
  } 
  if (r <= heap_size && array[r-1] > array[curr-1]) { 
    largest = r;
  }
  
  if (curr != largest) {
    swap(array, curr-1, largest-1);
    max_heapify(largest, array, heap_size);
  }
}

int app_main(int argc, const char* argv[]) {
  if (argc < 2) {
    printf("pass array size \n");
    return 1;
  }
  
  int array_size = atoi(argv[1]);
  if (array_size <= 0 ) {
    printf("got wrong heap size = %d from %s\n", array_size, argv[1]);
    return 1;
  }
  
  long in_memory_size = array_size * sizeof (int);//dunno,  may be not long ?
  int *heap = (int *) malloc(in_memory_size);
  if (heap == NULL) {
    printf("cant allocate memory for\n");
    return 1;
  }
  
  srand(time(NULL));
  int i;
  for(i = 0; i < array_size; i++) {
    heap[i] = random_num(array_size);
  }
  
  print_as_line(heap, array_size);
  
  free(heap);
  heap = NULL;
  
  return 0;
}

//TESTS
static char * test_left() {
  mu_assert(left(3) == 6, "left 3");
  mu_assert(left(2) == 4, "left 2");
  mu_assert(left(1) == 2, "left 1");
  mu_assert(left(100) == 200, "left 100");
  mu_assert(left(0) == 0, "left 0");
  return 0;
}

static char * test_right() {
  mu_assert(right(3) == 7, "right 3");
  mu_assert(right(2) == 5, "right 2");
  mu_assert(right(1) == 3, "right 1");
  mu_assert(right(100) == 201, "right 100");
  mu_assert(right(0) == 1, "right 0");
  return 0;
}

static char * test_parent() {
  mu_assert(parent(3) == 1, "parent 3");
  mu_assert(parent(2) == 1, "parent 2");
  mu_assert(parent(10) == 5, "parent 1");
  mu_assert(parent(11) == 5, "parent 100");
  mu_assert(parent(0) == 0, "parent 0");
  return 0;
}

static char * test_max_heapify() {
  int size = 5;
  int correct[] = {5,3,4,2,1};
  int i;
  
  int array[] = {5,3,4,2,1};
  max_heapify(1, array, size); 
  for (i = 0; i < size; i++) {
    mu_assert(array[i] == correct[i], "max_heapify nothing to do");
  }
  
  int array2[] = {5,1,4,2,3};
  max_heapify(2, array2, size);
  for (i = 0; i < size; i++) {
    mu_assert(array2[i] == correct[i], "max_heapify one element");
  }
  return 0;
}

static char * all_tests() {
    mu_run_test(test_left);
    mu_run_test(test_right);
    mu_run_test(test_parent);
    mu_run_test(test_max_heapify);
    return 0;
}

//test - app - runner
int main(int argc, const char* argv[]) {
  if (report_tests_result(all_tests()) != 0) {
    return 1;
  }
  return app_main(argc, argv);
}

