import ReconnectingWebSocket from 'reconnecting-websocket';

// websocket/index.js
var socket = new ReconnectingWebSocket("ws://127.0.0.1/ws");

let connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    console.log("receive msg: ", msg);
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = msg => {
  console.log("sending msg: ", msg);
  socket.send(
    JSON.stringify(msg)
  );
};

let closeSocket = (code, msg) => {
  socket.close(code, msg);
}

export { connect, sendMsg, closeSocket };
