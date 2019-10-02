import React, { Component } from "react";
import { BrowserRouter as Router, Route, NavLink } from "react-router-dom";
import { Container, Nav } from "react-bootstrap";
import { connect } from "./websocket";
import AddNotification from "./components/add_notification";

function About() {
  return <h2>About</h2>;
}

function Users() {
  return <h2>Users</h2>;
}

class App extends Component {
  componentDidMount() {
    connect();
  }

  render() {
    return (
      <Router>
        <Container>
          <Nav activeKey="/">
            <Nav.Item>
              <NavLink className="nav-link" role="button" to="/">
                Home
              </NavLink>
            </Nav.Item>
            <Nav.Item>
              <NavLink className="nav-link" role="button" to="/about/">
                About
              </NavLink>
            </Nav.Item>
            <Nav.Item>
              <NavLink className="nav-link" role="button" to="/users/">
                Users
              </NavLink>
            </Nav.Item>
          </Nav>

          <Route path="/" exact component={AddNotification} />
          <Route path="/about/" component={About} />
          <Route path="/users/" component={Users} />
        </Container>
      </Router>
    );
  }
}

export default App;
