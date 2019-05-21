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