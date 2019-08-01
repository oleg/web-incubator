#include <stdio.h>
#include <math.h>

#include "minunit.c"

int left(int i) {
  return i << 1; //*2  
}

int right(int i) {
  return (i << 1) + 1; //* 2 + 1  
}

int parent(int i) {
  return floor(i / 2.0);
}

void swap(int *array, int a, int b) {
  const temp = array[a];
  array[a] = array[b];
  array[b] = temp;
}

void print_as_line(int *array, int array_size) {
  int i;
  for(i = 0; i < array_size; i++) {
    printf("%2d ", array[i]);
  }
  //printf("\n");
}

void print_as_tree(char * buff, int *array, int array_size) {//TODO
  int i;
  for(i = 0; i < array_size; i++) {
    sprintf(buff, "%d ", array[i]);
  }
  sprintf(buff, "\n");  
}


static char * test_print_as_tree() {
  int data[6] = {1,1,1,2,2,2};
  char buff[6];
  
  print_as_tree(buff, data, 6);
  
  printf("==>> %s <<==", buff);
  mu_assert(22 == 22, "what?");
  return 0;
}

static char * all_tests() {
    mu_run_test(test_print_as_tree);
    return 0;
}

//test - app - runner
int main(int argc, const char* argv[]) {
  if (report_tests_result(all_tests()) != 0) {
    return 1;
  }
  return 0;
}