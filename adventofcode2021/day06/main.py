from __future__ import annotations
import fileinput
from dataclasses import dataclass
from operator import attrgetter


@dataclass
class Fish:
    daysLeft: int
    count: int

    def tick(self) -> [Fish | None]:
        if self.daysLeft == 0:
            self.daysLeft = 6
            return Fish(8, self.count)
        else:
            self.daysLeft -= 1
            return None

    def merge(self, fish):
        if self.daysLeft != fish.daysLeft:
            raise Exception(f"unequal daysLeft {self.daysLeft} and {fish.daysLeft}")
        self.count += fish.count


class Simulation:
    def __init__(self, fishes: list[Fish]):
        fishes.sort(key=attrgetter('daysLeft'))

        prev = fishes[0]
        for f in fishes[1:]:
            if prev.daysLeft == f.daysLeft:
                f.merge(prev)
                prev.daysLeft = -1
            prev = f

        res = [f for f in fishes if f.daysLeft != -1]
        self.fishes = res

    def run_for(self, days) -> Simulation:
        while days > 0:
            days -= 1
            extra = []
            for f in self.fishes:
                self.append(f.tick(), extra)
            for f in extra:
                self.append(f, self.fishes)
        return self

    @staticmethod
    def append(fish: Fish, fishes: list[Fish]):
        if fish is None:
            return
        for f in fishes:
            if f.daysLeft == fish.daysLeft:
                f.merge(fish)
                return
        fishes.append(fish)

    def count(self) -> int:
        return sum([f.count for f in self.fishes])

    def __repr__(self):
        return 'Sim:' + ','.join(','.join(str(x) for x in [f.daysLeft] * f.count) for f in self.fishes)


def run(fishes, days) -> int:
    simulation = Simulation([Fish(int(d), 1) for d in fishes.split(',')])
    simulation.run_for(days)
    return simulation.count()


if __name__ == '__main__':
    line = next(fileinput.input())
    print(run(line, 80))
    print(run(line, 256))


def test_fish_eq():
    assert Fish(8, 1) == Fish(8, 1)


def test_tick():
    f = Fish(0, 10)
    r = f.tick()
    assert r == Fish(8, 10)


def test_merge_different_days_left():
    import pytest
    with pytest.raises(Exception) as e:
        Fish(3, 2).merge(Fish(2, 1))
    assert "unequal daysLeft 3 and 2" in str(e.value)


def test_merge_same_days_left():
    f = Fish(3, 5)
    f.merge(Fish(3, 2))
    assert f == Fish(3, 7)


def test_fish():
    f = Fish(4, 1)
    r = f.tick()
    assert f.daysLeft == 3
    assert r is None

    r = f.tick()
    assert f.daysLeft == 2
    assert r is None

    r = f.tick()
    assert f.daysLeft == 1
    assert r is None

    r = f.tick()
    assert f.daysLeft == 0
    assert r is None

    r = f.tick()
    assert f.daysLeft == 6
    assert r == Fish(8, 1)


def test_simulation_simple_count():
    s = Simulation([Fish(2, 1)])
    assert s.count() == 1
    s = Simulation([Fish(1, 1), Fish(2, 1), Fish(1, 1)])
    assert s.count() == 3


def test_simulation_simple_count_2():
    s = Simulation([Fish(3, 4), Fish(2, 1), Fish(1, 2)])
    assert str(s) == 'Sim:1,1,2,3,3,3,3'


def test_simulation_run_for_10_days():
    def s(): return Simulation([Fish(d, 1) for d in [3, 4, 3, 1, 2]])

    assert str(s().run_for(0)) == 'Sim:1,2,3,3,4'
    assert str(s().run_for(1)) == 'Sim:0,1,2,2,3'
    assert str(s().run_for(2)) == 'Sim:6,0,1,1,2,8'
    assert str(s().run_for(3)) == 'Sim:5,6,0,0,1,7,8'
    assert str(s().run_for(10)) == 'Sim:5,6,0,0,1,0,1,2,2,3,7,8'
    assert str(s().run_for(18)) == 'Sim:4,5,6,6,0,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8'


def test_data():
    result = run("3,4,3,1,2", 80)
    assert result == 5934
