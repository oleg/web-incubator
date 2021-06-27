import React, {useEffect, useState} from 'react';
import {Row, Col, ProgressBar, Container, Badge} from 'react-bootstrap';

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

    return (
        <Container className="p-3">
            <ItemsList items={items}/>
        </Container>
    );
};

export default App;
