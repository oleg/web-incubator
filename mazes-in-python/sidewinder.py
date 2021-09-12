from __future__ import annotations
from random import Random
from grid import Grid
from cell import Cell


class Sidewinder:
    # dirs = ['north', 'east', 'south', 'west']
    # fs = [attrgetter(d) for d in dirs]
    #
    def __init__(self, seed: int | None = None) -> None:
        self.rnd = Random(seed)

    def generate(self, direction: str, grid: Grid):
        for row in grid.grid:
            run: list[Cell] = []
            for cell in row:
                run.append(cell)
                should_close_out = (cell.east is None) or ((cell.north is not None) and self.rnd.getrandbits(1))
                if should_close_out:
                    member = run[self.rnd.randint(0, len(run) - 1)]
                    if member.north is not None:
                        member.link(member.north)
                    run = []
                else:
                    if cell.east is not None:  # todo only to satisfy mypy?
                        cell.link(cell.east)


def test_north():
    g = Grid(5, 5)
    tr = Sidewinder(6)
    tr.generate('north', g)
    # @formatter:off
    assert str(g) == \
"""
+---+---+---+---+---+
|                   |
+   +---+   +   +---+
|   |       |       |
+---+   +   +   +---+
|       |   |       |
+   +---+---+   +   +
|           |   |   |
+---+   +   +   +   +
|       |   |   |   |
+---+---+---+---+---+
"""
# @formatter:on
