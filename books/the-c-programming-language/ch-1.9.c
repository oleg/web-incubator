#include <stdio.h>
#define MAXLINE 1000
#define EOS '\0'

int getln(char line[], int lim);
void copy(char to[], char from[]);

int main()
{
  int len, max;
  char line[MAXLINE];
  char longest[MAXLINE];

  max = 0;
  while((len = getln(line, MAXLINE)) > 0) {
    if (len > max) {
      max = len;
      copy(longest, line);
    }
  }
  if (max > 0) {
    printf("%s", longest);
  }
  
  return 0;
}

int getln(char line[], int limit)
{
    int c, i;
    for(i = 0; i < limit - 1 && (c = getchar()) != EOF && c != '\n'; i++)
        line[i] = c;

    if (c == '\n') {
        line[i] = c;
        i++;
    }
    line[i] = EOS;
    return i;
}

void copy(char to[],  char from[]) {
    int i;
    i = 0;

    while((to[i] = from[i]) != EOS)
        i++;
}
