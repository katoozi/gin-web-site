import { RECEIVE_NOTIFICATION } from "./actionTypes";


const initialState = {
  notifications: []
};


export function notificationReducer(state = initialState, action) {
  switch (action.type) {
    case RECEIVE_NOTIFICATION:
      return Object.assign({}, state, {
        notifications: [...state.notifications, action.text]
      });
    default:
      return state;
  }
}
