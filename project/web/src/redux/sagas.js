import { call, put, takeEvery } from "redux-saga/effects";
import { sendMsg } from "../websocket";
import {
  SEND_NOTIFICATION_REQUESTED,
  RECEIVE_NOTIFICATION,
  SEND_NOTIFICATION_FAILED
} from "./";

function* sendMessage(action) {
  try {
    const msg = yield call(sendMsg, action.text.text);
    yield put({ type: RECEIVE_NOTIFICATION, text: msg });
  } catch (e) {
    yield put({ type: SEND_NOTIFICATION_FAILED, text: e });
  }
}

function* rootSagas() {
  yield takeEvery(SEND_NOTIFICATION_REQUESTED, sendMessage);
}

export default rootSagas;
