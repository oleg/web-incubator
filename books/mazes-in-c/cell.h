#ifndef CELL_H_INCLUDED
#define CELL_H_INCLUDED

#define Z 0
#define N 1
#define E 2
#define S 4
#define W 8


typedef struct Cell Cell;
struct Cell {
  int row;
  int column;
  int links;
};

bool Cell_link(Cell *cell, Cell *other);
bool Cell_linked(Cell *cell, Cell *other);
int Cell_relation(Cell *cell, Cell *other);

void Cell_print(Cell *cell);
void Cell_run_tests();

#endif
