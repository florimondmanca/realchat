class WebSocketService {
  private ws: WebSocket = null;

  constructor() {}

  public connect(path: string, handler: (data: string) => void): void {
    this.ws = new WebSocket(`ws://localhost:8080${path}`);
    this.ws.onmessage = (event) => {
      if (typeof event.data === "string") {
        handler(event.data);
      }
    };
    this.ws.onerror = (event: any) => {
      console.error("Websocket error:" + event);
    };
  }

  public close(): void {
    if (!this.ws) {
      return;
    }
    console.log("Closing socket connection...");
    this.ws.close();
    this.ws = null;
  }

  public send(data: string): void {
    this.ws.send(data);
  }
}

export const webSocketService = new WebSocketService();
