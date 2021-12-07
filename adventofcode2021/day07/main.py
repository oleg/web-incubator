from __future__ import annotations
import fileinput
from statistics import median


def run(positions_str: str) -> int:
    positions = [int(d) for d in positions_str.split(',')]
    m = median(positions)
    return sum([abs(p - m) for p in positions])


if __name__ == '__main__':
    line = next(fileinput.input())
    print(run(line))


def test_sample():
    result = run("16,1,2,0,4,2,7,1,2,14")
    assert result == 37
