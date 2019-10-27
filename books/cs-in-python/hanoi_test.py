from typing import TypeVar, Generic, List
T = TypeVar('T')

class Tower(Generic[T]):
    def __init__(self) -> None:
        self._container: List[T] = []

    def push(self, item: T) -> None:
        if self._container and self._container[-1] > item:
            raise ValueError
        self._container.append(item)

    def pop(self) -> T:
        return self._container.pop()

    def __len__(self) -> int:
        return len(self._container)
    
    def __repr__(self) -> str:
        return repr(self._container)


class Hanoi():
    def __init__(self, num_discs: int) -> None:
        self.num_discs = num_discs
        self.tower_a= Tower()
        self.tower_b = Tower()
        self.tower_c = Tower()
        for i in range(1, self.num_discs + 1):
            self.tower_a.push(i)

    def solve(self) -> None:
        '''move all from a to c'''
        self.move(self.tower_a, self.tower_c, self.tower_b, self.num_discs)

    def move(self, frm: Tower[int], to: Tower[int], tmp: Tower[int], discs: int) -> None:
        if discs == 1:
            to.push(frm.pop())
        else:
            # A[1..n] -> B
            self.move(frm, tmp, to, discs - 1)
            # A[0]    -> C
            self.move(frm, to, tmp, 1)
            # B[0..n] -> C
            self.move(tmp, to, frm, discs - 1)
            
    def __repr__(self) -> str:
        return f"Hanoi({self.tower_a!r}, {self.tower_b!r}, {self.tower_c!r})"


# Tests
import pytest

def test_tower_push():
    t = Tower()
    t.push(10)
    assert t.pop() == 10
    
    t.push(11)
    assert t.pop() == 11

def test_tower_repr():
    t = Tower()
    t.push(7)
    assert repr(t) == '[7]'

def test_throws_error_if_adds_smaller_value():
    t = Tower()
    t.push(7)
    with pytest.raises(ValueError):
        t.push(6)
    
def test_move_1():
    hanoi = Hanoi(4)
    hanoi.move(hanoi.tower_a, hanoi.tower_b, hanoi.tower_c, 1)
    assert repr(hanoi) == 'Hanoi([1, 2, 3], [4], [])'
    
    hanoi.move(hanoi.tower_b, hanoi.tower_c, hanoi.tower_a, 1)
    assert repr(hanoi) == 'Hanoi([1, 2, 3], [], [4])'

def test_solution_1():
    hanoi = Hanoi(1)
    assert repr(hanoi) == 'Hanoi([1], [], [])'
    hanoi.solve()
    assert repr(hanoi) == 'Hanoi([], [], [1])'

def test_solution_2():
    hanoi = Hanoi(2)
    assert repr(hanoi) == 'Hanoi([1, 2], [], [])'
    hanoi.solve()
    assert repr(hanoi) == 'Hanoi([], [], [1, 2])'
    
def test_solution_3():
    hanoi = Hanoi(3)
    assert repr(hanoi) == 'Hanoi([1, 2, 3], [], [])'    
    hanoi.solve()
    assert repr(hanoi) == 'Hanoi([], [], [1, 2, 3])'

def test_solution_10():
    hanoi = Hanoi(10)
    assert repr(hanoi) == 'Hanoi([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], [], [])'    
    hanoi.solve()
    assert repr(hanoi) == 'Hanoi([], [], [1, 2, 3, 4, 5, 6, 7, 8, 9, 10])'
