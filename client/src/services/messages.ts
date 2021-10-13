import type { IMessagePayload, Message } from "../models/message";
import { addMessage, clearMessages } from "../stores/messages";
import { webSocketService } from "./websocket";

export const messageService = {
  async init(userName: string) {
    await webSocketService.connect("/chat", (data: string) => {
      const message: Message = JSON.parse(data);
      addMessage(message);
    });

    // Send initial user info (as per our protocol).
    const data = JSON.stringify({ userName });
    webSocketService.send(data);
  },

  send(payload: IMessagePayload) {
    const data = JSON.stringify(payload);
    webSocketService.send(data);
  },

  tearDown() {
    clearMessages();
    webSocketService.close();
  },
};
