#include <stdio.h>

short is_space(int check) {
  return check == ' ' || check == '\t' || check == '\n';
}

void main() 
{
  int i;
  int c, max_nw = 0, nl = 0;
  int histo[30];
  
  for(i = 0; i < 30; i++)
    histo[i] = 0;
  
  while((c = getchar()) != EOF) {
    if (!is_space(c)) {
      nl++;
    } else {
      if (nl > 29) {
	nl = 29;
      }
      int nw = ++histo[nl];
      if (nw > max_nw) {
	max_nw = nw;
      }
      nl = 0;
    }
  }
  
  printf("\nHistogram:\n");
  int level;
  for(level = max_nw; level > 0; level--) {

    for(i = 0; i < 30; i++){
      if (i == 0) {
	printf("%5d", level);
      } else if (histo[i] >= level ) {
	printf("  *");
      } else {
	printf("   ");
      }
    }
    printf("\n");
  }
  
  printf("     ");
  for(i = 1; i < 30; i++) {
    printf("%3d", i);
  }
  printf("\n");
}