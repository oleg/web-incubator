#include <stdio.h>

void main()
{
  long nl = 0;
  long nt = 0;
  long nw = 0;
  int c;
  
  while((c = getchar()) != EOF) {
    if (c == '\n') nl++;
    if (c == '\t') nt++;
    if (c == ' ') nw++;
  }
  printf("%li\t%li\t%li\t\n", nl, nt, nw);
}