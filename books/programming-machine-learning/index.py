import numpy as np


def load_pizza():
    return np.loadtxt('pizza.txt', skiprows=1, unpack=True)


def predict(X, w):
    return X * w


def loss(X, Y, w):
    """mean squared error"""
    return np.average((predict(X, w) - Y) ** 2)


def train(X, Y, iterations, lr):
    """
    :param X: estimation
    :param Y: ground truth
    :param iterations: number of iterations
    :param lr: learning rate
    :return:
    """
    w = 0
    for i in range(iterations):
        current_loss = loss(X, Y, w)
        # print(f"Iteration {i:4} => Loss {current_loss:6}")

        if loss(X, Y, w + lr) < current_loss:
            w += lr
        elif loss(X, Y, w - lr) < current_loss:
            w -= lr
        else:
            return w
    raise Exception(f"Couldn't converge within {iterations} iterations")


# Tests
def test_predict():
    x = np.array([1, 3, 2, 5])
    x_p = predict(x, 10)
    # assert np.testing.assert_equal(x_p, np.array([,,,,]))
    assert list(x_p) == [10, 30, 20, 50]


def test_loss_non_zero():
    x = np.array([1, 2, 2, 5])
    y = np.array([10, 30, 20, 50])
    ls = loss(x, y, 10)
    assert ls == 25


def test_loss_zero():
    x = np.array([1, 3, 2, 5])
    y = np.array([10, 30, 20, 50])
    ls = loss(x, y, 10)
    assert ls == 0
