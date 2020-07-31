#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

static void sig_int(int);

int main(void)
{
    int c;

    if (signal(SIGINT, sig_int) == SIG_ERR) {
        printf("signal error: %s\n", strerror(errno));
        exit(1);
    }

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

void sig_int(int signo)
{
    printf("interrupt");
}