import fileinput
from collections import defaultdict


def run(seq) -> int:
    d = defaultdict(lambda: defaultdict(lambda: 0))
    for line in seq:
        line = line.rstrip()
        for i, v in enumerate(line):
            d[i][v] += 1
    lst = [''] * len(d)
    for k, v in d.items():
        lst[k] = '0' if v['0'] > v['1'] else '1'
    g = ''.join(lst)
    gamma = int(g, 2)
    delta = gamma ^ int('1' * len(d), 2)
    return gamma * delta


if __name__ == '__main__':
    print(run(fileinput.input()))


def test_data():
    testdata = """00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
"""
    result = run(testdata.split("\n"))
    assert result == 198
