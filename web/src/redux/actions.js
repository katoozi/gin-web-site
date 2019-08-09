import {RECEIVE_NOTIFICATION, DELETE_NOTIFICATION} from "./actionTypes";

export function addNotification(text) {
  return {
    type: RECEIVE_NOTIFICATION,
    text
  };
}

export function deleteNotification(text) {
  return {
    type: DELETE_NOTIFICATION,
    text
  };
}
