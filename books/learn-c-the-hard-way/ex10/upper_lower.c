#include <stdio.h>
#include <ctype.h>

int main(int argc, char* argv[]) {

	char a = 'a';
	printf("%c\n", a);
	printf("%c\n", a & 0xff);
	printf("%c\n", toupper(a));

	char b = 'B';
	printf("%c\n", b);
	printf("%c\n", b & 0xff);
	printf("%c\n", tolower(b));

	char c = '\0';
	printf("%c\n", c);
	printf("%c\n", tolower(c));
	printf("%c\n", tolower(c));

	return 0;
}
