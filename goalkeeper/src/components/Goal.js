import React from 'react';
import {Badge, Col, Progress, Row} from 'reactstrap';

const Goal = props => {
    return (
        <Row>
            <Col xs="4">{props.goalName} <Badge color="danger">{props.epic}</Badge></Col>
            <Col>{props.dateStarted}</Col>
            <Col><Progress value={props.percentComplete}>{props.percentComplete}%</Progress></Col>
        </Row>
    );
};

// Goal.propTypes = {
//     goalName: PropTypes.string.isRequired,
//     percentComplete: PropTypes.number.isRequired,
// };

export default Goal;
