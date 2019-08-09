import React from "react";
import {
  Button,
  InputGroup,
  FormControl,
  ButtonToolbar
} from "react-bootstrap";
import ToastList from "./toast_list";
import { sendMsg, closeSocket } from "../websocket";
import { addNotification, deleteNotification } from "../redux/actions";
import { connect } from "react-redux";

class AddNotification extends React.Component {
  state = {
    input_state: false,
    message_input_text: ""
  };
  constructor(props) {
    super(props);
    this.messageHandler = this.messageHandler.bind(this);
    this.send = this.send.bind(this);
    this.delete = this.delete.bind(this);
    this.close = this.close.bind(this);
  }
  send() {
    let value = this.state.message_input_text;
    sendMsg(value);
    this.props.dispatch(addNotification(value));
  }

  delete() {
    let value = this.state.message_input_text;
    this.props.dispatch(deleteNotification(value));
  }

  close() {
    console.log("Close Socket");
    closeSocket(1000, "close by user");
    this.setState({
      input_state: true
    })
  }

  messageHandler(input_obj) {
    if (input_obj.keyCode === 13) {
      this.send();
      return;
    }
    this.setState({
      message_input_text: input_obj.target.value
    });
  }
  render() {
    return (
      <div className="app">
        <ButtonToolbar aria-label="WebSocket Settings" className="app md-3">
          <InputGroup>
            {/* <InputGroup.Prepend>
                <InputGroup.Text id="btnGroupAddon">@</InputGroup.Text>
              </InputGroup.Prepend> */}
            <FormControl
              type="text"
              disabled={this.state.input_state}
              onKeyUp={this.messageHandler}
              placeholder="message to send"
              aria-label="message"
              aria-describedby="btnGroupAddon"
            />
          </InputGroup>
          <Button variant="outline-primary" onClick={this.send} value="Hit">
            Hit
          </Button>
          <Button
            variant="outline-primary"
            onClick={this.delete}
            value="Hit"
          >
            Delete
          </Button>
          <Button variant="outline-danger" onClick={this.close}>
            Close Socket
          </Button>
        </ButtonToolbar>
        <ToastList/>
      </div>
    );
  }
}

AddNotification = connect()(AddNotification);

export default AddNotification;
