#ifndef CELL_H_INCLUDED
#define CELL_H_INCLUDED


#define N 0
#define E 1
#define S 2
#define W 3

typedef struct Cell Cell;
struct Cell {
  int row;
  int column;

  Cell *links[4];
  bool merged[4];
  
};

Cell *Cell_north(Cell *cell);
Cell *Cell_east(Cell *cell);
Cell *Cell_south(Cell *cell);
Cell *Cell_west(Cell *cell);

void Cell_set_north(Cell *cell, Cell *other);
void Cell_set_east(Cell *cell, Cell *other);
void Cell_set_south(Cell *cell, Cell *other);
void Cell_set_west(Cell *cell, Cell *other);

bool Cell_merge(Cell *cell, Cell *other);
bool Cell_merged(Cell *cell, Cell *other);

//util
void Cell_print(Cell *cell);
void Cell_run_tests();

#endif
