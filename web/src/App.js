import React, {Component} from "react";
import "./App.css";
import { connect, sendMsg, closeSocket } from "./websocket";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  close() {
    console.log("Close Socket");
    closeSocket(1000, 'close by user');
  }

  render() {
    return (
      <div className="App">
        <button onClick={this.send}>Hit</button>
        <button onClick={this.close}>Close Socket</button>
      </div>
    );
  }
}

export default App;
