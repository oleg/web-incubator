#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include "cell.h"

typedef struct {
  int i;
  int j;
} C;
  
//typedef struct Grid Grid;
typedef struct {
  int rows;
  int columns;
  int size;
  int cells[];
} Grid;

Grid* Grid_create(int rows, int columns) {
  Grid* grid = malloc(sizeof(Grid) + sizeof(int)*rows*columns);
  if (!grid) {
    printf("DEBUG: Grid_create malloc failed\n");
    return NULL;
  }
  grid->rows = rows;
  grid->columns = columns;
  grid->size = rows*columns;
  for (int i = 0; i < grid->size; i++) {
    grid->cells[i] = 0;
  }
  return grid;
}

int Grid_get_cell(Grid *grid, int i, int j) {
  return 0;
}

void Grid_set_cell(Grid *grid, int i, int j, int cell) {
  
}


bool Grid_link(Grid *grid, C c, C o) {
  int a = Grid_get_cell(grid, c.i, c.j);
  
  return false;
}

void test_create_grid() {
  printf("> should create grid 3x4 with non-null rows and non-null cells\n");
  /* for (int i = 0; i < (rows*columns); i++) { */
  /*   printf("+++: (%d,%d)=%d\n", i/rows,i - (i/rows)*rows, g->cells[i]); */
  /* } */

  Grid* grid = Grid_create(3, 4);

  if (grid == NULL)
    printf("ERROR: failed to create grid\n");

  if (grid->rows != 3)
    printf("ERROR: wrong rows size\n");

  if (grid->columns != 4)
    printf("ERROR: wrong columns size\n");
  
  if (grid->size != 12)
    printf("ERROR: wrong size\n");
  

  for (int i = 0; i < grid->size; i++)
    if (grid->cells[i] != 0)
      printf("ERROR: unexpected cell vaule\n");

  printf("<\n");
}

void test_link_neighbors_cells() {
  printf("> should \n");
  
  Grid* grid = Grid_create(3, 3);
  if (!Grid_link(grid, (C){0,0}, (C){0,1}))
    printf("ERROR: failed to link (0,0) to (0,1)\n");
  
}

void Grid_run_tests() {
  printf(">> Grid_run_tests >>\n");
  test_create_grid();
  test_link_neighbors_cells();
  printf("<< Grid_run_tests <<\n");  
}

int run_tests() {
  printf(">>> TESTS >>>\n");
  Cell_run_tests();
  Grid_run_tests();
  printf("<<< TESTS <<<\n");
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
