#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include "cell.h"


void Cell_print(Cell *cell) {
  //todo should return string not print it
  printf("DEBUG: Cell(%p) {\n", cell);
  printf("\trow: %d, column: %d\n", cell->row, cell->column);
  printf("\tlinks N: %p, E: %p, S: %p, W: %p\n", cell->links[N], cell->links[E], cell->links[S], cell->links[W]);
  printf("\tmerged N: %d, E: %d, S: %d, W: %d\n", cell->merged[N], cell->merged[E], cell->merged[S], cell->merged[W]);  
  printf("}\n");
}


Cell *Cell_north(Cell *cell) {
  return cell->links[N];
}

Cell *Cell_east(Cell *cell) {
  return cell->links[E];
}

Cell *Cell_south(Cell *cell) {
  return cell->links[S];
}

Cell *Cell_west(Cell *cell) {
  return cell->links[W];
}


void Cell_set_north(Cell *cell, Cell *other) {
  cell->links[N] = other;
}

void Cell_set_east(Cell *cell, Cell *other) {
  cell->links[E] = other;
}

void Cell_set_south(Cell *cell, Cell *other) {
  cell->links[S] = other;
}

void Cell_set_west(Cell *cell, Cell *other) {
  cell->links[W] = other;
}

bool Cell_merge(Cell *cell, Cell *other) {
  int NO = -1, ci = NO, oi = NO;
  
  for (int i = N; i <= W; i++) {
    if (cell->links[i] == other)
      ci = i;
    if (other->links[i] == cell) 
      oi = i;
  }
  
  if (ci > NO && oi > NO) {
    cell->merged[ci] = true;
    other->merged[oi] = true;
    return true;
  }
  //todo error: can't merge unrelated cells
  return false;
}

bool Cell_merged(Cell *cell, Cell *other) {
  for (size_t i = N; i <= W; i++) {
    if (cell->links[i] == other) {
      return cell->merged[i];
    }
  }
  return false;
}

//tests

void test_cell_links() {
  printf("> should return east cell\n");
  Cell cell = {.row = 5, .column = 5};
  Cell north = {.row = 4, .column = 5};
  Cell east = {.row = 5, .column = 6};
  Cell south = {.row = 6, .column = 5};
  Cell west = {.row = 5, .column = 4};

  Cell_set_north(&cell, &north);
  Cell_set_east(&cell, &east);
  Cell_set_south(&cell, &south);
  Cell_set_west(&cell, &west);

  if (Cell_north(&cell) != &north)
    printf("ERROR: wrong north vaule\n");
  if (Cell_east(&cell) != &east)
    printf("ERROR: wrong east vaule\n");
  if (Cell_south(&cell) != &south)
    printf("ERROR: wrong south vaule\n");
  if (Cell_north(&cell) != &north)
    printf("ERROR: wrong north vaule\n");
  
}

void test_create_cell() {
  printf("> should be able to create cell with specific row and column\n");
  Cell a = {.row = 1, .column = 2};
  if (a.row != 1)
    printf("ERROR: wrong row\n");
  if (a.column != 2)
    printf("ERROR: wrong column\n");
}


void test_link_two_cells() {
  printf("> should link two cells\n");
  Cell bt = {.row = 0, .column = 0}, *b = &bt; 
  Cell at = {.row = 0, .column = 1}, *a = &at;
  Cell ct = {.row = 0, .column = 2}, *c = &ct;

  Cell_set_east(b, a);
  Cell_set_west(a, b);
  
  Cell_set_east(a, c);
  Cell_set_west(c, a);

  
  if (!Cell_merge(a, b))
    printf("ERROR: unable to merge a to b\n");

  
  if (!Cell_merged(a, b))
    printf("ERROR: a not merged to b\n");
  if (!Cell_merged(b, a))
    printf("ERROR: b not merged to a\n");
  if (Cell_merged(a, c))
    printf("ERROR: a merged to c\n");
  if (Cell_merged(c, a))
    printf("ERROR: c merged to a\n");
}

void Cell_run_tests() {
  test_create_cell();
  test_cell_links();
  test_link_two_cells();
}
