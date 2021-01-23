def mergesort(seq):
    split_at = len(seq) // 2
    if not split_at: return seq
    a = mergesort(seq[:split_at])
    b = mergesort(seq[split_at:])
    m = merge(a, b)
    return m


def merge(a, b):
    c = []
    while a and b:
        if a[0] < b[0]:
            c.append(a.pop(0))
        else:
            c.append(b.pop(0))
    c.extend(a)
    c.extend(b)
    return c


assert mergesort([1, 4, 2, 3, 6, 5]) == [1, 2, 3, 4, 5, 6]
assert mergesort([8, 5, 4, 3, 2, 1, 9]) == [1, 2, 3, 4, 5, 8, 9]


def quicksort(seq):
    if seq:
        pivot, *tail = seq
        for e in tail:
            if e < pivot:
                pass
            else:
                pass
