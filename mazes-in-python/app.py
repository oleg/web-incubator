import click
from grid import Grid
from binary_tree import BinaryTree
from sidewinder import Sidewinder


@click.command()
@click.option('--height', default=10, help='height of the maze')
@click.option('--width', default=10, help='with of the maze')
@click.option('--direction', default='north', help='direction')
@click.option('--seed', help='seed', type=int)
@click.option('--algorithm', default='binary', type=click.Choice(['binary', 'sidewinder']))
def main(height: int, width: int, direction: str, seed: int, algorithm: str):
    grid = Grid(height, width)
    generator = {'binary': BinaryTree, 'sidewinder': Sidewinder}[algorithm](seed)
    generator.generate(direction, grid)
    click.echo(str(grid))


if __name__ == '__main__':
    main()
