import {
  RECEIVE_NOTIFICATION,
  DELETE_NOTIFICATION,
  SEND_NOTIFICATION_REQUESTED
} from "./actionTypes";

export function addNotification(data) {
  return {
    type: RECEIVE_NOTIFICATION,
    text: data.text
  };
}

export function deleteNotification(text) {
  return {
    type: DELETE_NOTIFICATION,
    text
  };
}

export function sendNotificationRequested(text) {
  return {
    type: SEND_NOTIFICATION_REQUESTED,
    text
  };
}
