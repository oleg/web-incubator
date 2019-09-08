#include <stdio.h>

int main(int argc, char *argv[])
{
  char *another = "Zed";

  int *i = another;
  
  printf("another each: %c %c %c %c\n", another[0], another[1], another[2], another[3]);  
  printf("another: %s\n", another);

  printf("i int: %d\n", i);
  printf("i str: %s\n", i);
  
  return 0;
}

  
  
