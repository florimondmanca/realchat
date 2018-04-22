import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { Message } from './message';
import { User } from './user';

const SERVER_URL = 'ws://localhost:8080/chat';

@Injectable()
export class SocketService {

  private socket: WebSocket;

  initSocket() {
    console.log('Initializing socket connection...');
    this.socket = new WebSocket(SERVER_URL);
  }

  closeSocket() {
    if (this.socket) {
      console.log('Closing socket connection...');
      this.socket.close();
      this.socket = null;
    }
  }

  send(message: Message) {
    this.socket.send(JSON.stringify(message));
  }

  onOpen(): Observable<any> {
    return new Observable<Message>(observer => {
      this.socket.onopen = (event) => observer.next();
    })
  }

  onClose(): Observable<any> {
    return new Observable<Message>(observer => {
      this.socket.onclose = (event) => observer.next();
    })
  }

  onMessage(): Observable<Message> {
    return new Observable<Message>(observer => {
      this.socket.onmessage = (event) => {
        const message: Message = JSON.parse(event.data);
        observer.next(message);
      }
    });
  }

}
