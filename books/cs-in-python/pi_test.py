def calculate_pi(n_terms: int) -> float:
    numerator: float = 4.0
    denominator: float = 1.0
    operation: float = 1.0
    pi: float = 0.0

    for _ in range(n_terms):
        pi += operation * (numerator / denominator)
        denominator += 2.0
        operation *= -1.0
    return pi

def test_0():
    assert calculate_pi(0) == 0.0

def test_1():
    assert calculate_pi(1) == 4
    
def test_10():
    assert calculate_pi(10) == 3.0418396189294032

def test_11():
    assert calculate_pi(11) == 3.232315809405594

def test_100():
    assert calculate_pi(100) == 3.1315929035585537
    
def test_10000():
    assert calculate_pi(10000) == 3.1414926535900345

def test_1000000():
    assert calculate_pi(1000000) == 3.1415916535897743
