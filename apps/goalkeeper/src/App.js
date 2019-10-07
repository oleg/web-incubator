import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Container, Nav, NavItem, NavLink} from 'reactstrap';
import GoalsPage from './components/GoalsPage'
import {createGoal, fetchGoals} from './actions';

class App extends Component {
    componentDidMount() {
        this.props.dispatch(fetchGoals())
    }

    onCreateGoal = ({goalName, percentComplete}) => {
        this.props.dispatch(createGoal({goalName, percentComplete}));
    }

    render() {
        return (
            <Container>
                <Nav>
                    <NavItem><NavLink href="#">page1</NavLink></NavItem>
                    <NavItem><NavLink href="#">page2</NavLink></NavItem>
                    <NavItem><NavLink href="#">page3</NavLink></NavItem>
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
