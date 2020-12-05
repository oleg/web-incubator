const suits = ['red', 'blue', 'black', 'gold'];
const ranks = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13];

function createDeck() {
    const dominoes = [];

    for (let suit of suits) {
        for (let rank of ranks) {
            dominoes.push(new Domino(suit, rank))
        }
    }

    return [...dominoes, ...dominoes]
}

function getRandom(max) {
    return Math.floor(Math.random() * max);
}

export class Deck {
    constructor() {
        this.dominoes = createDeck()
    }

    getNRandomDominoes(n) {
        const result = []
        for (let i = 0; i < n && this.dominoes.length > 0; i++) {
            const index = getRandom(this.dominoes.length)
            const value = this.dominoes.splice(index, 1)[0];
            result.push(value)
        }
        return result
    }
}

Deck.DECK_SIZE = 4 * 13 * 2;

export class Domino {
    constructor(suit, rank) {
        this.suit = suit
        this.rank = rank
    }
}

export class Player {
    constructor(name) {
        this.name = name
        this.dominoes = []
    }

    take(dominoes) {
        this.dominoes = this.dominoes.concat(dominoes)
    }
}

export class Game {
    constructor(...playerNames) {
        this.deck = new Deck()
        this.players = playerNames.map(name => new Player(name))
        this.players.forEach(player => player.take(this.deck.getNRandomDominoes(14)))
    }
}
