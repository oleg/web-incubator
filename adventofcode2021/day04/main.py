import fileinput
import re
from dataclasses import dataclass


@dataclass
class N:
    n: int
    selected: bool = False

    def select(self, val):
        if self.n == val:
            self.selected = True

    def value(self):
        if self.selected:
            return 0
        return self.n

    def __repr__(self):
        if self.selected:
            return f"({self.n})"
        return str(self.n)


@dataclass
class Card:
    lines: list[list[N]]

    def complete(self):
        for i in range(0, 5):
            all_selected = True
            for j in range(0, 5):
                all_selected &= self.lines[i][j].selected
            if all_selected:
                return True
        for i in range(0, 5):
            all_selected = True
            for j in range(0, 5):
                all_selected &= self.lines[j][i].selected
            if all_selected:
                return True
        return False

    def select(self, val):
        for line in self.lines:
            for n in line:
                n.select(val)

    def value(self):
        res = 0
        for line in self.lines:
            for n in line:
                res += n.value()
        return res

    def __repr__(self):
        res = ""
        for line in self.lines:
            res += ','.join([str(n) for n in line]) + "\n"
        return res


class Bingo:
    def __init__(self, numbers: str, cards):
        self.numbers = self.parse_numbers(numbers)
        self.cards = self.parse_cards(cards)

    @staticmethod
    def parse_numbers(numbers: str):
        return [int(n) for n in numbers.split(',')]

    @staticmethod
    def parse_cards(cards_iter):
        cards = []
        lines = []
        for c in cards_iter:
            numbers = [N(int(n)) for n in re.findall(r'\d+', c)]
            if len(numbers) != 0:
                lines.append(numbers)
            else:
                cards.append(Card(lines))
                lines = []
        return cards

    def draw(self):
        for n in self.numbers:
            for c in self.cards:
                c.select(n)
                if c.complete():
                    return n * c.value()
        return None


def run(seq) -> int:
    i = iter(seq)
    numbers = next(i)
    next(i)
    bingo = Bingo(numbers, i)
    return bingo.draw()


if __name__ == '__main__':
    print(run(fileinput.input()))


def test_numbers():
    numbers = Bingo.parse_numbers("7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1")
    assert numbers == [7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16,
                       13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1]


def test_cards():
    cards = Bingo.parse_cards("""22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6
""".split("\n"))
    assert cards == [
        Card([[N(22), N(13), N(17), N(11), N(0)],
              [N(8), N(2), N(23), N(4), N(24)],
              [N(21), N(9), N(14), N(16), N(7)],
              [N(6), N(10), N(3), N(18), N(5)],
              [N(1), N(12), N(20), N(15), N(19)]]),
        Card([[N(3), N(15), N(0), N(2), N(22)],
              [N(9), N(18), N(13), N(17), N(5)],
              [N(19), N(8), N(7), N(25), N(23)],
              [N(20), N(11), N(10), N(24), N(4)],
              [N(14), N(21), N(16), N(12), N(6)]])
    ]


def test_data():
    testdata = """7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
 """
    result = run(testdata.split("\n"))
    assert result == 4512
