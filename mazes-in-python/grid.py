from __future__ import annotations

from dataclasses import dataclass, field

from cell import Cell


@dataclass(init=False)
class Grid:
    rows: int
    columns: int
    grid: list[list[Cell]] = field(default_factory=list, compare=False)

    def __init__(self, rows, columns):
        self.rows = rows
        self.columns = columns
        self.grid = self.prepare_grid()
        self.configure_cells()

    def prepare_grid(self):
        return [[Cell(r, c) for c in range(0, self.columns)]
                for r in range(0, self.rows)]

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
        for row in self.grid:
            for item in row:
                yield item

    def __str__(self) -> str:
        return GridPrinter().print(self)


class GridPrinter:
    def print(self, grid: Grid) -> str:
        if not grid.columns or not grid.rows:
            return ""
        res = "\n"
        for row in grid.grid:
            top = []
            east = ["|"]
            for c in row:
                top.append((' ' if c.is_linked(c.north) else '-') * 3)
                east.append(' ' if c.is_linked(c.east) else '|')
            res += self.__wrap__(top)
            res += "   ".join(east) + "\n"
        res += self.__wrap__(["-" * 3] * grid.columns)
        return res

    @staticmethod
    def __wrap__(line: list) -> str:
        return "+" + "+".join(line) + "+\n"


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
    g = Grid(5, 5)
    assert g[3, 3].north == g[2, 3]
    assert g[3, 3].east == g[3, 4]
    assert g[3, 3].south == g[4, 3]
    assert g[3, 3].west == g[3, 2]


def test_cells_may_have_empty_n_e_s_w_attributes():
    g = Grid(5, 5)
    assert g[0, 0].north is None
    assert g[0, 0].east == g[0, 1]
    assert g[4, 4].south is None
    assert g[2, 0].west is None


def test_iter():
    g = Grid(2, 2)
    assert list(g) == [Cell(0, 0), Cell(0, 1), Cell(1, 0), Cell(1, 1)]


def test_str_empty_0x0():
    assert str(Grid(0, 0)) == ""


def test_str_empty_0x1():
    assert str(Grid(0, 1)) == ""


def test_str_empty_1x0():
    assert str(Grid(1, 0)) == ""


def test_str_single_1x1():
    # @formatter:off
    assert str(Grid(1, 1)) == \
"""
+---+
|   |
+---+
"""
# @formatter:on


def test_str_complete_4x4():
    # @formatter:off
    assert str(Grid(4, 4)) == \
"""
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
|   |   |   |   |
+---+---+---+---+
"""
# @formatter:on


def test_show_4x4_complex():
    g = Grid(4, 4)

    g[0, 0].link(g[0, 1])
    g[0, 1].link(g[0, 2])
    g[0, 2].link(g[0, 3])

    g[1, 0].link(g[0, 0])
    g[1, 1].link(g[0, 1])
    g[1, 1].link(g[1, 2])
    g[1, 3].link(g[0, 3])

    g[2, 0].link(g[2, 1])
    g[2, 1].link(g[2, 2])
    g[2, 2].link(g[1, 2])
    g[2, 2].link(g[2, 3])

    g[3, 0].link(g[2, 0])
    g[3, 0].link(g[3, 1])
    g[3, 2].link(g[2, 2])
    g[3, 3].link(g[2, 3])

    # @formatter:off
    assert str(g) == \
"""
+---+---+---+---+
|               |
+   +   +---+   +
|   |       |   |
+---+---+   +---+
|               |
+   +---+   +   +
|       |   |   |
+---+---+---+---+
"""
# @formatter:on
