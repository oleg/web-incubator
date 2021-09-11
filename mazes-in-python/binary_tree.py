import random

from grid import Grid


# todo test me
def binary_tree(grid: Grid):
    for cell in grid:
        if ns := [c for c in [cell.north, cell.east] if c]:
            if n := random.choice(ns):
                cell.link(n)
