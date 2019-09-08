#include <stdio.h>
#define MAXLINE 1000
#define EOS '\0'

#define TAB_STOP_SIZE 4


int get_line(char line[], int maxline);
void copy(char from[], char to[]);


int main()
{
  char line[MAXLINE];

  while (get_line(line, MAXLINE) > 0)
    printf("%s", line);

  return 0;
}

int get_line(char line[], int limit)
{
  int c, i, j, shift;

  shift = 0;
  for (i = 0; i < limit - 1 - TAB_STOP_SIZE && (c = getchar()) != EOF && c != '\n'; ++i) {
    if (c != '\t') {
      line[i + shift] = c;
    } else {
      for (j = 0; j < TAB_STOP_SIZE; j++, shift++) {
        line[i + shift] = ' ';
      }
      shift--;
    }
  }
  
  if (c == '\n') {
    line[i + shift] = c;
    ++i;
  }
  line[i + shift] = EOS;
  
  return i;
}

void copy(char from[], char to[])
{
  int i = 0;
  
  while ((to[i] = from[i]) != EOS)
    ++i;
}
