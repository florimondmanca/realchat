import { Component, OnInit, OnDestroy } from '@angular/core';
import { Subscription } from 'rxjs';
import { SocketService, User, Message, Event } from '../shared';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit, OnDestroy {

  user: string;
  conn: Subscription;
  messages: Message[] = [];
  messageContent: string;

  constructor(private socketService: SocketService) { }

  ngOnInit() {
    this.initConnection();
  }

  initConnection() {
    this.socketService.initSocket();
    this.conn = this.socketService.onMessage().subscribe(
      (message: Message) => {
        console.log(message);
        this.messages.push(message);
      }
    );
    this.socketService.onOpen().subscribe(
      () => console.log('connected')
    );

    this.socketService.onClose().subscribe(
      () => console.log('disconnected')
    );
  }

  sendMessage() {
    if (!this.messageContent) {
      return
    }
    this.socketService.send({
      userName: this.user || "Anonymous",
      body: this.messageContent,
      timestamp: (+ new Date()).toString()
    });
    this.messageContent = null;
  }

  ngOnDestroy() {
    if (this.conn) {
      this.conn.unsubscribe();
    }
    this.socketService.closeSocket();
  }

}
