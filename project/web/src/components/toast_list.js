import React from "react";
import PropTypes from "prop-types";
import { ListGroup } from "react-bootstrap";
import { connect } from "react-redux";

class ToastList extends React.Component {
  render() {
    const { notifications } = this.props;
    let notifications_list = notifications.map((notification, index) => (
      <ListGroup.Item key={index}>{notification.text}</ListGroup.Item>
    ));
    return <ListGroup>{notifications_list}</ListGroup>;
  }
}

ToastList.propTypes = {
  notifications: PropTypes.arrayOf(
    PropTypes.shape({
      text: PropTypes.string.isRequired
    })
  ).isRequired
};

const mapStateToProps = state => {
  return {
    notifications: state.notification.notifications
  };
};

const VisibleToastList = connect(mapStateToProps)(ToastList);

export default VisibleToastList;
