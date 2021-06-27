import React, {useEffect, useState} from 'react';
import {Badge, Col, Container, ProgressBar, Row} from 'react-bootstrap';

import './App.css';

const ItemsList = ({items}) => {
    return <div className="items-list">
        {items.map(i => <Item item={i} key={i.id}/>)}
    </div>;
}

const Item = ({item}) => {
    return <Row>
        <Col xs="4">{item.title} <Badge color="danger">{item.tags[0]}</Badge></Col>
        <Col>{item.dateStarted}</Col>
        <Col><ProgressBar now={item.percentComplete}/></Col>
    </Row>;
}

const App = () => {
    const loadItems = () => {
        fetch("http://localhost:3001/items", {
            mode: "cors",
        })
            .then(res => res.json())
            .then(
                res => setItems(res),
                error => alert(error),
            )
    }
    const [items, setItems] = useState([])
    useEffect(loadItems, []);

    const [title, setTitle] = useState([]);
    const [percentComplete, setPercentComplete] = useState([]);

    const onItemCreate = (event) => {
        event.preventDefault()
        console.log(title)
        console.log(percentComplete)
        let data = {
            "dateStarted": "2021-01-02",
            "tags": [],
            "title": title,
            "percentComplete": percentComplete
        }
        console.log(JSON.stringify(data))

        fetch('http://localhost:3001/items', {
            method: 'POST',
            mode: 'cors',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(data)
        })
            .then(response => response.json())
            .then(
                res => loadItems(),
                error => alert(error)
            );
    };

    return (
        <Container className="p-3">
            <ItemsList items={items}/>
            <Col><Row>
                <form onSubmit={onItemCreate}>
                    <input type="text" onChange={ (e) => setTitle(e.target.value)} value={title} placeholder="Title"/>
                    <input type="text" onChange={(e) => setPercentComplete(e.target.value)} value={percentComplete}
                           placeholder="Percent Complete"/>
                    <button type="submit">Add</button>
                </form>
            </Row></Col>
        </Container>
    );
};

export default App;
