import React, { Component } from "react";
import "./App.css";
import {
  Button,
  ButtonToolbar,
  Container,
  InputGroup,
  FormControl
} from "react-bootstrap";
import { connect, sendMsg, closeSocket } from "./websocket";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
    this.state = {};
    this.messageHandler = this.messageHandler.bind(this);
    this.send = this.send.bind(this);
  }

  send() {
    let value = this.state.message_input_text;
    console.log(value);
    sendMsg(value);
  }

  close() {
    console.log("Close Socket");
    closeSocket(1000, "close by user");
  }

  messageHandler(input_obj) {
    if (input_obj.keyCode === 13) {
      this.send();
      return
    }
    this.setState({
      message_input_text: input_obj.target.value
    });
  }

  render() {
    return (
      <Container>
        <div className="app">
          <ButtonToolbar aria-label="WebSocket Settings" className="app md-3">
            <InputGroup>
              {/* <InputGroup.Prepend>
                <InputGroup.Text id="btnGroupAddon">@</InputGroup.Text>
              </InputGroup.Prepend> */}
              <FormControl
                type="text"
                onKeyUp={this.messageHandler}
                placeholder="message to send"
                aria-label="message"
                aria-describedby="btnGroupAddon"
              />
            </InputGroup>
            <Button variant="outline-primary" onClick={this.send} value="Hit">
              Hit
            </Button>
            <Button variant="outline-danger" onClick={this.close}>
              Close Socket
            </Button>
          </ButtonToolbar>
        </div>
      </Container>
    );
  }
}

export default App;
