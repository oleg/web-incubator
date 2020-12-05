#include <stdio.h>

int recur(int val)
{
  if (val == 10) {
    return 1;
  }
  return recur(val + 1) + val;
}


int main(int argc, char *argv[])
{

  printf("The result is: %d\n", recur(0));
  return 0;
}
