def fib(x: int) -> int:
    if x < 2: return x
    return fib(x-1) + fib(x-2)

# [0, 1, 2, 3, 4, 5, 6, 7, 8]
# [0, 1, 1, 2, 3, 5, 8, 13, 21]    

def test_fib_zero():
    assert fib(0) == 0

def test_fib_one():
    assert fib(1) == 1
    
def test_fib_two():
    assert fib(2) == 1

def test_fib_three():
    assert fib(3) == 2

def test_fib_four():
    assert fib(4) == 3
    
def test_fib_five():
    assert fib(5) == 5

def test_fib_six():
    assert fib(6) == 8

def test_fib_seven():
    assert fib(7) == 13

def test_fib_eight():
    assert fib(8) == 21
