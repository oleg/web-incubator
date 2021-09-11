import click

from binary_tree import binary_tree
from grid import Grid


@click.command()
@click.option('--height', default=10, help='height of the maze')
@click.option('--width', default=10, help='with of the maze')
def main(height, width):
    grid = Grid(height, width)
    binary_tree(grid)
    click.echo(str(grid))


if __name__ == '__main__':
    main()
