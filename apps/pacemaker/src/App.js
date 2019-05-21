import React, {Component} from 'react';
import {Container, Nav, NavItem, NavLink} from 'reactstrap';
import GoalsPage from './components/GoalsPage'
import {connect} from 'react-redux';
import {createGoal} from './actions';

class App extends Component {

    onCreateGoal = ({goalName, percentComplete}) => {
        this.props.dispatch(createGoal({goalName, percentComplete}));
    }

    render() {
        return (
            <Container>
                <Nav>
                    <NavItem><NavLink href="#">домашняя страница</NavLink></NavItem>
                    <NavItem><NavLink href="#">другая страница</NavLink></NavItem>
                    <NavItem><NavLink href="#">страница номер 3</NavLink></NavItem>
                </Nav>
                <GoalsPage goals={this.props.goals} onCreateGoal={this.onCreateGoal}/>
            </Container>
        );
    }
}

function mapStateToProps(state) {
    return {
        goals: state.goals
    }
}

export default connect(mapStateToProps)(App);
