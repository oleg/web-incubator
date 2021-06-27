import React, {useEffect, useState} from 'react';
import {Badge, Col, Container, ProgressBar, Row, Form} from 'react-bootstrap';
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
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
    const [startDate, setStartDate] = useState(new Date());

    const onItemCreate = (event) => {
        event.preventDefault()
        let data = {
            "dateStarted": startDate.toISOString().split('T')[0],
            "tags": [],
            "title": title,
            "percentComplete": percentComplete
        }
        // console.log(JSON.stringify(data))

        fetch('http://localhost:3001/items', {
            method: 'POST',
            mode: 'cors',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(data)
        })
            .then(response => response.json())
            .then(
                res => {
                    loadItems()
                    setTitle("")
                    setPercentComplete("")
                    setStartDate(new Date())
                },
                error => alert(error)
            );
    };

    return (
        <Container className="p-3">
            <ItemsList items={items}/>
            <Form onSubmit={onItemCreate}>
                <input type="text" onChange={(e) => setTitle(e.target.value)} value={title} placeholder="Title"/>
                <br/>
                <input type="text" onChange={(e) => setPercentComplete(e.target.value)} value={percentComplete}
                       placeholder="Percent Complete"/>
                <br/>
                <DatePicker selected={startDate} onChange={(date) => setStartDate(date)} dateFormat="yyyy-MM-dd"/>
                <br/>
                <button type="submit">Add</button>
            </Form>
        </Container>
    );
};

export default App;
