#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <unistd.h>

#define BUFF_SIZE 4096


int main(int argc, char *argv[])
{
    int n;
    char buf[BUFF_SIZE];

    while((n = read(STDIN_FILENO, buf, BUFF_SIZE)) > 0) {
        if (write(STDOUT_FILENO, buf, n) != n) {
            printf("write error: %s\n", strerror(errno));
            exit(1);
        }
    }

    if (n < 0) {
        printf("read error: %s\n", strerror(errno));
        exit(1);
    }
    exit(0);
}