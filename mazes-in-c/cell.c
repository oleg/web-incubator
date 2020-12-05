#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include "cell.h"

void Cell_print(Cell *cell) {
  //todo should return string not print it
  printf("DEBUG: Cell(%p) {\n", cell);
  printf("\trow: %d, column: %d, links:%d\n", cell->row, cell->column, cell->links);
  printf("}\n");
}

int K[3][3] = {{Z, N, Z},
	       {W, Z, E},
	       {Z, S, Z},};

int Cell_relation(Cell *c, Cell *o) {
  int ir = 1 + c->row - o->row;
  int ic = 1 + c->column - o->column;
  if (ir < 0 || ir >= 3)
    return false;
  if (ic < 0 || ic >= 3)
    return false;

  return K[ir][ic];
}

bool Cell_link(Cell *c, Cell *o) {
  int k = Cell_relation(c, o);
  if (!k)
    return false;

  c->links |= k;
  return true;
}

bool Cell_linked(Cell *c, Cell *o) {
  int k = Cell_relation(c, o);
  if (!k)
    return false;

  return c->links & k;
}

//tests
void test_create_cell() {
  printf("> should be able to create cell with specific row, column, and links\n");
  Cell a = {.row = 1, .column = 2};
  if (a.row != 1)
    printf("ERROR: wrong row\n");
  if (a.column != 2)
    printf("ERROR: wrong column\n");
  if (a.links != Z)
    printf("ERROR: wrong links vaule\n");
  printf("<\n");  
}


void test_non_linked_by_default() {
  printf("> should see cells as unlinked by default\n");
  Cell a0t = {.row = 0, .column = 0}, *a0 = &a0t; 
  Cell a1t = {.row = 0, .column = 1}, *a1 = &a1t;
  Cell a2t = {.row = 0, .column = 2}, *a2 = &a2t;

  Cell b0t = {.row = 1, .column = 0}, *b0 = &b0t; 
  Cell b1t = {.row = 1, .column = 1}, *b1 = &b1t;
  Cell b2t = {.row = 1, .column = 2}, *b2 = &b2t;

  Cell c0t = {.row = 2, .column = 0}, *c0 = &c0t; 
  Cell c1t = {.row = 2, .column = 1}, *c1 = &c1t;
  Cell c2t = {.row = 2, .column = 2}, *c2 = &c2t;

  Cell* cells[] = {a0, a1, a2,
		   b0, b1, b2,
		   c0, c1, c2};
  
  for (int i = 0; i < 9; i++)
    for (int j = 0; j < 9; j++)
      if (Cell_linked(cells[i], cells[j]))
	printf("ERROR: cells %d and %d must not be linked\n", i, j);
  printf("<\n");
}

void test_link_neighbor_cells() {
  printf("> should link neighbor cells\n");
  Cell a1t = {.row = 0, .column = 1}, *a1 = &a1t;
  
  Cell b0t = {.row = 1, .column = 0}, *b0 = &b0t;
  Cell b1t = {.row = 1, .column = 1}, *b1 = &b1t;
  Cell b2t = {.row = 1, .column = 2}, *b2 = &b2t;
  
  Cell c1t = {.row = 2, .column = 1}, *c1 = &c1t;

  
  if (!Cell_link(b1, b2))
    printf("ERROR: unable to link b1 to b2\n");

  if (!Cell_linked(b1, b2))
    printf("ERROR: b1 is not linked to b2\n");

  if (Cell_linked(b1, b0))
    printf("ERROR: b1 is linked to b0\n");

  if (Cell_linked(b1, b1))
    printf("ERROR: b1 is linked to b1\n");
  
  if (Cell_linked(b1, a1))
    printf("ERROR: b1 is linked to a1\n");

  if (Cell_linked(b1, c1))
    printf("ERROR: b1 is linked to c1\n");


  printf("<\n");
}

void test_link_non_neighbor_cells() {
  printf("> should not link non-neighbor cells\n");
  Cell a0t = {.row = 0, .column = 0}, *a0 = &a0t; 
  Cell a1t = {.row = 0, .column = 1}, *a1 = &a1t;
  Cell a2t = {.row = 0, .column = 2}, *a2 = &a2t;

  Cell b0t = {.row = 1, .column = 0}, *b0 = &b0t; 
  Cell b1t = {.row = 1, .column = 1}, *b1 = &b1t;
  Cell b2t = {.row = 1, .column = 2}, *b2 = &b2t;

  Cell c0t = {.row = 2, .column = 0}, *c0 = &c0t; 
  Cell c1t = {.row = 2, .column = 1}, *c1 = &c1t;
  Cell c2t = {.row = 2, .column = 2}, *c2 = &c2t;

  Cell *linkable[] = {a1, b0, b2, c1};
  Cell *unlinkable[] = {b1, a0, a2, c0, c2};

  for (int i = 0; i < 4; i++)
    if (!Cell_link(b1, linkable[i]))
      printf("ERROR: must be able to link (1,1) and (%d,%d)\n",
	     linkable[i]->row, linkable[i]->column);
  
  for (int i = 0; i < 5; i++)
    if (Cell_link(b1, unlinkable[i]))
      printf("ERROR: must not be able to link (1,1) and (%d,%d)\n",
	     unlinkable[i]->row, unlinkable[i]->column);

  printf("<\n");
}

void Cell_run_tests() {
  printf(">> Cell_run_tests >>\n");
  test_create_cell();
  test_non_linked_by_default();
  test_link_non_neighbor_cells();
  test_link_neighbor_cells();
  printf("<< Cell_run_tests <<\n");
}
