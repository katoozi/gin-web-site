import { combineReducers, createStore, applyMiddleware } from "redux";
import * as reducers from "./reducers";
import createSagaMiddleware from "redux-saga";
import rootSagas from "./sagas";

// create the saga middleware
const sagaMiddleware = createSagaMiddleware();

const rootReducer = combineReducers({
  notification: reducers.notificationReducer
});

const store = createStore(rootReducer, applyMiddleware(sagaMiddleware));

sagaMiddleware.run(rootSagas);

export default store;
export * from "./actions";
export * from "./actionTypes";
export * from "./reducers";
export * from "./sagas";
