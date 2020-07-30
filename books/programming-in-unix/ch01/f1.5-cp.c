#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

int main(int argc, char *argv[])
{
    int c;

    while((c = getc(stdin)) != EOF) {
        if (putc(c, stdout) == EOF) {
            printf("output error: %s\n", strerror(errno));
            exit(1);
        }
    }

    if (ferror(stdin)) {
        printf("input error: %s\n", strerror(errno));
        exit(1);
    }

    exit(0);
}