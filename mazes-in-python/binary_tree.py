from __future__ import annotations
from random import Random
from operator import attrgetter

from grid import Grid


class BinaryTree:
    dirs = ['north', 'east', 'south', 'west']
    fs = [attrgetter(d) for d in dirs]

    def __init__(self, seed: int | None = None) -> None:
        self.rnd = Random(seed)

    def generate(self, direction: str, grid: Grid):
        index = self.dirs.index(direction)
        ln = len(self.fs)
        for cell in grid:
            cells = [self.fs[(index + i) % ln](cell) for i in [0, 1]]
            if ns := [v for v in cells if v]:
                if n := self.rnd.choice(ns):
                    cell.link(n)


def test_north():
    g = Grid(5, 5)
    tr = BinaryTree(6)
    tr.generate('north', g)
    # @formatter:off
    assert str(g) == \
"""
+---+---+---+---+---+
|                   |
+   +   +---+---+   +
|   |   |           |
+   +---+---+   +   +
|   |           |   |
+   +   +---+   +   +
|   |   |       |   |
+---+   +---+---+   +
|       |           |
+---+---+---+---+---+
"""
# @formatter:on


def test_east():
    g = Grid(5, 5)
    tr = BinaryTree(6)
    tr.generate('south', g)
    # @formatter:off
    assert str(g) == \
"""
+---+---+---+---+---+
|           |   |   |
+   +---+---+   +   +
|               |   |
+   +---+---+---+   +
|       |       |   |
+   +---+   +---+   +
|       |           |
+   +---+   +---+---+
|                   |
+---+---+---+---+---+
"""
# @formatter:on
