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


const Word = ({word}) =>
    <div className="word-block">
        <span className="word">{word.text}</span>
        <span className="superscript">{word.freq}</span>
    </div>

const WordsList = ({words}) =>
    <div>{words.map(w => <Word word={w} key={w.text}/>)}</div>

const App = () => {
    const [words, setWords] = useState([]);
    useEffect(() => {
        fetch("http://localhost:3001/words")
            .then(res => res.json())
            .then(
                result => setWords(result),
                error => alert(error),
            )
    }, []);

    return (
        <Container className="p-3">
            <WordsForm onUpdate={setWords}/>
            <WordsList words={words}/>
        </Container>
    );
};

export default App;
