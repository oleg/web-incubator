import random

from grid import Grid


def binary_tree(grid: Grid):
    for cell in grid:
        ns = [c for c in [cell.north, cell.east] if c]
        if ns:
            n = random.choice(ns)
            if n:
                cell.link(n)


if __name__ == "__main__":
    g = Grid(15, 15)
    binary_tree(g)
    print(g)
