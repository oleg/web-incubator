import React, {Component} from 'react';
import './App.css';

export default class App extends Component {
    render() {
        return (
            <div className="App">
                <div>{this.props.game.players.map(p => <PlayerView player={p}/>)}</div>
                Deck <DominoesView open={false} dominoes={this.props.game.deck.dominoes}/>
            </div>
        );
    }
}

class PlayerView extends Component {
    render() {
        return <div className="Player">
            <div>{this.props.player.name}</div>
            <DominoesView open={true} dominoes={this.props.player.dominoes}/>
        </div>
    }
}

class DominoesView extends Component {
    render() {
        return <div>{this.props.dominoes.map(d => <DominoView key={d} domino={d} open={this.props.open}/>)}</div>
    }
}

class DominoView extends Component {

    constructor(props, context) {
        super(props, context)
        this.state = {open: props.open}
    }

    turn = () => {
        this.setState({open: !this.state.open})
    }

    render() {
        let [color, text] = this.state.open ? [this.props.domino.suit, this.props.domino.rank] :  ['black', '*']
        return <div className="Domino" style={{color: color}} onClick={this.turn}>{text}</div>
    }
}


