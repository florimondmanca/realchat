import { EventEmitter } from "events";
import type { Message } from "../models";

class SocketService {
  private socket: WebSocket;
  private events = new EventEmitter();

  constructor() {}

  public init(): void {
    this.socket = new WebSocket("ws://localhost:8080/chat");
    this.socket.onerror = (event: any) =>
      console.log("Websocket error:" + event);
    this.events.emit("ready");
  }

  public close(): void {
    if (!this.socket) {
      return;
    }
    console.log("Closing socket connection...");
    this.socket.close();
    this.socket = null;
  }

  public send(userName: string, body: string) {
    const payload = { userName, body };
    this.socket.send(JSON.stringify(payload));
  }

  private ready$(): Promise<void> {
    return new Promise((resolve) => {
      this.events.on("ready", () => resolve());
    });
  }

  public onMessage(cb: (message: Message) => void) {
    this.ready$().then(() => {
      this.socket.onmessage = (event) => {
        if (typeof event.data === "string") {
          const message: Message = JSON.parse(event.data);
          cb(message);
        }
      };
    });
  }
}

export const socketService = new SocketService();
