#include <stdio.h>
#define MAX_LINE_SIZE 1000
#define EOS '\0'

int get_line(char line[], int limit);
void copy_line(char to[], char from[]);

/**
 * print the length of arbitrarily long input lines, and as much as possible of the text.
**/
int main()
{
    int len, max, num, i;
    char line[MAX_LINE_SIZE];
    char longest[MAX_LINE_SIZE];

    num = max = 0;
    for (i = 0; (len = get_line(line, MAX_LINE_SIZE)) > 0; i++) {
      if (len > max) {//or >= ?
            max = len;
            num = i + 1;
            copy_line(line, longest);
        }
    }

    if (max > 0) {
        printf("the %d line is the longest. Here is the it's first %d symbols: \n%s", num, MAX_LINE_SIZE, longest);
    }

    return 0;
}

int get_line(char line[], int limit)
{
    int c, i, j;
    for(i = 0, j = 0; (c = getchar()) != EOF && c != '\n'; j++) {
        if (i < limit - 1) {
            line[i] = c;
            i++;
        }
        j++;
    }

    if (c == '\n') {
        line[i] = c;
        i++;
    }
    line[i] = EOS;
    return j;
}

void copy_line(char from[],  char to[]) {
    int i = 0;

    while((to[i] = from[i]) != EOS)
        i++;
}
