import { call, put, takeEvery } from "redux-saga/effects";
import { sendMsg } from "../websocket";
import { SEND_NOTIFICATION_REQUESTED, RECEIVE_NOTIFICATION } from "./";

function* sendMessage(action) {
  try {
    const msg = yield call(sendMsg, action.text.text);
    yield put({ type: RECEIVE_NOTIFICATION, text: msg });
  } catch (e) {
    yield put({
      type: e.type,
      text: e.message,
      name: e.name
    });
  }
}

function* rootSagas() {
  yield takeEvery(SEND_NOTIFICATION_REQUESTED, sendMessage);
}

export default rootSagas;
