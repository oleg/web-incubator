/*
 * remove:
 * - trailing blanks 
 * - tabs
 * - entirely blank lines
 */

#include <stdio.h>
#define MAX_LINE_SIZE 1000
#define MAX_OUT_LINE_SIZE 10
#define EOS '\0'
#define EOL '\n'

int get_line(char line[], int limit);
void trim_line_to(char line[], char to[]);

int main()
{
    int len, max, i;
    char line[MAX_LINE_SIZE], trimmed[MAX_LINE_SIZE];

    while ((len = get_line(line, MAX_LINE_SIZE)) > 0) {
        trim_line_to(line, trimmed);
        if (trimmed[0] != EOL) {
            printf("%s", trimmed);
        }
    }

    return 0;
}

int get_line(char line[], int limit)
{
    int c, i;
    for (i = 0; (i < limit - 1) && ((c = getchar()) != EOF) && c != EOL; i++) {
        line[i] = c;
    }

    if (c == EOL) {
        line[i] = c;
        i++;
    }
    line[i] = EOS;
    return i;
}

void trim_line_to(char line[], char to[])
{
    int i, j, c;
    int was_space;

    j = 0;
    was_space = 0;

    for (i = 0; (c = line[i]) != EOS; i++) {
        if (c == '\t') {
	    c = ' ';
        }
	
        if (c == ' ') {
            if (was_space != 1) {
                to[j++] = line[i];
            }
            was_space = 1;
        } else {
            to[j++] = line[i];
            was_space = 0;
        }
    }
    to[i] = EOS;
}
