#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include "cell.h"

typedef enum {
	      INVALID = 0,
	      N = 1,
	      E = 2,
	      S = 4,
	      W = 8
} Direction;

//typedef struct Grid Grid;
typedef struct {
  int rows;
  int columns;
  int size;
  int cells[];
} Grid;

bool Direction_valid(Direction dir) {
  switch (dir) {
  case N:
  case E:
  case S:
  case W:
    return true;
  default:
    return false;
  }
}

Direction Direction_opposite(Direction dir) {
  switch (dir) {
  case N: return S;
  case E: return W;
  case S: return N;
  case W: return E;
  default: return INVALID;
  }
}

void Grid_print(Grid *grid) {
  //return string?
  /* for (int i = 0; i < (rows*columns); i++) { */
  /*   printf("+++: (%d,%d)=%d\n", i/rows,i - (i/rows)*rows, g->cells[i]); */
  /* } */

  printf("Grid\nrows=%d, columns=%d, size=%d", grid->rows, grid->columns, grid->size);

  for (int c = 0; c < grid->size; c++) {
    if (c % grid->columns == 0)
      printf("\n");
    printf("%5d  ", grid->cells[c]);
  }
  
  printf("\n");
}

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

int* grid_cell(Grid *grid, int i, int j) {
  int c = i*grid->rows + j;
  if (c < 0 || c >= grid->size) {
    return NULL;
  }
  return &grid->cells[c];
}

int* grid_cell_neighbor(Grid *grid, int i, int j, Direction dir) {
  if (!Direction_valid(dir))
    return NULL;

  if (dir == N) i--;
  if (dir == E) j++;
  if (dir == S) i++;
  if (dir == W) j--;

  return grid_cell(grid, i, j);
}

//todo test
bool Grid_linked(Grid *grid, int i, int j, Direction dir) {
  if (!Direction_valid(dir))
    return false;
  int *c = grid_cell(grid, i, j);
  if (c == NULL)
    return false;

  return *c & dir;
}


bool Grid_link(Grid *grid, int i, int j, Direction dir) {
  //todo assert dir is correct otherwise return error
  int *c = grid_cell(grid, i, j);
  if (c == NULL)
    return false;

  int *o = grid_cell_neighbor(grid, i, j, dir);
  if (o == NULL)
    return false;

  int odir = Direction_opposite(dir);
  //todo assert odir is correct otherwise return an error

  *c |= dir;
  *o |= odir;

  return true;
}

void test_direction_valid() {
  printf("> should return true if direction valid\n");

  if (!Direction_valid(N))
    printf("ERROR: direction N must be valid\n");
  if (!Direction_valid(E))
    printf("ERROR: direction E must be valid\n");
  if (!Direction_valid(S))
    printf("ERROR: direction S must be valid\n");
  if (!Direction_valid(W))
    printf("ERROR: direction W must be valid\n");


  if (Direction_valid(INVALID))
    printf("ERROR: 0 is not valid direction\n");
  if (Direction_valid(3))
    printf("ERROR: 3 is not valid direction\n");
  if (Direction_valid(9))
    printf("ERROR: 9 is not valid direction\n");
  if (Direction_valid(100))
    printf("ERROR: 100 is not valid direction\n");
  
  printf("<\n");
}

void test_direction_opposite() {
  printf("> should return opposite direction\n");

  if (Direction_opposite(N) != S)
    printf("ERROR: opposite to N should be S\n");
  if (Direction_opposite(E) != W)
    printf("ERROR: opposite to E should be W\n");
  if (Direction_opposite(S) != N)
    printf("ERROR: opposite to S should be N\n");
  if (Direction_opposite(W) != E)
    printf("ERROR: opposite to W should be E\n");


  if (Direction_opposite(INVALID) != INVALID)
    printf("ERROR: opposite to INVALID should be INVALID\n");
  if (Direction_valid(3) != INVALID)
    printf("ERROR: opposite to 3 should be INVALID\n");
  if (Direction_valid(9) != INVALID)
    printf("ERROR: opposite to 9 should be INVALID\n");
  if (Direction_valid(100) != INVALID)
    printf("ERROR: opposite to 100 should be INVALID\n");
  
  printf("<\n");  
}

