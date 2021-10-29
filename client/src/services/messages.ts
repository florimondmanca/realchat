// NOTE: go-socket.io only supports SocketIO server version 1.x and 2.x.
// So we use an only SocketIO client version that does not ship with types.
// See: https://socket.io/docs/v4/client-installation/#version-compatibility
// See: https://github.com/googollee/go-socket.io/issues/443
import io from "socket.io-client";

import type { Message } from "../models/message";
import { addMessage, clearMessages, addUser, removeUser } from "../stores/chat";

let socket: any;

export const messageService = {
  init(userName: string) {
    socket = io("ws://localhost:8000", {
      autoConnect: true,
      transports: ["websocket"],
    });

    socket.on("connect", () => {
      console.debug("connected", socket.id);
    });

    socket.on("join", (userName) => {
      console.debug("join", userName);
      addUser(userName);
    });

    socket.on("leave", (userName) => {
      console.debug("leave", userName);
      removeUser(userName);
    });

    socket.on("msg", (id, timestampSeconds, userName, body) => {
      const message: Message = {
        id,
        timestampSeconds,
        userName,
        body,
      };
      console.debug("msg", message);
      addMessage(message);
    });

    socket.on("disconnect", () => {
      console.debug("disconnect");
    });

    // Send initial user info.
    socket.emit("join", userName, (response) => {
      console.log("join.response:", response);
    });
  },

  send(body: string) {
    socket.emit("msg", body);
  },

  tearDown() {
    clearMessages();
    socket.close();
  },
};
