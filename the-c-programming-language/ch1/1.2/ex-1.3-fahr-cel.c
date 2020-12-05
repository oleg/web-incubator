#include <stdio.h>

int main(void)
{
  float fahr, celsius;
  int step, lower, upper;
  
  lower = 0;
  upper = 300;
  step = 20;
  
  fahr = lower;
  
  printf("   F  \t   C\n");
  while (fahr <= upper) {
    celsius = 5.0 / 9.0 * (fahr - 32.0);
    printf("%6.0f\t%6.2f\n", fahr, celsius);
    fahr += step;
  }

  return 0;
}