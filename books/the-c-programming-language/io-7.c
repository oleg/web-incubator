#include <stdio.h>

#define IN 1 /* внутри слова */
#define OUT 0 /* вне слова */

void main()
{
  int c, nl, nw, nc, state;
  
  state = OUT;
  nl = nw = nc = 0;
  
  while((c = getchar()) != EOF) {
    nc++;
    if (c == '\n') nl++;
    
    if (c == ' ' || c == '\t' || c == '\n') {
      if (state == IN) {
        nw++;
      }
      state = OUT;
    } else {
      state = IN;
    }
  }
  
  printf("nl: %i, nw: %i, nc: %i\n", nl, nw, nc);
}