#include <stdio.h>

#define FS ' '
#define LS '~'
#define SIZE (LS - FS) + 1

void main() 
{
  int c, i, j;
  int histo[SIZE];
  
  for(i = 0; i < SIZE; i++)
    histo[i] = 0;
  
  while((c = getchar()) != EOF) {
    if (c >= FS && c <= LS) {
      histo[c - FS]++;
    }
  }
  
  
  
  for(i = 0; i < SIZE; i++) {
    printf("  %c  ", i + FS);
    for(j = 0; j < histo[i]; j++)
      printf("=");
    printf("\n");
  }
}