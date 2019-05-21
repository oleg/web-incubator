#include <stdio.h>
#define MAX_LINE_SIZE 1000
#define MAX_OUT_LINE_SIZE 10
#define EOS '\0'

int get_line(char line[], int limit);
void reverse(char line[], char result[], int len);

int main()
{
    int len, max, i;
    char line[MAX_LINE_SIZE], reversed[MAX_LINE_SIZE];

    while ((len = get_line(line, MAX_LINE_SIZE)) > 0) {
        reverse(line, reversed, len);
        printf("%s", reversed);
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

void reverse(char line[], char result[], int len)
{
    int i, c, top;
    top = len - 2;

    for (i = 0; i <= top; i++) {
        result[i] = line[top - i];
    }
    result[++top] = '\n';
    result[++top] = EOS;
}