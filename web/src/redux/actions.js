import {RECEIVE_NOTIFICATION} from "./actionTypes";

export function addNotification(text) {
  return {
    type: RECEIVE_NOTIFICATION,
    text
  };
}
