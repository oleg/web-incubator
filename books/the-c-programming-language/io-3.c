#include <stdio.h>

void main()
{
  long nl = 0;
  int c;
  
  while((c = getchar()) != EOF) {
    if (c == '\n') {
      nl++;
    }
  }
  printf("%li\n", nl);
}