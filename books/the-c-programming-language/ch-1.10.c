#include <stdio.h>
#define MAXLINE 1000
#define EOS '\0'

char line[MAXLINE], longest[MAXLINE];

int getln();
void copy();

int main()
{
  int len, max;

  max = 0;
  while ((len = getln()) > 0) {
    if (len > max) {
      max = len;
      copy();
    }
  }
  if (max > 0) {
    printf("%s", longest);
  }
  
  return 0;
}

int getln()
{
    extern char line[MAXLINE];

    int c, i;

    for(i = 0; i < MAXLINE - 1 && (c = getchar()) != EOF && c != '\n'; i++)
        line[i] = c;

    if (c == '\n') {
        line[i] = c;
        i++;
    }
    line[i] = EOS;
    return i;
}

void copy() {
    extern char line[MAXLINE], longest[MAXLINE];

    int i;
    i = 0;

    while((longest[i] = line[i]) != EOS)
        i++;
}
