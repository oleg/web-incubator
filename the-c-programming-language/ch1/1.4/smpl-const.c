#include <stdio.h>

#define LOWER 0
#define UPPER 300
#define STEP 20

int main()
{
  printf("   F  \t   C\n");
  for(int fahr = LOWER; fahr <= UPPER; fahr += STEP) {
    float celsius = 5.0 / 9.0 * (fahr - 32);
    printf("%3d\t%6.1f\n", fahr, celsius);
  }
}