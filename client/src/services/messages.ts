import type { IMessagePayload, Message } from "../models/message";
import { addMessage, clearMessages } from "../stores/messages";
import { webSocketService } from "./websocket";

export const messageService = {
  init() {
    webSocketService.connect("/chat", (data: string) => {
      const message: Message = JSON.parse(data);
      addMessage(message);
    });
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