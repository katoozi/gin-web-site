import { RECEIVE_NOTIFICATION, DELETE_NOTIFICATION } from "./actionTypes";

export function addNotification(data) {
  return {
    type: RECEIVE_NOTIFICATION,
    text: data.text,
    action_type: RECEIVE_NOTIFICATION
  };
}

export function deleteNotification(text) {
  return {
    type: DELETE_NOTIFICATION,
    text
  };
}
