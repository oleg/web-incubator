#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

int main(void)
{
    printf("process id: %ld\n", (long)getpid());
    exit(0);
}