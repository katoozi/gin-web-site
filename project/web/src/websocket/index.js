import ReconnectingWebSocket from 'reconnecting-websocket';
import Swal from 'sweetalert2/dist/sweetalert2.js';
import 'sweetalert2/src/sweetalert2.scss';

// websocket/index.js
var socket = new ReconnectingWebSocket("ws://127.0.0.1/ws");

let connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    let data = JSON.parse(msg.data);
    if(data.action_type === 'RECEIVE_NOTIFICATION'){
      Swal.fire({
        position: 'top-end',
        type: 'success',
        title: data.text,
        showConfirmButton: false,
        timer: 1500
      });
    }else{
      console.log(msg);
    }
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
