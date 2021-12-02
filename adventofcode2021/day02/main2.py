import dataclasses
import fileinput


@dataclasses.dataclass
class Sub:
    aim: int = 0
    depth: int = 0
    horizontal: int = 0

    def position(self):
        return self.depth * self.horizontal


class Command:
    def __init__(self, cmd: str):
        self.depth = 0
        self.horizontal = 0
        match (cmd.rstrip().split(" ")):
            case ["down", v]:
                self.depth = int(v)
            case ["up", v]:
                self.depth = -int(v)
            case ["forward", v]:
                self.horizontal = int(v)
            case ['']:
                pass
            case _:
                raise Exception("Unknown command")

    def exec(self, sub: Sub):
        sub.aim += self.depth
        sub.horizontal += self.horizontal
        sub.depth += sub.aim * self.horizontal


def run(seq) -> int:
    sub = Sub()
    for s in seq:
        Command(s).exec(sub)
    return sub.position()


if __name__ == '__main__':
    print(run(fileinput.input()))


def test_forward():
    command = Command("forward 5")
    assert command.depth == 0
    assert command.horizontal == 5


def test_down():
    command = Command("down 3")
    assert command.depth == 3
    assert command.horizontal == 0


def test_up():
    command = Command("up 5")
    assert command.depth == -5
    assert command.horizontal == 0


def test_data():
    testdata = """forward 5
down 5
forward 8
up 3
down 8
forward 2
"""
    result = run(testdata.split("\n"))
    assert result == 900
