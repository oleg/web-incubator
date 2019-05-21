#include <stdio.h>

#define UPPER 300
#define LOWER 0
#define STEP 20

void main()
{
  printf("   F  \t   C\n");
  for(int fahr = UPPER; fahr >= LOWER; fahr -= STEP) {
    float celsius = 5.0 / 9.0 * (fahr - 32);
    printf("%3d\t%6.1f\n", fahr, celsius);
  }
}