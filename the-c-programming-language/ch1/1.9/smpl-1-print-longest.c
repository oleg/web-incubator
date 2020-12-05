#include <stdio.h>

#define MAXLINE 1000
#define EOS '\0'

int get_line(char line[], int maxline);

void copy(char from[], char to[]);

int main()
{
  int len, max;
  char line[MAXLINE];
  char longest[MAXLINE];

  max = 0;
  while ((len = get_line(line, MAXLINE)) > 0)
    if (len > max) {
      max = len;
      copy(line, longest);
    }

  if (max > 0)
    printf("%s", longest);

  return 0;
}

int get_line(char line[], int limit)
{
  int c, i;

  for (i = 0; i < limit - 1 && (c = getchar()) != EOF && c != '\n'; ++i)
    line[i] = c;

  if (c == '\n') {
    line[i] = c;
    ++i;
  }
  line[i] = EOS;
  
  return i;
}

void copy(char from[], char to[])
{
  int i = 0;
  
  while ((to[i] = from[i]) != EOS)
    ++i;
}
