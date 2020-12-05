#include <stdio.h>

int main(int argc, char *argv[])
{

  int ages[] = { 11, 22, 33 };
  printf("%d\n", sizeof(ages));
  printf("%d %d %d\n", ages + 0, *(ages + 0), ages[0]);
  printf("%d %d %d\n", ages + 1, *(ages + 1), ages[1]);
  printf("%d %d %d\n", ages + 2, *(ages + 2), ages[2]);
  
  printf("===\n");

  char name[] = { 'a', 'b', 'c' };
  printf("%d\n", sizeof(name));
  printf("%d %c %c\n", name + 0, *(name + 0), name[0]);
  printf("%d %c %c\n", name + 1, *(name + 1), name[1]);
  printf("%d %c %c\n", name + 2, *(name + 2), name[2]);
  
  return 0;
}
