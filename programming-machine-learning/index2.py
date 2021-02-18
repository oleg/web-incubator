import numpy as np


def load_pizza():
    return np.loadtxt('pizza.txt', skiprows=1, unpack=True)


def loss(x, y, w, b):
    """mean squared error"""
    return np.average((predict(x, w, b) - y) ** 2)


def predict(x, w, b):
    # x can be an np.array?
    return x * w + b


def gradient(x, y, w, b):
    w_gradient = 2 * np.average(x * (predict(x, w, b) - y))
    b_gradient = 2 * np.average((predict(x, w, b) - y))
    return w_gradient, b_gradient


def train(x, y, iterations, lr):
    w = b = 0
    for i in range(iterations):
        w_gradient, b_gradient = gradient(x, y, w, b)
        w -= w_gradient * lr
        b -= b_gradient * lr
    return w, b


# Tests
def test_train():
    x, y = load_pizza()
    w, b = train(x, y, iterations=20_000, lr=0.001)
    assert w == 1.0811301699901938
    assert b == 13.172267656369339


def test_predict():
    x = np.array([1, 3, 2, 5])
    x_p = predict(x, 10, 0)
    # assert np.testing.assert_equal(x_p, np.array([,,,,]))
    assert list(x_p) == [10, 30, 20, 50]


def test_loss_non_zero():
    x = np.array([1, 2, 2, 5])
    y = np.array([10, 30, 20, 50])
    ls = loss(x, y, 10, 0)
    assert ls == 25


def test_loss_zero():
    x = np.array([1, 3, 2, 5])
    y = np.array([10, 30, 20, 50])
    ls = loss(x, y, 10, 0)
    assert ls == 0
