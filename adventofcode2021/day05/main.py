import fileinput
import re
from collections import defaultdict
from dataclasses import dataclass


@dataclass
class Point:
    x: int
    y: int


@dataclass
class Line:
    fr: Point
    to: Point

    def is_horizontal(self) -> bool:
        return self.fr.y == self.to.y

    def is_vertical(self) -> bool:
        return self.fr.x == self.to.x

    def __iter__(self):
        if self.is_horizontal():
            return self.horizontal_gen()
        if self.is_vertical():
            return self.vertical_gen()
        return iter([])

    def horizontal_gen(self):
        start, end = self.order(self.fr.x, self.to.x)
        for x in range(start, end + 1):
            yield x, self.fr.y

    def vertical_gen(self):
        start, end = self.order(self.fr.y, self.to.y)
        for y in range(start, end + 1):
            yield self.fr.x, y

    @staticmethod
    def order(a, b):
        if a < b:
            return a, b
        return b, a


@dataclass
class Field:
    lines: list[Line]

    def count_dangerous(self) -> int:
        m = defaultdict(lambda: defaultdict(lambda: 0))
        for line in self.lines:
            for x, y in line:
                m[x][y] += 1

        dangerous = 0
        for x, m2 in m.items():
            for y, count in m2.items():
                if count >= 2:
                    dangerous += 1

        return dangerous


def run(seq) -> int:
    field = parse_field(seq)
    return field.count_dangerous()


def parse_field(seq) -> Field:
    lines = []
    for line_str in seq:
        if len(line_str.rstrip()) != 0:
            fr_x, fr_y, to_x, to_y = [int(n) for n in re.findall(r'\d+', line_str)]
            lines.append(Line(Point(fr_x, fr_y), Point(to_x, to_y)))
    return Field(lines)


if __name__ == '__main__':
    print(run(fileinput.input()))


def test_line_iter_horizontal2():
    line = Line(Point(0, 9), Point(5, 9))
    points = [p for p in line]
    assert points == [(0, 9), (1, 9), (2, 9), (3, 9), (4, 9), (5, 9)]


def test_line_iter_horizontal():
    line = Line(Point(10, 4), Point(15, 4))
    points = [p for p in line]
    assert points == [(10, 4), (11, 4), (12, 4), (13, 4), (14, 4), (15, 4)]


def test_line_iter_vertical():
    line = Line(Point(3, 13), Point(3, 9))
    points = [p for p in line]
    assert points == [(3, 9), (3, 10), (3, 11), (3, 12), (3, 13)]


testdata = """0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
 """


def test_parse_field():
    field = parse_field(testdata.split("\n"))
    assert len(field.lines) == 10
    assert field.lines[0] == Line(Point(0, 9), Point(5, 9))


def test_data():
    result = run(testdata.split("\n"))
    assert result == 5
