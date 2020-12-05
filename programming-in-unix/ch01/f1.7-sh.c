#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <errno.h>
#include <sys/wait.h>

#define MAX_LINE 10000

int main(void)
{
    char buf[MAX_LINE];
    pid_t pid;
    int status;

    printf("%% ");
    while((fgets(buf, MAX_LINE, stdin)) != NULL) {
        int last = strlen(buf) - 1;
        if (buf[last] == '\n') {
            buf[last] = 0;
        }

        if ((pid = fork()) < 0) {
            printf("fork error: %s\n", strerror(errno));
            exit(1);
        } else if (pid == 0) {
            execlp(buf, buf, (char *)0);
            printf("couldn′t execute: %s :%s\n", buf, strerror(errno));
            exit(127);
        }


        if ((pid = waitpid(pid, &status, 0)) < 0) {
            printf("waitpid error\n");
            exit(1);
        }
        printf("%% ");
    }

    exit(0);
}