#!/usr/bin/env python3

import csv
from collections import Counter
from pathlib import Path

import click


@click.group()
def cli():
    pass


@cli.command()
@click.argument('subtitles', type=click.File('r'))
@click.option('--minus')
@click.option('--add-to')
def ask(subtitles, minus, add_to):
    freq = get_words_count(subtitles)
    if minus:
        known = get_words(minus, create_if_missing=True)
        for kw in known:
            del freq[kw]
    selected = []
    for k, v in freq.most_common():
        r = input(k + ': ')
        if r == 's':
            selected.append(k)
        elif r == 'q':
            append_words(add_to, selected)
            return


@cli.command()
@click.argument('subtitles', type=click.File('r'))
@click.option('--minus')
@click.option('--only-words', is_flag=True)
def show(subtitles, minus, only_words):
    freq = get_words_count(subtitles)
    if minus:
        known = get_words(minus, create_if_missing=True)
        for kw in known:
            del freq[kw]
    for k, v in freq.most_common():
        if only_words:
            print(k)
        else:
            print(f"{k} {v}")


@cli.command()
@click.argument('subtitles', type=click.File('r'))
@click.argument('name')
def save(subtitles, name):
    freq = get_words_count(subtitles)
    dest = in_worder(name)
    if dest.exists():
        raise RuntimeError("dest already exist")
    save_dict_count(dest, freq)


def worder_home():
    wh = Path.home() / ".worder"
    #    if not wh.exists():
    #        wh.mkdir()
    wh.mkdir(exist_ok=True)
    return wh


def in_worder(filename):
    return worder_home() / filename


def get_words_count(openfile):
    freq = Counter()
    for line in openfile:
        words = [w for w in line
            .replace('-', ' ')
            .replace('.', ' ')
            .replace('!', ' ')
            .replace('?', ' ')
            .replace(';', ' ')
            .lower()
            .split() if w.isalpha()]
        freq.update(words)
    return freq


def save_dict_count(dest, freq):
    with open(dest, 'w+') as f:
        writer = csv.writer(f)
        for pair in freq.most_common():
            writer.writerow(pair)


def get_words(filename, create_if_missing):
    wf = in_worder(filename)

    if create_if_missing:
        wf.touch(exist_ok=True)

    with open(wf, 'r') as f:
        return [word.rstrip() for word in f]


def append_words(add_to, words):
    wf = in_worder(add_to)
    with open(wf, 'a+') as f:
        for w in words:
            f.write(w + '\n')


# usage:
# ./worder.py show data/Remarque_Arc_de_Triomphe.doc.txt --minus known | less
# ./worder.py ask data/Remarque_Arc_de_Triomphe.doc.txt --minus known --add-to known

# todo commands
# - list all imported/saved counts
# - substruct dictionary from saved count
# show words one by one and ask for translation
# show words one by one and add to list of known / ignored / studing


if __name__ == '__main__':
    cli()
