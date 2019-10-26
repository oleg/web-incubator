from typing import Generator

def fib(x: int) -> Generator[int, None, None]:
    yield 0
    if x > 0: yield 1
    last: int = 0
    next: int = 1
    for _ in range(1, x):
        last, next = next, last + next
        yield next


# [0, 1, 2, 3, 4, 5, 6, 7,  8]
# [0, 1, 1, 2, 3, 5, 8, 13, 21]    

def test_fib_zero():
    assert list(fib(0)) == [0]

def test_fib_one():
    assert list(fib(1)) == [0, 1]
    
def test_fib_two():
    assert list(fib(2)) == [0, 1, 1]

def test_fib_three():
    assert list(fib(3)) == [0, 1, 1, 2]

def test_fib_four():
    assert list(fib(4)) == [0, 1, 1, 2, 3]
    
def test_fib_five():
    assert list(fib(5)) == [0, 1, 1, 2, 3, 5]

def test_fib_six():
    assert list(fib(6)) == [0, 1, 1, 2, 3, 5, 8]

def test_fib_seven():
    assert list(fib(7)) == [0, 1, 1, 2, 3, 5, 8, 13]

def test_fib_eight():
    assert list(fib(8)) == [0, 1, 1, 2, 3, 5, 8, 13, 21]
