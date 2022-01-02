import fileinput


def is_lowest(i, j, lines):
    t = int(lines[i][j])
    for (a, b) in zip([i - 1, i + 1, i, i], [j, j, j - 1, j + 1]):
        if a in range(len(lines)):
            if b in range(len(lines[a])):
                if int(lines[a][b]) <= t:
                    return False
    return True


def run(lines) -> int:
    lowest = []
    for i in range(len(lines)):
        for j in range(len(lines[i])):
            if is_lowest(i, j, lines):
                lowest.append(int(lines[i][j]) + 1)
    return sum(lowest)


if __name__ == '__main__':
    print(run([x.strip() for x in fileinput.input()]))

testdata = """2199943210
3987894921
9856789892
8767896789
9899965678
"""


def test_data():
    result = run(testdata.splitlines())
    assert result == 15
