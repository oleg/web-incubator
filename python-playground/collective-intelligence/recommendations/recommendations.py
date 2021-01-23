import csv
from collections import defaultdict
from math import sqrt


def load_critics():
    critics = defaultdict(dict)
    with open('critics.csv') as file:
        for row in csv.reader(file):
            critics[row[0]][row[1]] = float(row[2])
    return critics


def sim_distance(prefs, person1, person2):
    items = [pow(prefs[person1][item] - prefs[person2][item], 2)
             for item in prefs[person1]
             if item in prefs[person2]]
    if len(items) == 0:
        return 0
    return 1 / (1 + sqrt(sum(items)))