void test_create_grid() {
  printf("> should create grid 3x4 with non-null rows and non-null cells\n");

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

void test_address_cell() {
  printf("> should address cells by coordinates\n");
  Grid* g = Grid_create(3, 3);

  if (grid_cell(g, -1, -1) != NULL)
    printf("ERROR: non existing cell (-1,-1) is not null\n");
  
  if (grid_cell(g, 4, 4) != NULL)
    printf("ERROR: non existing cell (4, 4) is not null\n");
  
  
  int *c00 = grid_cell(g, 0, 0);
  if (c00 == NULL)
    printf("ERROR: (0,0) is null\n");
  if (*c00 != 0)
    printf("ERROR: wrong initial vaule of the cell (0,0)\n");

  *c00 = 7;

  int *c00x2 = grid_cell(g, 0, 0);
  if (*c00x2 != 7)
    printf("ERROR: wrong updated vaule of the cell (0,0)\n");

  
  int *c11 = grid_cell(g, 1, 1);
  if (c11 == NULL)
    printf("ERROR: (1, 1) is null\n");
  if (*c11 != 0)
    printf("ERROR: wrong initial vaule of the cell (1,1)\n");

  *c11 = 9;

  int *c11x2 = grid_cell(g, 1, 1);
  if (*c11x2 != 9)
    printf("ERROR: wrong updated vaule of the cell (1,1)\n");
  
  printf("<\n");
}

void test_non_neighbors_cells() {
  printf("> should ignore non neighbor links\n");
  Grid* grid = Grid_create(3, 3);

  if (Grid_link(grid, -10, -10, E))
    printf("ERROR: linked invalid cell\n");

  printf("<\n");
}

void test_link_neighbors_cells() {
  printf("> should be able to link two cells\n");
  
  Grid* grid = Grid_create(3, 3);

  if (Grid_linked(grid, 0, 0, E))
    printf("ERROR: (0,0) is linked to (0,1) before linking\n");

  if (!Grid_link(grid, 0, 0, E))
    printf("ERROR: failed to link (0,0) to (0,1)\n");

  if (!Grid_linked(grid, 0, 0, E))
    printf("ERROR: (0,0) is not linked to (0,1) after linking\n");

  if (!Grid_linked(grid, 0, 1, W))
    printf("ERROR: (0,1) is not linked to (0,0) after linking\n");


  if (Grid_linked(grid, 0, 0, N))
    printf("ERROR: (0,0) is linked to (-1,0)\n");

  if (Grid_linked(grid, 0, 0, S))
    printf("ERROR: (0,0) is linked to (1,0)\n");

  if (Grid_linked(grid, 0, 0, W))
    printf("ERROR: (0,0) is linked to (0,-1)\n");

  if (Grid_linked(grid, 0, 0, -100))
    printf("ERROR: invalid direction\n");
  
  printf("<\n");
}

void test_grid_cell_neighbor() {
  printf("> should return neighbor cells\n");
  
  Grid* g = Grid_create(3, 3);
  
  *grid_cell(g, 0, 0) = 0;
  *grid_cell(g, 0, 1) = 1;
  *grid_cell(g, 0, 2) = 2;
  
  *grid_cell(g, 1, 0) = 10;
  *grid_cell(g, 1, 1) = 11;
  *grid_cell(g, 1, 2) = 12;
  
  *grid_cell(g, 2, 0) = 20;
  *grid_cell(g, 2, 1) = 21;
  *grid_cell(g, 2, 2) = 22;  

  if (grid_cell_neighbor(g, 0, 0, N) != NULL)
    printf("ERROR: neighbor (0,0)N must be null\n");
  if (*grid_cell_neighbor(g, 0, 0, E) != 1)
    printf("ERROR: neighbor (0,0)E must be 1\n");
  if (*grid_cell_neighbor(g, 0, 0, S) != 10)
    printf("ERROR: neighbor (0,0)S must be 10\n");
  if (grid_cell_neighbor(g, 0, 0, W) != NULL)
    printf("ERROR: neighbor (0,0)W must be null\n");

  if (grid_cell_neighbor(g, 0, 0, 100) != NULL)
    printf("ERROR: neighbor for non existing direction must be null, %p, %d\n", grid_cell_neighbor(g, 0, 0, 100), *grid_cell_neighbor(g, 0, 0, 100));
  
  printf("<\n");
}

void Grid_run_tests() {
  printf(">> Grid_run_tests >>\n");
  test_direction_valid();
  test_direction_opposite();  
  test_create_grid();
  test_address_cell();
  test_grid_cell_neighbor();
  test_non_neighbors_cells();
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
