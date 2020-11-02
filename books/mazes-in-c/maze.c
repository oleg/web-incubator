#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include "cell.h"

typedef struct Row Row;
struct Row {
  size_t n_cells;
  Cell* cells[];
};

typedef struct Grid Grid;
struct Grid {
  size_t n_rows;
  Row* rows[];
};

 /* st_employees* make_employees(unsigned n) { */
 /*    st_employees* s = malloc(sizeof(s_employees)+n*sizeof(char*)); */
 /*    if (!s) { perror("malloc make_employees"); exit(EXIT_FAILURE); }; */
 /*    s->No_of_Employees = n; */
 /*    for (unsigned i=0; i<n; i++) s->Employe_Names[i] = NULL; */
 /*    return s; */
 /* } */

Grid* Grid_create(size_t rows, size_t columns) {
  Grid* grid = malloc(sizeof(Grid)+rows*(sizeof(Row)));
  if (!grid) {
    //todo return or print an error message?
    return NULL;
  }
  grid->n_rows = rows;
  return grid;
}

void test_create_grid_non_null() {
  printf("> should create grid 5x5 with non-null rows and non-null cells\n");
  
  Grid* grid = Grid_create(5, 5);

  if (grid == NULL)
    printf("ERROR: failed to create grid\n");
  
  for (size_t i = 0; i < 5; i++)
    if (grid->rows[i] == NULL)
      printf("ERROR: row %zu is null\n", i);
}


void test_create_grid_correct_sizes() {
  printf("> should create grid 5x7 with correct rows and cells sizes\n");
  
  Grid* grid = Grid_create(5, 7);

  if (grid->n_rows != 5)
    printf("ERROR: wrong grid rows size %zu\n", grid->n_rows);
  
  for (size_t i = 0; i < grid->n_rows; i++)
    if (grid->rows[i]->n_cells != 7)
      printf("ERROR: row %zu has wrong cell size %zu\n", i, grid->rows[i]->n_cells);
}


void Grid_run_tests() {
  test_create_grid_non_null();
  test_create_grid_correct_sizes();
}

int run_tests() {
  printf("=> running tests\n");
  Cell_run_tests();
  Grid_run_tests();
  printf("=> tests finished\n");
  return EXIT_SUCCESS;
}


int main(int argc, char *argv[]) {
  if (argc > 1) {
    if (!strcmp(argv[1], "test")) {
      return run_tests();
    }
  }
  
  return EXIT_SUCCESS;
}
