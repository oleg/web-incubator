from dataclasses import dataclass, field
from typing import List
from cell import Cell

@dataclass(init = False)
class Grid:
    rows: int
    columns: int
    grid: List[List[Cell]] = field(default_factory = None, compare = False)


    def __init__(self, rows, columns):
        self.rows = rows
        self.columns = columns
        self.grid = self.prepare_grid()
        self.configure_cells()
    
    def prepare_grid(self):
        return [[Cell(r, c) for c in range(0, self.columns)] for r in range(0, self.rows)]

    def configure_cells(self):
        for c in self:
            row, col = c.row, c.column
            c.north = self[row - 1, col]
            c.east = self[row, col + 1]
            c.south = self[row + 1, col]
            c.west = self[row, col - 1]
            
    def __getitem__(self, index):
        r, c = index
        if r not in range(0, self.rows) or c not in range(0, self.columns):
            return None
        return self.grid[r][c]

    def __iter__(self):
        return (self.grid[r][c] for r in range(0, self.rows) for c in range(0, self.columns))


def test_can_create_grid():
    g = Grid(3, 3)
    assert g.rows == 3
    assert g.columns == 3

def test_getitem_check_boundaries():
    g = Grid(3, 3)
    assert g[-1, -1] is None
    assert g[0, -1] is None
    assert g[-100, -100] is None
    assert g[3, 3] is None
    assert g[13, 14] is None

def test_creates_cells():
    g = Grid(2, 2)
    assert g[0, 0] == Cell(0, 0)
    assert g[0, 1] == Cell(0, 1)
    assert g[1, 0] == Cell(1, 0)
    assert g[1, 1] == Cell(1, 1)

def test_cells_has_correct_n_e_s_w_attributes():
    g = Grid(5,5)
    assert g[3,3].north == g[2,3]
    assert g[3,3].east  == g[3,4]
    assert g[3,3].south == g[4,3]
    assert g[3,3].west  == g[3,2]

def test_cells_may_have_empty_n_e_s_w_attributes():
    g = Grid(5,5)
    assert g[0,0].north is None
    assert g[0,0].east  == g[0,1]
    assert g[4,4].south is None
    assert g[2,0].west  is None

def test_iter():
    g = Grid(2, 2)
    assert list(g) == [Cell(0, 0), Cell(0, 1), Cell(1, 0), Cell(1, 1)]
