import ReconnectingWebSocket from "reconnecting-websocket";
import store, { addNotification } from "../redux/";

let hostname = window.location.hostname + ":" + window.location.port;
var socket = new ReconnectingWebSocket("ws://" + hostname + "/ws");

let connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    console.log("receive msg: ", msg);
    store.dispatch(addNotification(msg));
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

const sendMsg = async msg => {
  try {
    socket.send(JSON.stringify(msg));
    return msg;
  } catch (e) {
    // eslint-disable-next-line no-throw-literal
    throw "There is an error. Try again later!!!";
  }
};

const closeSocket = async (code, msg) => {
  socket.close(code, msg);
};

export { connect, sendMsg, closeSocket };
