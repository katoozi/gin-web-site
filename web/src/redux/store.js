import { createStore } from "redux";
import onlineWebSocket from "./index";

const store = createStore(onlineWebSocket);

export default store;
