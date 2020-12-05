import numpy as np


def load_pizza():
    return np.loadtxt('pizza.txt', skiprows=1, unpack=True)


def train(x, y, iterations, lr):
    """
    :param x: estimation
    :param y: ground truth
    :param iterations: number of iterations
    :param lr: learning rate
    :return:
    """
    w = 0
    b = 0
    for i in range(iterations):
        current_loss = loss(x, y, w, b)
        # print(f"Iteration {i:4} => Loss {current_loss:6}")

        if loss(x, y, w + lr, b) < current_loss:
            w += lr
        elif loss(x, y, w - lr, b) < current_loss:
            w -= lr
        elif loss(x, y, w, b + lr) < current_loss:
            b += lr
        elif loss(x, y, w, b - lr) < current_loss:
            b -= lr
        else:
            return w, b
    raise Exception(f"Could not converge within {iterations} iterations")


def loss(x, y, w, b):
    """mean squared error"""
    predicted = predict(x, w, b)
    return np.average((predicted - y) ** 2)


def predict(x, w, b):
    # x can be an np.array?
    return x * w + b


# Tests
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


def test_predict_data():
    x, y = load_pizza()
    w, b = train(x, y, 10000, 0.01)

    assert w == 1.1000000000000008
    assert b == 12.929999999999769
