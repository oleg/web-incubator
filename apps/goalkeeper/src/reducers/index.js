export default function goals(state = {goals: []}, action) {
    if (action.type === 'CREATE_GOAL') {
        return {goals: state.goals.concat(action.payload)}
    }
    if (action.type === 'FETCH_GOALS_SUCCEEDED') {
        return {goals: action.payload.goals}
    }
    return state
}