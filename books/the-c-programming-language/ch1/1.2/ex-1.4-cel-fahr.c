#include <stdio.h>

main()
{
  float fahr, celsius;
  int step, lower, upper;
  
  lower = -50;
  upper = 50;
  step = 5;
  
  celsius = lower;
  
  printf("   C  \t   F\n");
  while (celsius <= upper) {
    fahr = 9.0 / 5.0 * celsius + 32.0;
    printf("%6.0f\t%6.2f\n", celsius, fahr);
    celsius += step;
  }
}