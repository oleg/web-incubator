#include <stdio.h>

void main()
{
  int c;
  
  while((c = getchar()) != EOF) {
    if (c == '\t') {
      printf("\\t");
    } else if (c == '\b') { //чет эт не понятно зачем
      printf("\\b");
    } else if (c == '\\') { //the same
      printf("\\");
    } else {
      putchar(c);
    }
  }
}