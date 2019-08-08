import React from "react";
import {
  Button,
  InputGroup,
  FormControl,
  ButtonToolbar
} from "react-bootstrap";
import { sendMsg, closeSocket } from "../websocket";
import { addNotification } from "../redux/actions";
import { connect } from "react-redux";

class AddNotification extends React.Component {
  constructor(props) {
    super(props);
    this.messageHandler = this.messageHandler.bind(this);
    this.send = this.send.bind(this);
  }
  send() {
    let value = this.state.message_input_text;
    console.log(value);
    sendMsg(value);
    this.props.dispatch(addNotification(value));
  }

  close() {
    console.log("Close Socket");
    closeSocket(1000, "close by user");
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
    );
  }
}

AddNotification = connect()(AddNotification);

export default AddNotification;
