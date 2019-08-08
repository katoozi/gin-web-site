import React from 'react';
import PropTypes from "prop-types";
import { ListGroup } from "react-bootstrap";
import { connect } from "react-redux";


class ToastList extends React.Component{
    render() {
        let notifications = this.props.notifications;
        notifications.map((notification) => <ListGroup.Item>{notification}</ListGroup.Item>)
        return (
            <ListGroup>
                {notifications}
            </ListGroup>
        )
    }
}

ToastList.propTypes = {
  notifications: PropTypes.arrayOf(
      PropTypes.string.isRequired,
  ).isRequired
};

const mapStateToProps = state => {
  return {
    notifications: state.notifications
  };
};

const VisibleToastList = connect(
  mapStateToProps,
)(ToastList);

export default VisibleToastList;
