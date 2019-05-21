#include <stdio.h>
#include <stdlib.h>

#include "rand.c"
#include "arrays.c"

int main(int argc, const char* argv[]) {
  if (argc < 2) {
    printf("pass array size \n");
    return 1;
  }
  
  int array_size = atoi(argv[1]);
  if (array_size <= 0 ) {
    printf("got wrong array size = %d from %s\n", array_size, argv[1]);
    return 1;
  }
  
  long in_memory_size = array_size * sizeof (int);//dunno,  may be not long ?
  int *array = (int *) malloc(in_memory_size);
  if (array == NULL) {
    printf("cant allocate memory for\n");
    return 1;
  }
  
  srand(time(NULL));
  int i;
  for(i = 0; i < array_size; i++) {
    array[i] = random_num(array_size);
  }
  print_as_line(array, array_size);
  
  for(i = 1; i < array_size; i++) {
    int next_index = i;
    int next = array[next_index];
    
    int previous_index = next_index - 1;
    int previous = array[previous_index];
    
    while((previous_index >= 0) && (previous > next)) {
      
      swap(array, previous_index, next_index);
      next_index = previous_index;
      previous_index -= 1;
      
      next = array[next_index];
      previous = array[previous_index];
    }
  }
  print_as_line(array, array_size);
  
  free(array);
  array = NULL;
  
  return 0;
}

