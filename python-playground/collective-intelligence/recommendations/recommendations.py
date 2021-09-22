from __future__ import annotations
import csv
from collections import defaultdict
from math import sqrt
import typing


def load_critics() -> dict[str, dict[str, typing.Any]]:
    critics: dict[str, dict[str, typing.Any]] = defaultdict(dict)
    with open('critics.csv') as file:
        for row in csv.reader(file):
            critics[row[0]][row[1]] = float(row[2])
    return critics


# Returns a distance-based similarity score for p1 and p2
def sim_distance(prefs, p1, p2):
    items = [pow(prefs[p1][item] - prefs[p2][item], 2)
             for item in prefs[p1]
             if item in prefs[p2]]
    if len(items) == 0:
        return 0
    return 1 / (1 + sqrt(sum(items)))


# Returns the Pearson correlation coefficient for p1 and p2
def sim_pearson(prefs, p1, p2):
    shared_items = {item for item in prefs[p1] if item in prefs[p2]}
    n = float(len(shared_items))
    if n == 0:
        return 0
    sum1 = sum(prefs[p1][it] for it in shared_items)
    sum2 = sum(prefs[p2][it] for it in shared_items)

    sum1_sqr = sum(pow(prefs[p1][it], 2) for it in shared_items)
    sum2_sqr = sum(pow(prefs[p2][it], 2) for it in shared_items)

    p_sum = sum(prefs[p1][it] * prefs[p2][it] for it in shared_items)

    num = p_sum - sum1 * sum2 / n
    den = sqrt(
        (sum1_sqr - pow(sum1, 2) / n) *
        (sum2_sqr - pow(sum2, 2) / n))

    if den == 0:
        return 0
    return num / den


def top_matches(prefs, person, n=5, similarity=sim_pearson):
    scores = [(similarity(prefs, person, other), other)
              for other in prefs
              if other != person]
    scores.sort()
    scores.reverse()
    return scores[0:n]


def test_sim_distance():
    c = load_critics()
    d = sim_distance(c, 'Lisa Rose', 'Gene Seymour')
    assert d == 0.29429805508554946


def test_sim_parson():
    c = load_critics()
    d = sim_pearson(c, 'Lisa Rose', 'Gene Seymour')
    assert d == 0.39605901719066977


def test_top_matches():
    c = load_critics()
    tm = top_matches(c, 'Lisa Rose')
    assert tm == [(0.9912407071619299, 'Toby'),
                  (0.7470178808339965, 'Jack Matthews'),
                  (0.5940885257860044, 'Mick LaSalle'),
                  (0.5669467095138396, 'Claudia Puig'),
                  (0.40451991747794525, 'Michael Phillips')]
