from __future__ import annotations
from typing import Dict
from dataclasses import dataclass, field

@dataclass(unsafe_hash = True)
class Cell:
    row: int = field(hash = True)
    column: int = field(hash = True)
    links: Dict[Cell, bool] = field(default_factory = dict, init = False, compare = False)
    north: Cell = field(default = None, init = False, compare = False)
    east: Cell = field(default = None, init = False, compare = False)
    south: Cell = field(default = None, init = False, compare = False)
    west: Cell = field(default = None, init = False, compare = False)

    def is_linked(self, cell: Cell) -> bool:
        return cell in self.links

    def link(self, cell: Cell, bidi = True):
       self.links[cell] = True
       if bidi:
           cell.link(self, False)

    def unlink(self, cell: Cell, bidi = True):
        del self.links[cell]
        if bidi:
            cell.unlink(self, False)

    def neighbors(self):
        return [c for c in [self.north, self.east, self.south, self.west] if c]


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
