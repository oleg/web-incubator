#include <stdlib.h>
#include <stdio.h>
#include <dirent.h>
#include <errno.h>
#include <string.h>

int main(int argc, char *argv[])
{
    DIR           *dp;
    struct dirent *dirp;
    if (argc != 2) {
        printf("usage: s01-ls directory\n");
        exit(1);
    }

    if ((dp = opendir(argv[1])) == NULL) {
        printf("can't open %s: %s\n", argv[1], strerror(errno));
        exit(2);
    }

    while((dirp = readdir(dp)) != NULL) {
        printf("%s\n", dirp->d_name);
    }
    closedir(dp);

    exit(0);
}