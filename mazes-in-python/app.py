import click

from binary_tree import BinaryTree
from grid import Grid


@click.command()
@click.option('--height', default=10, help='height of the maze')
@click.option('--width', default=10, help='with of the maze')
@click.option('--direction', default="north", help='direction')
@click.option('--seed', help='seed', type=int)
def main(height, width, direction, seed):
    grid = Grid(height, width)
    BinaryTree(seed).generate(direction, grid)
    click.echo(str(grid))


if __name__ == '__main__':
    main()
