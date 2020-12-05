import * as api from '../api';

export function fetchGoalsSucceeded(goals) {
    return {
        type: 'FETCH_GOALS_SUCCEEDED',
        payload: {
            goals
        }
    }
}

export function fetchGoals() {
    return dispatch => {
        api.fetchGoals().then(resp => {
            dispatch(fetchGoalsSucceeded(resp.data))
        })
    }
}


let _id = 1;
export function uniqueId() {
    return _id++;
}


export function createGoal({goalName, percentComplete}) {
    return {
        type: 'CREATE_GOAL',
        payload: {
            id: uniqueId(),
            goalName,
            percentComplete,
            epic: 'other'
        },
    };
}