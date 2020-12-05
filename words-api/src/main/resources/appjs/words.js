import React from "react";
import ListsBox from "./lists";
import Util from "./util";
import {worderApi} from "./api-client";


class WordsForm extends React.Component {

    constructor(props) {
        super(props);
        Util._bind(this, 'handleSubmit');
    }

    handleSubmit(e) {
        e.preventDefault();
        var word = this.refs.word.value.trim();
        if (!word) {
            return;
        }
        this.props.onWordSubmit({word: word});
        this.refs.word.value = '';
    }

    render() {
        return <form className="words-form" onSubmit={this.handleSubmit}>
            <input type="text" placeholder="Word" ref="word"/>
            <input type="submit" value="Add"/>
        </form>;
    }
}

class Word extends React.Component {
    constructor(props) {
        super(props);
        this.state = {hovered: false, clicked: false};
        Util._bind(this, '_updateState', '_clicked');
    }

    _updateState(isHovered) {
        this.setState({hovered: isHovered});
    }

    _clicked(e) {
        if (this.props.moveWord) {
            this.props.moveWord(this.props.word);
            this.setState({clicked: true});
        }
    }

    render() {
        if (this.state.clicked) {
            return null;
        }
        var hover = this._updateState.bind(this, true);
        var unhover = this._updateState.bind(this, false);
        return <div className="word-block" onMouseEnter={hover} onMouseLeave={unhover} onClick={this._clicked}>
            <span className="word">{this.props.word}</span>
            <span className="superscript">{this.props.count}</span>
        </div>
    }
}

const WordsList = (props) => {
    var words = props.words.map((word) => <Word word={word} count={0} moveWord={props.moveWord} key={word} />);
    return <div className="wordsList">{words}</div>
};

class WordsBox extends React.Component {
    constructor(props) {
        super(props);
        this.state = {words: []};

        Util._bind(this, 'componentWillReceiveProps', 'componentDidMount',
            '_handleWordSubmit', '_loadWordsFromServer');
    }

    componentWillReceiveProps(nextProps) {
        this._loadWordsFromServer(nextProps);
    }

    componentDidMount() {
        this._loadWordsFromServer(this.props);
    }

    _handleWordSubmit(data) {
        worderApi.postJson(
            data,
            worderApi.listUrl(this.props.listName),
            data => {
                this.setState(data)
                this._loadWordsFromServer(this.props)
            });
    }

    _loadWordsFromServer(props) {
        worderApi.getJson(worderApi.listUrl(props.listName), (data) => this.setState(data));
    }

    //_moveWord: function(word) {
    //    //var loggedUser = WorderApi.getLoggedUser();
    //    //WorderApi.postJson({"word": word}, "/api/user/" + loggedUser + "/list/" + "study", function() {});
    //},
    render() {
        return <div className="post">
            <h2 className="title">{this.props.listName}</h2>
            <div className="entry">
                <WordsList words={this.state.words} moveWord={null}/>
                <WordsForm onWordSubmit={this._handleWordSubmit}/>
            </div>
        </div>;
    }
}

export {WordsBox, Word};