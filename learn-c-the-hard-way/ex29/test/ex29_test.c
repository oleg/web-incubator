#include <stdio.h>
#include <dlfcn.h>
#include "dbg.h"

typedef int (*lib_function)(const char *data);

int run_test(char *argv[]);

int main(int argc, char *argv[])
{
  int rc;
  
  char *args0[] = {"ex29", "build/libex29.so", "print_a_message", "hello there"};
  rc = run_test(args0);
  check(rc == 0, "Failed to print a message");

  char *args1[] = {"ex29", "build/libex29.so", "uppercase", "hello there"};
  rc = run_test(args1);
  check(rc == 0, "Failed to uppercase");

  char *args2[] = {"ex29", "build/libex29.so", "lowercase", "HELLO THERE"};
  rc = run_test(args2);
  check(rc == 0, "Failed to lowercase");

  char *args3[] = {"ex29", "build/libex29.so", "fail_on_purpose", "i fail"};
  rc = run_test(args3);
  check(rc != 0, "Failed to fail on purpose");

  char *args4[] = {"ex29", "build/libex29.so", "fail_on_purpose", NULL};
  rc = run_test(args4);
  check(rc != 0, "Failed to fail on purpose with null");

  char *args5[] = {"ex29", "build/libex29.so", "asdf", "fdsa"};
  rc = run_test(args5);
  check(rc != 0, "Failed to fail with wrong function");

  return 0;
  
 error:
  return 1;
}

int run_test(char *argv[])
{
  char *lib_file = argv[1];
  char *func_to_run = argv[2];
  char *data = argv[3];

  void *lib = dlopen(lib_file, RTLD_NOW);
  check(lib != NULL, "Failed to open the library %s: %s", lib_file, dlerror());

  lib_function func = dlsym(lib, func_to_run);
  check(func != NULL, "Did not find %s function in the library %s: %s", func_to_run, lib_file, dlerror());
  
  int rc;
  rc = func(data);
  check(rc == 0, "Function %s return %d for data: %s", func_to_run, rc, data);

  rc = dlclose(lib);
  check(rc == 0, "Failed to close %s", lib_file);

  return 0;
  
 error:
  return 1;
}
