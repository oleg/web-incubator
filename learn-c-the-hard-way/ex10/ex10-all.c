#include <stdio.h>

int main(int argc, char* argv[]) {

	if (argc < 2) {
		printf("ERROR: You need one argument.\n");
		return 1;
	}

	int i = 0,j = 0;
	for (i = 1; i < argc; i++) {
		for (j = 0; argv[i][j] != '\0'; j++) {
			char letter = argv[i][j];

			switch (letter) {
				case 'a': 
				case 'A':
				      printf("%d: 'A'\n", j);
				      break;
				case 'e': 
				case 'E':
				      printf("%d: 'E'\n", j);
				      break;
				case 'i': 
				case 'I':
				      printf("%d: 'I'\n", j);
				      break;
				case 'o': 
				case 'O':
				      printf("%d: 'O'\n", j);
				      break;
				case 'u': 
				case 'U':
				      printf("%d: 'U'\n", j);
				      break;
				case 'y': 
				case 'Y':
				      if (j > 2) {
					  printf("%d: 'Y'\n", j);
				      }
				      break;
				default:
				      printf("%d: %c is not a vowel\n", j, letter);

			}
		}
		printf("\n");
	}

	return 0;
}

