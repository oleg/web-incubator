import {Deck, Domino, Game, Player} from './Game'

//Deck
test('domino has correct suit and rank', () => {
    let domino = new Domino('red', 10)

    expect(domino.suit).toBe('red');
    expect(domino.rank).toBe(10);
});

test('deck has right size', () => {
    let deck = new Deck()

    expect(deck.dominoes.length).toBe(Deck.DECK_SIZE);
});

test('getNRandomDominoes retrieves dominoes of the given size from the deck', () => {
    let deck = new Deck()

    let dominoes1 = deck.getNRandomDominoes(1)
    expect(dominoes1.length).toBe(1);
    expect(deck.dominoes.length).toBe(Deck.DECK_SIZE - 1);

    let dominoes9 = deck.getNRandomDominoes(9)
    expect(dominoes9.length).toBe(9);
    expect(deck.dominoes.length).toBe(Deck.DECK_SIZE - 1 - 9);
});

test('getNRandomDominoes n is more than deck size', () => {
    let deck = new Deck()

    let taken = deck.getNRandomDominoes(Deck.DECK_SIZE + 10)
    expect(taken.length).toBe(Deck.DECK_SIZE);
    expect(deck.dominoes.length).toBe(0);
});


///Player
test('player has name', () => {
    let player = new Player('Oleg')

    expect(player.name).toBe('Oleg');
})
test('player has no dominoes', () => {
    let player = new Player('Zhenya')

    expect(player.dominoes.length).toBe(0);
});
test('player can take a domino', () => {
    let player = new Player('Zhenya')
    player.take([new Domino('black', 10)])

    expect(player.dominoes.length).toBe(1);
    expect(player.dominoes[0]).toEqual(new Domino('black', 10));
});

///Game
test('when game is created it has deck and players', () => {
    let game = new Game('Oleg', 'Zhenya')

    expect(game.deck.dominoes.length > 0).toBeTruthy();
    expect(game.players.map(p => p.name)).toEqual(['Oleg', 'Zhenya']);
});

test('when game is created each player take 14 dominoes from the deck', () => {
    let game = new Game('Oleg', 'Zhenya')

    expect(game.deck.dominoes.length).toBe(Deck.DECK_SIZE - 28);
    expect(game.players.map(p => p.dominoes.length)).toEqual([14, 14]);
});
