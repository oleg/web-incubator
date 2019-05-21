#include <stdio.h>

main()
{
  float fahr, celsius;
  int step, lower, upper;
  
  lower = 0;
  upper = 300;
  step = 20;
  
  fahr = lower;
  
  while(fahr <= upper) {
    celsius = 5.0 / 9.0 * (fahr - 32.0);
    printf("%3.0f\t%.2f\n", fahr, celsius);
    fahr += step;
  }
}