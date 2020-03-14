import random
from grid import Grid

def binary_tree(grid: Grid):
    for cell in grid:
        neighbor = random.choice([cell.north, cell.east])
        if neighbor:
            cell.link(neighbor)


if __name__ == "__main__":
    g = Grid(5, 5)
    binary_tree(g)
