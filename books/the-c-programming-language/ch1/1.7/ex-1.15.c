#include <stdio.h>

float to_celsius(float fahr) {
  return 5.0 / 9.0 * (fahr - 32.0);
}


int main() 
{
  float fahr;
  int step, lower, upper;
  
  lower = 0;
  upper = 300;
  step = 20;
  
  fahr = lower;
  
  while (fahr <= upper) {
    printf("%3.0f\t%.2f\n", fahr, to_celsius(fahr));
    fahr += step;
  }
  
  return 0;
}
