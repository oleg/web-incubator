#include <stdio.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

int main(void)
{
    fprintf(stderr, "EACCES: %s\n", strerror(EACCES));

    errno = EACCES;
    perror("EACCES");

    exit(0);
}