#include <stdio.h>

int main() 
{
  int c, nl = 0;
  int histo[30];
  
  for (int i = 0; i < 30; i++) 
    histo[i] = 0;
  
  while ((c = getchar()) != EOF) {
    if (c != ' ' && c != '\t' && c != '\n') {
      nl++;
    } else {
      if (nl < 29) {
	histo[nl]++;
      } else {
	histo[29]++;
      }
      nl = 0;
    }
  }
  
  for (int i = 0; i < 30; i++) {
    printf("%2d: ", i);
    for (int j = 0; j < histo[i]; j++) {
      printf("=");
    }
    printf("\n");
  }
}
