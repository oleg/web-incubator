#include <stdio.h>

#define MAXLINE 1000

int max;
char line[MAXLINE];
char longest[MAXLINE];


int getline(void);
void copy(void);


int main()
{
    extern int max;
    extern char longest[];
    
    int len;
    
    max = 0;
    while ((len = getline()) > 0)
        if (len > max) {
            max = len;
            copy();
        }

    if (max > 0)
        printf("%s", longest);
    
    return 0;
}


int getline(void)
{
    extern char line[];
    int c, i;

    for (i = 0; i < MAXLINE - 1  && (c = getchar()) != EOF && c != '\n'; ++i)
            line[i] = c;

    if (c == '\n') {
        line[i] = c;
        ++i;
    }

    line[i] = '\0';
    return i;

}

void copy(void)
{
  extern char longest[], line[];
  
  int i = 0;
  
  while ((longest[i] = line[i]) != '\0')
    ++i;
}
