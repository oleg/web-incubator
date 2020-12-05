#include <stdio.h>

#define mu_assert(test, message) do { if (!(test)) return message; } while (0)
#define mu_run_test(test) do { char *message = test(); tests_run++; if (message) return message; } while (0)

int tests_run = 0;

int report_tests_result(char * result) {
    if (result != 0) {
        printf("%s\n", result);
	printf("ERROR IN TESTS (%d)\n\n", tests_run);
    }
    else {
        printf("ALL %d TESTS PASSED\n\n", tests_run);
    }
    return result != 0;  
}