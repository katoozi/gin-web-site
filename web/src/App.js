import React, { Component } from "react";
import "./App.css";
import { Container } from "react-bootstrap";
import { connect } from "./websocket";
import ToastList from "./components/toast_list";
import AddNotification from "./components/add_notification";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  render() {
    return (
      <Container>
        <div className="app">
          <AddNotification />
        </div>
        <ToastList />
      </Container>
    );
  }
}

export default App;
