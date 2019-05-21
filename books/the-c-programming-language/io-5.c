#include <stdio.h>

void main()
{
  int skip_spaces = 0;
  int c;
  
  while((c = getchar()) != EOF) {
    
    if (skip_spaces == 0) {
      putchar(c);
    }
    
    if (skip_spaces == 1 && c != ' ') {
        putchar(c);
        skip_spaces = 0;
    }
    
    if (c == ' ') {
      skip_spaces = 1;
    }
  }
}