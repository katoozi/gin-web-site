import { RECEIVE_NOTIFICATION, DELETE_NOTIFICATION } from "./actionTypes";


const initialState = {
  notifications: []
};


export function notificationReducer(state = initialState, action) {
  switch (action.type) {
    case RECEIVE_NOTIFICATION:
      return Object.assign({}, state, {
        notifications: [...state.notifications, action.text]
      });
    case DELETE_NOTIFICATION:
      return Object.assign({}, state, {
        notifications: [...state.notifications.filter((notification) => notification !== action.text)]
      });
    default:
      return state;
  }
}
