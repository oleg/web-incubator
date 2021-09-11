from __future__ import annotations

from dataclasses import dataclass, field


@dataclass(unsafe_hash=True, repr=False)
class Cell:
    row: int = field(hash=True)
    column: int = field(hash=True)

    north: Cell | None = field(default=None, init=False, compare=False)
    east: Cell | None = field(default=None, init=False, compare=False)
    south: Cell | None = field(default=None, init=False, compare=False)
    west: Cell | None = field(default=None, init=False, compare=False)

    links: dict[Cell, bool] = field(default_factory=dict, init=False, compare=False)

    def is_linked(self, cell: Cell | None) -> bool:
        return cell in self.links

    def link(self, cell: Cell, bidi=True):
        self.links[cell] = True
        if bidi:
            cell.link(self, False)

    def unlink(self, cell: Cell, bidi=True):
        del self.links[cell]
        if bidi:
            cell.unlink(self, False)

    def neighbors(self):
        return [c for c in [self.north, self.east, self.south, self.west] if c]

    def __repr__(self):
        # return f"Cell({self.row},{self.column})"
        return "Cell([{},{}] [n={}, e={}, s={}, w={}] [{}])" \
            .format(self.row, self.column, self.north, self.east, self.south, self.west, self.__links_str__())

    def __str__(self):
        return "C({},{})".format(self.row, self.column)

    def __links_str__(self):
        return ', '.join([str(kv[0]) for kv in self.links.items() if kv[1]])


def test_properties():
    c = Cell(1, 2)
    assert c.row == 1
    assert c.column == 2


def test_neighbors():
    n = Cell(0, 1)
    e = Cell(1, 2)
    s = Cell(2, 1)
    w = Cell(1, 0)

    m = Cell(1, 1)
    m.north = n
    m.east = e
    m.south = s
    m.west = w

    assert m.north == n
    assert m.east == e
    assert m.south == s
    assert m.west == w


def test_links_is_empty_by_default():
    c = Cell(1, 1)
    assert c.links == {}


def test_is_linked_returns_false_for_unlinked_cells():
    a = Cell(0, 0)
    b = Cell(0, 1)
    assert not a.is_linked(b)
    assert not b.is_linked(a)


def test_is_linked_returns_true_if_linked():
    a = Cell(0, 0)
    b = Cell(0, 1)
    a.link(b)
    assert a.is_linked(b)
    assert b.is_linked(a)


def test_is_linked_returns_false_after_unlink():
    a = Cell(0, 0)
    b = Cell(0, 1)
    a.link(b)
    a.unlink(b)
    assert not a.is_linked(b)
    assert not b.is_linked(a)


def test_neighbors_north_and_south():
    a1 = Cell(0, 1)
    a2 = Cell(2, 1)
    m = Cell(1, 1)

    m.north = a1
    m.south = a2

    assert m.neighbors() == [a1, a2]


def test_neighbors_all():
    n = Cell(0, 1)
    e = Cell(1, 2)
    s = Cell(2, 1)
    w = Cell(1, 0)

    m = Cell(1, 1)
    m.north = n
    m.east = e
    m.south = s
    m.west = w

    assert m.neighbors() == [n, e, s, w]


def test_repr_single():
    assert repr(Cell(0, 0)) == "Cell([0,0] [n=None, e=None, s=None, w=None] [])"


def test_repr_with_neighbor():
    w = Cell(0, 0)
    e = Cell(0, 1)
    w.east = e
    e.west = w
    assert repr(w) == "Cell([0,0] [n=None, e=C(0,1), s=None, w=None] [])"
    assert repr(e) == "Cell([0,1] [n=None, e=None, s=None, w=C(0,0)] [])"


def test_repr_with_neighbor_linked():
    n = Cell(0, 0)
    s = Cell(1, 0)
    n.south = s
    s.north = n
    n.link(s)
    assert repr(n) == "Cell([0,0] [n=None, e=None, s=C(1,0), w=None] [C(1,0)])"
    assert repr(s) == "Cell([1,0] [n=C(0,0), e=None, s=None, w=None] [C(0,0)])"


def test_str():
    assert str(Cell(0, 0)) == "C(0,0)"
    assert str(Cell(5, 2)) == "C(5,2)"


def test_str_with_neighbor():
    n = Cell(0, 0)
    e = Cell(0, 1)
    n.east = e
    e.north = n
    assert str(n) == "C(0,0)"


def test_str_with_neighbor_linked():
    n = Cell(0, 0)
    e = Cell(0, 1)
    n.east = e
    e.north = n
    n.link(e)
    assert str(n) == "C(0,0)"
