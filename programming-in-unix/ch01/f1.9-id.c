#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(void)
{
    printf("user id: %d, group id: %d\n", getuid(), getgid());
    exit(0);
}