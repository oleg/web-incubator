def length(seq):
    return 1 + length(seq[1:]) if seq else 0


assert length([0, 1, 2, 3, 4, 5]) == 6
assert length([0, 1, 2, 3]) == 4


def scope():
    x = 1
    if x == 1:
        y = 10
    else:
        z = 20
    print(y)


scope()


def maker(n):
    def action(x):
        return x * n

    return action


def maker2(n):
    return lambda x: n * x
