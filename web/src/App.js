import React, { Component } from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import { Container } from "react-bootstrap";
import { connect } from "./websocket";
import AddNotification from "./components/add_notification";

function About() {
  return <h2>About</h2>;
}

function Users() {
  return <h2>Users</h2>;
}

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  render() {
    return (
      <Router>
        <Container>
          <nav>
            <ul>
              <li>
                <Link to="/">Home</Link>
              </li>
              <li>
                <Link to="/about/">About</Link>
              </li>
              <li>
                <Link to="/users/">Users</Link>
              </li>
            </ul>
          </nav>

          <Route path="/" exact component={AddNotification} />
          <Route path="/about/" component={About} />
          <Route path="/users/" component={Users} />
        </Container>
      </Router>
    );
  }
}

export default App;
