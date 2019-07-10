#include <stdio.h>

int main()
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
  printf("%ld\t%ld\t%ld\t\n", nl, nt, nw);
}