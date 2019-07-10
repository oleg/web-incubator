#include <stdio.h>

int main()
{
  int c;
  
  while ((c = getchar()) != EOF) {
    if (c == ' ' || c == '\t')
      printf("\n");
    else
      putchar(c);
  }
}