#include <stdio.h>
#define MAX_LINE_SIZE 1000
#define MAX_OUT_LINE_SIZE 10
#define EOS '\0'

int get_line(char line[], int limit);
void trim_line_to(char line[], char to[]);

int main()
{
    int len, max, i;
    char line[MAX_LINE_SIZE], trimmed[MAX_LINE_SIZE];

    while ((len = get_line(line, MAX_LINE_SIZE)) > 0) {
        trim_line_to(line, trimmed);
        if (trimmed[0] != '\n') {
            printf("%s", trimmed);
        }
    }

    return 0;
}

int get_line(char line[], int limit)
{
    int c, i;
    for (i = 0; (i < limit - 1) && ((c = getchar()) != EOF) && c != '\n'; i++) {
        line[i] = c;
    }

    if (c == '\n') {
        line[i] = c;
        i++;
    }
    line[i] = EOS;
    return i;
}

void trim_line_to(char line[], char to[])
{
    int i, j, c;
    int wasSpace;
    wasSpace = 0;
    j = 0;

    for (i = 0; (c = line[i]) != EOS; i++) {
        if (c == '\t') {
            c = ' ';
        }
        if (c == ' ') {
            if (wasSpace != 1) {
                to[j++] = line[i];
            }
            wasSpace = 1;
        } else {
            to[j++] = line[i];
            wasSpace = 0;
        }
    }
    to[i] = EOS;
}