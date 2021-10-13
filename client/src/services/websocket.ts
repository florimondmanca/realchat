import { EventEmitter } from "events";
import type { Message } from "../models";

class WebSocketService {
  private ws: WebSocket = null;
  private events = new EventEmitter();

  constructor() {}

  public init(): void {
    this.ws = new WebSocket("ws://localhost:8080/chat");
    this.ws.onerror = (event: any) => {
      console.error("Websocket error:" + event);
    };
    this.events.emit("ready");
  }

  public close(): void {
    if (!this.ws) {
      return;
    }
    console.log("Closing socket connection...");
    this.ws.close();
    this.ws = null;
  }

  public send(userName: string, body: string) {
    const payload = { userName, body };
    this.ws.send(JSON.stringify(payload));
  }

  private ready$(): Promise<void> {
    return new Promise((resolve) => {
      this.events.on("ready", () => resolve());
    });
  }

  public onMessage(cb: (message: Message) => void) {
    this.ready$().then(() => {
      this.ws.onmessage = (event) => {
        if (typeof event.data === "string") {
          const message: Message = JSON.parse(event.data);
          cb(message);
        }
      };
    });
  }
}

export const webSocketService = new WebSocketService();
