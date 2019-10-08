import { SEND_NOTIFICATION_FAILED } from "../redux";

export default function MsgException(message) {
  this.message = message;
  this.name = "MsgException";
  this.type = SEND_NOTIFICATION_FAILED;
}
