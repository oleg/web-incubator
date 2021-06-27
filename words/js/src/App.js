import React, {useEffect, useRef, useState} from 'react';
import Container from 'react-bootstrap/Container';

import './App.css';

const WordsForm = ({onUpdate}) => {
    const wform = useRef(null);
    const onSubmit = (event) => {
        event.preventDefault()
        fetch('http://localhost:3001/upload', {
            method: 'POST',
            mode: 'cors',
            body: new FormData(wform.current)
        })
            .then(response => response.json())
            .then(
                res => onUpdate(res),
                error => alert(error)
            );
    };
    return <form onSubmit={onSubmit} ref={wform}>
        <input type="file" name="subfile"/>
        <input type="submit" value="Add"/>
        <br/>
    </form>;
}

const Word = ({word, onClick}) => {
    return <div className="word-block" onClick={() => onClick(word)}>
        <span className="word">{word.text}</span>
        <span className="superscript">{word.percentComplete}</span>
    </div>;
}

const ListsList = ({lists, loadList}) => {
    return <div className="list-name-box">
        {lists.map(l =>
        <span className="list-name" key={l.name}>
            <a href="#" key={l.name} onClick={() => loadList(l.name)}>{l.name}</a>
        </span>)}
    </div>;
}

const WordsList = ({words, onClick, onSave}) => {
    const lForm = useRef(null);
    const [listName, setListName] = useState([]);
    const onSubmit = (event) => {
        event.preventDefault()
        fetch('http://localhost:3001/lists', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            mode: 'cors',
            body: JSON.stringify({
                name: listName,
                words: words
            })
        })
            // .then(res => res)
            .then(
                res => onSave(),
                error => alert(error)
            );
    };
    return <div className="words-list">
        <form onSubmit={onSubmit} ref={lForm}>
            <input type="text" value={listName} onChange={(e) => setListName(e.target.value)} className="list-name"/>
            <input type="submit" value="Save"/>
        </form>
        <div>
            {words.map(w => <Word word={w} key={w.text} onClick={onClick}/>)}
        </div>
    </div>;
}

const RemoveList = ({lists, onChange}) => {
    return <div>
        list to remove <select onChange={(e)=>onChange(e.target.value)}>
            {lists.map(l => <option value={l.name} key={l.name}>{l.name}</option>)}
        </select>
        {/*<input type="submit" value="remove"/>*/}
    </div>;
}

const sortWords = a => {
    a.sort((a, b) => b.freq - a.freq)
    return a
}

const App = () => {
    const [words, setWords] = useState([]);
    const moveToWords = (word) => {
        setWords(sortWords(words.concat(word)))
        setBuffer(sortWords(buffer.filter(x => x !== word)))
    }
    const [buffer, setBuffer] = useState([]);
    const moveToBuffer = (word) => {
        setBuffer(sortWords(buffer.concat(word)))
        setWords(sortWords(words.filter(x => x !== word)))
    }

    const loadList = (list) => {
        fetch("http://localhost:3001/words?" + new URLSearchParams({name: list}), {
            mode: "cors",
        })
            .then(res => res.json())
            .then(
                res => setWords(res),
                error => alert(error),
            )
    }
    const loadListByName = async (list) => {
        const response = await fetch("http://localhost:3001/words?" + new URLSearchParams({name: list}), {mode: "cors"})
        return await response.json();
    }
    const loadLists = () => {
        fetch("http://localhost:3001/lists", {
            mode: "cors",
        })
            .then(res => res.json())
            .then(
                res => setLists(res.lists),
                error => alert(error),
            )
    }
    const [lists, setLists] = useState([])
    useEffect(loadLists, []);

    return (
        <Container className="p-3">
            <ListsList lists={lists} loadList={loadList}/>
            <br/>
            <WordsForm onUpdate={setWords}/>
            <br/>
            <span>buffer</span>
            <WordsList words={buffer} onClick={moveToWords} onSave={loadLists} key="w1"/>
            <br/>
            <span>selected</span>
            <RemoveList lists={lists} onChange={(v) => {
                loadListByName(v)
                    .then(w => w.map( x => x.text))
                    .then(w => words.filter(x => !w.includes(x.text)))
                    .then(w => setWords(w))
            }}/>
            <WordsList words={words} onClick={moveToBuffer} onSave={loadLists} kye="w2"/>
        </Container>
    );
};

export default App;
