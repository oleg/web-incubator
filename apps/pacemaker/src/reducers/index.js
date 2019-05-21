import {uniqueId} from '../actions';

const mockGoals = [
    // {goalName: 'Beyond Rhythm Guitar', percentComplete: 0},
    {id: uniqueId(), dateStarted: '2016-10-11', epic: 'Guitar', goalName: 'Complete Technique for Modern Guitar', percentComplete: 6},
    {id: uniqueId(), dateStarted: '2017-10-11', epic: 'Guitar', goalName: 'First Pieces for Classical Guitar', percentComplete: 5},
    {id: uniqueId(), dateStarted: '2019-10-11', epic: 'Guitar', goalName: 'Guitar Fretboard Fluency', percentComplete: 18},
    {id: uniqueId(), dateStarted: '2020-10-11', epic: 'Guitar', goalName: 'First Chord Progressions for Guitar', percentComplete: 7},
    // {id:uniqueId(), dateStarted: '2018-10-11', epic: 'Guitar', goalName: 'Heavy Metal Lead Guitar', percentComplete: 0},
    // {id:uniqueId(), dateStarted: '2018-10-11', epic: 'Guitar', goalName: 'Heavy Metal Rhythm Guitar', percentComplete: 0},
    // {id:uniqueId(), dateStarted: '2018-10-11', epic: 'Guitar', goalName: 'Melodic Rock Soloing for Guitar', percentComplete: 0},
    {id: uniqueId(), dateStarted: '2015-10-11', epic: 'Guitar', goalName: 'Modern Music Theory for Guitarists', percentComplete: 10},
    // {id:uniqueId(), dateStarted: '2018-10-11', epic: 'Guitar', goalName: 'Progressive Metal Guitar', percentComplete: 0},
    // {id:uniqueId(), dateStarted: '2018-10-11', epic: 'Guitar', goalName: 'Rock Guitar Mode Mastery', percentComplete: 0},
    // {id:uniqueId(), dateStarted: '2018-10-11', epic: 'Guitar', goalName: 'Rock Rhythm Guitar Playing', percentComplete: 0},
    // {id:uniqueId(), dateStarted: '2018-10-11', epic: 'Guitar', goalName: 'Sight Reading Mastery for Guitar', percentComplete: 0},
    {id: uniqueId(), dateStarted: '2018-10-11', epic: 'Guitar', goalName: 'The Guitar Finger-Gym', percentComplete: 0},
    {id: uniqueId(), dateStarted: '2018-10-11', epic: 'Programming', goalName: 'React In Action', percentComplete: 20},
    {id: uniqueId(), dateStarted: '2018-10-11', epic: 'Programming', goalName: 'Smalltalk Best Practice Patterns', percentComplete: 17},
    {id: uniqueId(), dateStarted: '2018-10-11', epic: 'Programming', goalName: 'Jenkins 2: Up and Running', percentComplete: 44}
];

export default function goals(state = {goals: mockGoals}, action) {
    if (action.type === 'CREATE_GOAL') {
        return {goals: state.goals.concat(action.payload)}
    }
    return state
}