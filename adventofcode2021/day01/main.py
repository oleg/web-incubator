import collections
import fileinput
from collections.abc import Iterable
from itertools import islice


def count_increased(mesures: Iterable[int]) -> int:
    count = 0
    for window in sliding_window(mesures, 2):
        if window[1] > window[0]:
            count += 1
    return count


def count_increased_windows(mesures: Iterable[int]) -> int:
    count = 0
    for window in sliding_window(mesures, 4):
        if sum(window[1:4]) > sum(window[0:3]):
            count += 1
    return count


def sliding_window(iterable, n):
    it = iter(iterable)
    window = collections.deque(islice(it, n), maxlen=n)
    if len(window) == n:
        yield tuple(window)
    for x in it:
        window.append(x)
        yield tuple(window)


def run(seq) -> None:
    mesures = [int(s.rstrip()) for s in seq]
    print(count_increased(mesures))
    print(count_increased_windows(mesures))


if __name__ == '__main__':
    run(fileinput.input())


def test_count_increased():
    assert count_increased([]) == 0
    assert count_increased([1]) == 0
    assert count_increased([1, 1]) == 0
    assert count_increased([1, 2]) == 1
    assert count_increased([2, 1]) == 0
    assert count_increased([199, 200, 208, 210, 200, 207, 240, 269, 260, 263]) == 7


def test_count_increased_windows():
    assert count_increased_windows([]) == 0
    assert count_increased_windows([1]) == 0
    assert count_increased_windows([1, 1]) == 0
    assert count_increased_windows([1, 1, 1]) == 0
    assert count_increased_windows([1, 1, 1, 1]) == 0
    assert count_increased_windows([1, 1, 2, 1]) == 0
    assert count_increased_windows([1, 1, 1, 2]) == 1
    assert count_increased_windows([5, 1, 1, 2]) == 0
    assert count_increased_windows([199, 200, 208, 210, 200, 207, 240, 269, 260, 263]) == 5
