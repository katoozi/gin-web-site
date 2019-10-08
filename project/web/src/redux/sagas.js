import { call, put, takeEvery } from "redux-saga/effects";
import { sendMsg, MsgException } from "../websocket";
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
    if (e instanceof MsgException) {
      yield put({
        type: SEND_NOTIFICATION_FAILED,
        text: e.message,
        name: e.name
      });
    }
  }
}

function* rootSagas() {
  yield takeEvery(SEND_NOTIFICATION_REQUESTED, sendMessage);
}

export default rootSagas;
