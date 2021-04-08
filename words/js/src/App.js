import React, {useEffect, useRef, useState} from 'react';
import Container from 'react-bootstrap/Container';

import './App.css';

const WordsForm = ({onUpdate}) => {
    const wform = useRef(null);
    const onSubmit = (event) => {
        event.preventDefault()
        fetch('http://localhost:3001/upload', {
            method: 'POST',
            body: new FormData(wform.current)
        })
            .then(response => response.json())
            .then(
                result => onUpdate(result),
                error => alert(error)
            );
    };
    return <form className="words-form" onSubmit={onSubmit} ref={wform}>
        <input type="file" name="subfile"/>
        <input type="submit" value="Add"/>
    </form>;
}


const Word = ({word, onClick}) => {
    return <div className="word-block" onClick={() => onClick(word)}>
        <span className="word">{word.text}</span>
        <span className="superscript">{word.freq}</span>
    </div>;
}

const WordsList = ({words, onClick}) =>
    <div>{words.map(w => <Word word={w} key={w.text} onClick={onClick}/>)}</div>

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

    useEffect(() => {
        fetch("http://localhost:3001/words")
            .then(res => res.json())
            .then(
                result => setWords(result),
                error => alert(error),
            )
    }, []);
    const [buffer, setBuffer] = useState([]);
    const moveToBuffer = (word) => {
        setBuffer(sortWords(buffer.concat(word)))
        setWords(sortWords(words.filter(x => x !== word)))
    }
    return (
        <Container className="p-3">
            <WordsForm onUpdate={setWords}/>
            <WordsList words={buffer} onClick={moveToWords} key="w1"/>
            <hr/>
            <WordsList words={words} onClick={moveToBuffer} kye="w2"/>
        </Container>
    );
};

export default App;
