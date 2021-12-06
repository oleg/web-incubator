from __future__ import annotations
import fileinput
from dataclasses import dataclass


@dataclass
class Fish:
    daysLeft: int

    def tick(self) -> [Fish | None]:
        if self.daysLeft == 0:
            self.daysLeft = 6
            return Fish(8)
        else:
            self.daysLeft -= 1
            return None


@dataclass
class Simulation:
    fishes: list[Fish]

    def run_for(self, days) -> Simulation:
        while days > 0:
            days -= 1
            extra = []
            for f in self.fishes:
                if r := f.tick():
                    extra.append(r)
            self.fishes.extend(extra)
        return self

    def count(self) -> int:
        return len(self.fishes)

    def __repr__(self):
        return 'Sim:' + ','.join(str(f.daysLeft) for f in self.fishes)


def run(line, days) -> int:
    simulation = Simulation([Fish(int(d)) for d in line.split(',')])
    simulation.run_for(days)
    return simulation.count()


if __name__ == '__main__':
    line = next(fileinput.input())
    print(run(line, 80))
    #print(run(line, 256))


def test_fish_eq():
    assert Fish(8) == Fish(8)


def test_fish():
    f = Fish(4)
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
    assert r == Fish(8)


def test_simulation_simple_count():
    s = Simulation([Fish(2)])
    assert s.count() == 1
    s = Simulation([Fish(1), Fish(2), Fish(1)])
    assert s.count() == 3


def test_simulation_run_for_10_days():
    def s(): return Simulation([Fish(d) for d in [3, 4, 3, 1, 2]])

    assert str(s().run_for(0)) == 'Sim:3,4,3,1,2'
    assert str(s().run_for(1)) == 'Sim:2,3,2,0,1'
    assert str(s().run_for(2)) == 'Sim:1,2,1,6,0,8'
    assert str(s().run_for(3)) == 'Sim:0,1,0,5,6,7,8'
    assert str(s().run_for(4)) == 'Sim:6,0,6,4,5,6,7,8,8'
    assert str(s().run_for(5)) == 'Sim:5,6,5,3,4,5,6,7,7,8'
    assert str(s().run_for(6)) == 'Sim:4,5,4,2,3,4,5,6,6,7'
    assert str(s().run_for(7)) == 'Sim:3,4,3,1,2,3,4,5,5,6'
    assert str(s().run_for(8)) == 'Sim:2,3,2,0,1,2,3,4,4,5'
    assert str(s().run_for(9)) == 'Sim:1,2,1,6,0,1,2,3,3,4,8'
    assert str(s().run_for(10)) == 'Sim:0,1,0,5,6,0,1,2,2,3,7,8'
    assert str(s().run_for(18)) == 'Sim:6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8'


def test_data():
    result = run("3,4,3,1,2", 80)
    assert result == 5934
